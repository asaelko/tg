package mtproto

import (
	"bytes"
	"fmt"
)

type Auth struct {
	UserId 	    int32
	AuthKey     []byte
	AuthKeyHash []byte
	ServerSalt  []byte
}

func (m *MTProto) Auth() error {
	var phone_number string
	var authSentCode TL_auth_sentCode

	fmt.Print("Enter phone number: ")
	fmt.Scanf("%s", &phone_number)

	// (TL_auth_sendCode)
	flag := true
	for flag {
		resp := make(chan TL, 1)
		m.queueSend <- packetToSend{
			TL_auth_sendCode{
				0,    // flags
				true, //allowFlashcall
				phone_number,
				TL_boolTrue{},
				appId,
				appHash,
			},
			resp,
		}
		x := <-resp

		switch x.(type) {
		case TL_auth_sentCode:
			authSentCode = x.(TL_auth_sentCode)
			flag = false
		case TL_rpc_error:
			x := x.(TL_rpc_error)
			if x.error_code != 303 {
				return fmt.Errorf("RPC error_code: %d", x.error_code)
			}
			var newDc int32
			n, _ := fmt.Sscanf(x.error_message, "PHONE_MIGRATE_%d", &newDc)
			if n != 1 {
				n, _ := fmt.Sscanf(x.error_message, "NETWORK_MIGRATE_%d", &newDc)
				if n != 1 {
					return fmt.Errorf("RPC error_string: %s", x.error_message)
				}
			}

			newDcAddr, ok := m.dclist[newDc]
			if !ok {
				return fmt.Errorf("Wrong DC index: %d", newDc)
			}
			err := m.Reconnect(newDc, newDcAddr)
			if err != nil {
				return err
			}
		default:
			return fmt.Errorf("Got: %T", x)
		}

	}

	var code int

	fmt.Print("Enter code: ")
	fmt.Scanf("%d", &code)

	if authSentCode.phone_registered {
		resp := make(chan TL, 1)
		m.queueSend <- packetToSend{
			TL_auth_signIn{phone_number, authSentCode.phone_code_hash, fmt.Sprintf("%d", code)},
			resp,
		}
		x := <-resp
		switch x.(type) {
		case TL_auth_authorization:
			auth, ok := x.(TL_auth_authorization)

			fmt.Printf("%q %q", auth, ok)
			if !ok {
				return fmt.Errorf("RPC: %#v", x)
			}
		case TL_rpc_error:
			x := x.(TL_rpc_error)
			if x.error_message == "SESSION_PASSWORD_NEEDED" {
				err := m.authPasswordNeed(phone_number)
				if err != nil {
					return err
				}
			}
		}
		//userSelf := auth.user.(TL_userSelf)
		//fmt.Printf("Signed in: id %d name <%s %s>\n", userSelf.id, userSelf.first_name, userSelf.last_name)
	} else {
		return fmt.Errorf("Cannot sign up yet")
	}

	return nil
}

func (m *MTProto) authPasswordNeed(phone_number string) error {
	var password string

	resp := make(chan TL, 1)
	m.queueSend <- packetToSend{
		TL_account_getPassword{},
		resp,
	}
	x := <-resp
	switch x.(type) {
	case TL_account_password:
		x := x.(TL_account_password)
		fmt.Printf("TFA enabled\n")
		fmt.Print("Enter password: ")
		fmt.Scanf("%s", &password)

		var passwordHash []byte = m.makePasswordHash(x.current_salt, password)

		resp = make(chan TL, 1)
		m.queueSend <- packetToSend{
			TL_auth_checkPassword{passwordHash},
			resp,
		}
		_ = <-resp

	case TL_rpc_error:
		x := x.(TL_rpc_error)

		return fmt.Errorf(x.error_message)
	}

	return nil
}

func (m *MTProto) makePasswordHash(salt []byte, password string) []byte {
	var passwordBuffer bytes.Buffer
	passwordBuffer.Write(salt)
	passwordBuffer.WriteString(password)
	passwordBuffer.Write(salt)
	return sha256(passwordBuffer.Bytes())
}
