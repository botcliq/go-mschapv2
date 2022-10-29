package mschapv2

import (
	"encoding/binary"
	"fmt"
)

type ChallengePacket struct {
	Identifier uint8
	Challenge  [16]byte
	Name       string
}

func (p *ChallengePacket) String() string {
	return fmt.Sprintf("Code: Challenge Challenge: %#v Name: %s", p.Challenge, p.Name)
}
func (p *ChallengePacket) OpCode() OpCode {
	return OpCodeChallenge
}
func (p *ChallengePacket) Encode() (b []byte) {
	len := 4 + 1 + 16 + len(p.Name)
	b = make([]byte, len)
	b[0] = byte(p.OpCode())
	b[1] = byte(p.Identifier)
	binary.BigEndian.PutUint16(b[2:4], uint16(len))
	b[4] = 16
	copy(b[5:21], p.Challenge[:])
	copy(b[21:], p.Name)
	return b
}
