package mschapv2

import (
	"encoding/hex"
	"fmt"
)

func Decode(b []byte) (p Packet, err error) {
	if len(b) == 1 {
		//eap - mschapv2
		return &SimplePacket{
			Code: OpCode(b[0]),
		}, nil
	}
	if len(b) < 4 {
		return nil, fmt.Errorf("[MSCHAPV2.Decode] protocol error 1, len(b)[%d] < 2", len(b))
	}
	code := OpCode(b[0])
	Identifier := uint8(b[1])
	switch code {
	case OpCodeChallenge:
		if len(b) < 21 {
			return nil, fmt.Errorf("[MsChapV2PacketFromEap] protocol error 2 Challenge packet len is less than 21 ")
		}
		resp := &ChallengePacket{}
		copy(resp.Challenge[:], b[5:21])
		resp.Name = string(b[21:])
		resp.Identifier = Identifier
		return resp, nil
	case OpCodeResponse:
		if len(b) < 53 {
			return nil, fmt.Errorf("[MsChapV2PacketFromEap] protocol error 3 Response packet len is less than 53 ")
		}
		resp := &ResponsePacket{}
		copy(resp.PeerChallenge[:], b[5:21])
		copy(resp.NTResponse[:], b[29:53])
		resp.Name = string(b[54:])
		resp.Identifier = Identifier
		return resp, nil
	case OpCodeSuccess:
		resp := &SuccessPacket{}
		hex.Decode(resp.Auth[:], b[6:46])
		resp.Message = string(b[49:])
		resp.Identifier = Identifier
		return resp, nil
	default:
		return nil, fmt.Errorf("[MsChapV2PacketFromEap] can not parse opcode:%s", p.OpCode())
	}
	return p, nil
}
