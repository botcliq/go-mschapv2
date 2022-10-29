package mschapv2

import (
	"encoding/binary"
	"fmt"
)

type ResponsePacket struct {
	Identifier    uint8
	PeerChallenge [16]byte //16byte
	NTResponse    [24]byte //24byte
	Name          string
}

func (p *ResponsePacket) String() string {
	return fmt.Sprintf("Code: Response PeerChallenge: %#v NTResponse:%#v Name:%#v", p.PeerChallenge, p.NTResponse, p.Name)
}
func (p *ResponsePacket) OpCode() OpCode {
	return OpCodeResponse
}
func (p *ResponsePacket) Encode() (b []byte) {
	len := 4 + 1 + 49 + len(p.Name)
	b = make([]byte, len)
	b[0] = byte(p.OpCode())
	b[1] = byte(p.Identifier)
	binary.BigEndian.PutUint16(b[2:4], uint16(len))
	b[4] = 49
	copy(b[5:21], p.PeerChallenge[:])
	copy(b[29:53], p.NTResponse[:])
	copy(b[54:], p.Name)
	return b
}
