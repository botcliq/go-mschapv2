package mschapv2

import (
	"bytes"
	"encoding/binary"
	"encoding/hex"
	"fmt"
)

type SuccessPacket struct {
	Identifier uint8
	Auth       [20]byte // the binary format of auth_string
	Message    string
}

func (p *SuccessPacket) String() string {
	return fmt.Sprintf("Code: Success AuthString: %#v Message: %s", p.Auth, p.Message)
}

func (p *SuccessPacket) OpCode() OpCode {
	return OpCodeSuccess
}
func (p *SuccessPacket) Encode() (b []byte) {
	len := 4 + 2 + 40 + 3 + len(p.Message)
	b = make([]byte, len)
	b[0] = byte(p.OpCode())
	b[1] = byte(p.Identifier)
	binary.BigEndian.PutUint16(b[2:4], uint16(len))
	copy(b[4:6], "S=")
	hex.Encode(b[6:46], p.Auth[:])
	out := bytes.ToUpper(b[6:46])
	copy(b[6:46], out)
	copy(b[46:49], " M=")
	copy(b[49:], p.Message)
	return b
}