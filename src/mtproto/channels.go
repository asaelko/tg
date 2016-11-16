package mtproto

import (
	"strings"
	"errors"
	"fmt"
)

type Channel struct {
	mtproto *MTProto
	full TL_channelFull

	Id int32
	Name string
	AccessHash int64
	LastMessageId int32
}

func newChannel(m *MTProto) *Channel {
	channel := new(Channel)
	channel.mtproto = m
	channel.full = TL_channelFull{}
	return channel
}

func (m *MTProto)SearchChannel(name string) (*Channel, error) {
	resp := make(chan TL, 1)
	m.queueSend <- packetToSend{
		TL_contacts_search{
			name,
			1,
		},
		resp,
	}
	x := <-resp
	res, ok := x.(TL_contacts_found)
	if !ok {
		return nil, fmt.Errorf("RPC: %#v", x)
	}
	channel := newChannel(m)

	for _, foundChat := range res.chats {
		switch foundChat.(type){
			case TL_channel:
				foundChat := foundChat.(TL_channel)
				if strings.ToLower(foundChat.username) == strings.ToLower(name) {
					//dump(foundChat)
					channel.Id = foundChat.id
					channel.AccessHash = foundChat.access_hash
					channel.Name = name
				}
		}
	}

	if channel.isEmpty() {
		return nil, errors.New("Channel not found")
	}

	return channel, nil
}

func (channel *Channel)isEmpty() bool {
	if channel.Id != 0 {
		return true
	}

	return false
}

func (channel *Channel)GetFullInfo() error {
	resp := make(chan TL, 1)

	channel.mtproto.queueSend <- packetToSend{
		TL_channels_getFullChannel{
			TL_inputChannel{
				channel.Id,
				channel.AccessHash,
			},
		},
		resp,
	}
	x := <-resp
	switch x.(type){
		case TL_channelFull:
			channel.full = x.(TL_channelFull)
		default:
			return errors.New("Cannot request full channel")
	}

	return nil
}

func (channel *Channel)GetMessages(offset_id, offset_date, add_offset, limit, max_id, min_id int32) {
	resp := make(chan TL, 1)

	channel.mtproto.queueSend <- packetToSend{
		TL_messages_getHistory{
			TL_inputPeerChannel{
				channel.Id,
				channel.AccessHash,
			},
			offset_id,
			offset_date,
			add_offset,
			limit,
			max_id,
			min_id,
		},
		resp,
	}
	x := <-resp
	dump(x)
}

func (channel *Channel)GetViews(msgIds []int32, increment bool) map[int32]int32{

	resp := make(chan TL, 1)

	var tlIncrement TL
	if increment {
		tlIncrement = TL_boolTrue{}
	} else {
		tlIncrement = TL_boolFalse{}
	}

	channel.mtproto.queueSend <- packetToSend{
		TL_messages_getMessagesViews{
			TL_inputPeerChannel{
				channel.Id,
				channel.AccessHash,
			},
			msgIds,
			tlIncrement,
		},
		resp,
	}
	x := <-resp
	dump(x);
	return nil
}