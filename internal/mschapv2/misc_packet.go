package mschapv2

import (
	"fmt"
)

type SimplePacket struct {
	Code OpCode
}

func (p *SimplePacket) OpCode() OpCode {
	return p.Code
}

func (p *SimplePacket) String() string {
	return fmt.Sprintf("Code: %s", p.OpCode())
}

func (p *SimplePacket) Encode() (b []byte) {
	b = make([]byte, 1)
	b[0] = byte(p.OpCode())
	return b
}

type ReplySuccessPacketRequest struct {
	AuthenticatorChallenge [16]byte
	Response               *ResponsePacket
	Username               []byte
	Password               []byte
	Message                string
}

func ReplySuccessPacket(req *ReplySuccessPacketRequest) (p *SuccessPacket) {
	Auth := GenerateAuthenticatorResponse(req.Password, req.Response.NTResponse, req.Response.PeerChallenge, req.AuthenticatorChallenge, req.Username)
	return &SuccessPacket{
		Identifier: req.Response.Identifier,
		Auth:       Auth,
		Message:    req.Message,
	}
}
