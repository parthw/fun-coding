package main

import (
	"encoding/json"
)

type MyDelegate struct {
	msgCh chan []byte
}

func (d *MyDelegate) NotifyMsg(msg []byte) {
	d.msgCh <- msg
}

func (d *MyDelegate) NodeMeta(limit int) []byte {
	// not use, noop
	return []byte("")
}
func (d *MyDelegate) LocalState(join bool) []byte {
	// not use, noop
	return []byte("")
}
func (d *MyDelegate) GetBroadcasts(overhead, limit int) [][]byte {
	// not use, noop
	return nil
}
func (d *MyDelegate) MergeRemoteState(buf []byte, join bool) {
	// not use
}

type MyMessage struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

func (m *MyMessage) Bytes() []byte {
	data, err := json.Marshal(m)
	if err != nil {
		return []byte("")
	}
	return data
}
func ParseMyMessage(data []byte) (*MyMessage, bool) {
	msg := new(MyMessage)
	if err := json.Unmarshal(data, &msg); err != nil {
		return nil, false
	}
	return msg, true
}
