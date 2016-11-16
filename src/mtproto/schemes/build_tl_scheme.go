package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

type nametype struct {
	name  string
	_type string
}

type constuctor struct {
	id        string
	predicate string
	params    []nametype
	_type     string
}

func normalize(s string) string {
	x := []byte(s)
	for i, r := range x {
		if r == '.' {
			x[i] = '_'
		}
	}
	y := string(x)
	if y == "type" {
		return "_type"
	}
	return y
}

func main() {
	var err error
	var parsed interface{}

	// read json file from stdin
	data, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		fmt.Println(err)
		return
	}

	// parse json
	d := json.NewDecoder(bytes.NewReader(data))
	d.UseNumber()
	err = d.Decode(&parsed)
	if err != nil {
		fmt.Println(err)
		return
	}

	// process constructors
	_order := make([]string, 0, 1000)
	_cons := make(map[string]constuctor, 1000)
	_types := make(map[string][]string, 1000)

	parsefunc := func(data []interface{}, kind string) {
		for _, data := range data {
			data := data.(map[string]interface{})

			// id
			idx, err := strconv.Atoi(data["id"].(string))
			if err != nil {
				fmt.Println(err)
				return
			}
			_id := fmt.Sprintf("0x%08x", uint32(idx))

			// predicate
			_predicate := normalize(data[kind].(string))

			if _predicate == "vector" {
				continue
			}

			// params
			_params := make([]nametype, 0, 16)
			params := data["params"].([]interface{})
			for _, params := range params {
				params := params.(map[string]interface{})
				_params = append(_params, nametype{normalize(params["name"].(string)), normalize(params["type"].(string))})
			}

			// type
			_type := normalize(data["type"].(string))

			_order = append(_order, _predicate)
			_cons[_predicate] = constuctor{_id, _predicate, _params, _type}
			if kind == "predicate" {
				_types[_type] = append(_types[_type], _predicate)
			}
		}
	}
	parsefunc(parsed.(map[string]interface{})["constructors"].([]interface{}), "predicate")
	parsefunc(parsed.(map[string]interface{})["methods"].([]interface{}), "method")

	// constants
	fmt.Print(`package mtproto
			import (
				"fmt"
				"reflect"
				"strconv"
			)
			const (
				`)

	for _, key := range _order {
		c := _cons[key]
		fmt.Printf("crc_%s = %s\n", c.predicate, c.id)
	}
	fmt.Print(")\n\n")

	// type structs
	for _, key := range _order {
		c := _cons[key]
		fmt.Printf("type TL_%s struct {\n", c.predicate)
		for _, t := range c.params {
			fmt.Printf("%s\t", t.name)

			var bite int
			var flag_type string
			var tl_type string

			flags, _ := fmt.Sscanf(t._type, "flags_%d?%s", &bite, &flag_type)
			if flags == 2 {
				t._type = flag_type
			}

			switch t._type {
			case "true":
				fmt.Printf("bool")
			case "int":
				fmt.Printf("int32")
			case "long":
				fmt.Printf("int64")
			case "string":
				fmt.Printf("string")
			case "double":
				fmt.Printf("float64")
			case "bytes":
				fmt.Printf("[]byte")
			case "Vector<int>":
				fmt.Printf("[]int32")
			case "Vector<long>":
				fmt.Printf("[]int64")
			case "Vector<string>":
				fmt.Printf("[]string")
			case "Vector<double>":
				fmt.Printf("[]float64")
			case "!X":
				fmt.Printf("TL")
			case "#":
				fmt.Printf("int32")
			default:
				var inner string
				n, _ := fmt.Sscanf(t._type, "Vector<%s", &inner)
				if n == 1 {
					tl_type = inner[:len(inner)-1]
					fmt.Print("[]TL")
				} else {
					tl_type = t._type
					fmt.Print("TL")
				}
			}
			if flags == 2 {
				fmt.Printf(" `flag_byte:\"%d\"`", bite)
			}
			if len(tl_type) != 0 {
				fmt.Printf(" // %s", tl_type)
			}
			fmt.Printf("\n")
		}
		fmt.Printf("}\n\n")
	}

	// encode funcs
	for _, key := range _order {
		c := _cons[key]
		fmt.Printf("func (e TL_%s) encode() []byte {\n", c.predicate)

		fmt.Printf("x := NewEncodeBuf(512)\n")
		fmt.Printf("x.UInt(crc_%s)\n", c.predicate)
		var paramsBuffer bytes.Buffer
		var s string
		var has_flags bool = false

		// lets generate flags

		for _, t := range c.params {
			var bite int
			var flag_type string

			flags, _ := fmt.Sscanf(t._type, "flags_%d?%s", &bite, &flag_type)
			if flags == 2 {
				t._type = flag_type
			}

			switch t._type {
			case "#":
				fmt.Printf("var flags int32 = 0\n")
				s = fmt.Sprintf("x.Int(e.%s)\n", t.name)
				has_flags = true
			case "true":
				fmt.Printf("if(e.%s) { flags = flags | 1 << %d }\n", t.name, bite)
			}
		}
		if has_flags {
			// reflection for flagged types
			fmt.Printf("s := TL_%s{}\n", c.predicate)
			fmt.Printf("sRef := reflect.TypeOf(s)\n")
			fmt.Printf(`
					for i := 0; i < sRef.NumField(); i++ {
							field := sRef.Field(i)
							if flag_byte, ok := field.Tag.Lookup("flag_byte"); ok {
								b, _ := strconv.ParseUint(flag_byte, 10, 64)
								flags = flags | 1<<b
							}
						}
				`)
		}

		// lets generate data
		for _, t := range c.params {
			var bite int
			var flag_type string

			flags, _ := fmt.Sscanf(t._type, "flags_%d?%s", &bite, &flag_type)
			if flags == 2 {
				t._type = flag_type
			}

			switch t._type {
			case "#":
				s = fmt.Sprintf("x.Int(flags)\n")
			case "true":
				s = fmt.Sprint("")
			case "int":
				s = fmt.Sprintf("x.Int(e.%s)\n", t.name)
			case "long":
				s = fmt.Sprintf("x.Long(e.%s)\n", t.name)
			case "string":
				s = fmt.Sprintf("x.String(e.%s)\n", t.name)
			case "double":
				s = fmt.Sprintf("x.Double(e.%s)\n", t.name)
			case "bytes":
				s = fmt.Sprintf("x.StringBytes(e.%s)\n", t.name)
			case "Vector<int>":
				s = fmt.Sprintf("x.VectorInt(e.%s)\n", t.name)
			case "Vector<long>":
				s = fmt.Sprintf("x.VectorLong(e.%s)\n", t.name)
			case "Vector<string>":
				s = fmt.Sprintf("x.VectorString(e.%s)\n", t.name)
			case "!X":
				s = fmt.Sprintf("x.Bytes(e.%s.encode())\n", t.name)
			case "Vector<double>":
				panic(fmt.Sprintf("Unsupported %s", t._type))
			default:
				var inner string
				n, _ := fmt.Sscanf(t._type, "Vector<%s", &inner)
				if n == 1 {
					s = fmt.Sprintf("x.Vector(e.%s)\n", t.name)
				} else {
					s = fmt.Sprintf("x.Bytes(e.%s.encode())\n", t.name)
				}
			}
			paramsBuffer.WriteString(s)
		}
		fmt.Printf("%s", paramsBuffer.String())
		fmt.Printf("return x.buf\n")
		fmt.Printf("}\n\n")
	}

	// decode funcs
	fmt.Println(`
func (m *DecodeBuf) ObjectGenerated(constructor uint32) (r TL) {
	switch constructor {`)

	for _, key := range _order {
		c := _cons[key]
		fmt.Printf("case crc_%s:\n", c.predicate)

		var hasFlags bool = false

		for _, t := range c.params {
			switch t._type {
			case "#":
				hasFlags = true
			}
		}
		if hasFlags {
			fmt.Printf("rInit := TL_%s{}\n", c.predicate)
			fmt.Printf("rInit.flags = m.Int()\n") // get flags

			for _, t := range c.params {
				var bite int
				var flag_type string

				flags, _ := fmt.Sscanf(t._type, "flags_%d?%s", &bite, &flag_type)
				if flags == 2 {
					t._type = flag_type
					fmt.Printf("if rInit.flags & (1 << %d) == (1<<%d) {", bite, bite)
				}

				switch t._type {
				case "#":
					// do nothing
				case "true":
					fmt.Printf("rInit.%s = true\n", t.name)
				case "int":
					fmt.Printf("rInit.%s = m.Int()\n", t.name)
				case "long":
					fmt.Printf("rInit.%s = m.Long()\n", t.name)
				case "string":
					fmt.Printf("rInit.%s = m.String()\n", t.name)
				case "double":
					fmt.Printf("rInit.%s = m.Double()\n", t.name)
				case "bytes":
					fmt.Printf("rInit.%s = m.StringBytes()\n", t.name)
				case "Vector<int>":
					fmt.Printf("rInit.%s = m.VectorInt()\n", t.name)
				case "Vector<long>":
					fmt.Printf("rInit.%s = m.VectorLong()\n", t.name)
				case "Vector<string>":
					fmt.Printf("rInit.%s = m.VectorString()\n", t.name)
				case "!X":
					fmt.Printf("rInit.%s = m.Object()\n", t.name)
				case "Vector<double>":
					panic(fmt.Sprintf("Unsupported %s", t._type))
				default:
					var inner string
					n, _ := fmt.Sscanf(t._type, "Vector<%s", &inner)
					if n == 1 {
						fmt.Printf("rInit.%s = m.Vector()\n", t.name)
					} else {
						fmt.Printf("rInit.%s = m.Object()\n", t.name)
					}
				}

				if flags == 2 {
					fmt.Printf(" }\n")
				}
			}
			fmt.Printf("r = rInit\n\n")
		} else {
			fmt.Printf("r = TL_%s{\n", c.predicate)

			for _, t := range c.params {
				var bite int
				var flag_type string

				flags, _ := fmt.Sscanf(t._type, "flags_%d?%s", &bite, &flag_type)
				if flags == 2 {
					t._type = flag_type
				}

				switch t._type {
				case "int":
					fmt.Print("m.Int(),\n")
				case "long":
					fmt.Print("m.Long(),\n")
				case "string":
					fmt.Print("m.String(),\n")
				case "double":
					fmt.Print("m.Double(),\n")
				case "bytes":
					fmt.Print("m.StringBytes(),\n")
				case "Vector<int>":
					fmt.Print("m.VectorInt(),\n")
				case "Vector<long>":
					fmt.Print("m.VectorLong(),\n")
				case "Vector<string>":
					fmt.Print("m.VectorString(),\n")
				case "!X":
					fmt.Print("m.Object(),\n")
				case "Vector<double>":
					panic(fmt.Sprintf("Unsupported %s", t._type))
				default:
					var inner string
					n, _ := fmt.Sscanf(t._type, "Vector<%s", &inner)
					if n == 1 {
						fmt.Print("m.Vector(),\n")
					} else {
						fmt.Print("m.Object(),\n")
					}
				}
			}
			fmt.Print("}\n\n")

		}

	}

	fmt.Println(`
	default:
		m.err = fmt.Errorf("Unknown constructor: %08x", constructor)
		return nil

	}

	if m.err != nil {
		return nil
	}

	return
}`)

}
