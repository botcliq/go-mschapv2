package mschapv2

import (
	"strconv"
)

type Packet interface {
	String() string
	OpCode() OpCode
	Encode() (b []byte)
}

type OpCode uint8

const (
	OpCodeChallenge      OpCode = 1
	OpCodeResponse       OpCode = 2
	OpCodeSuccess        OpCode = 3
	OpCodeFailure        OpCode = 4
	OpCodeChangePassword OpCode = 7
)

func (c OpCode) String() string {
	switch c {
	case OpCodeChallenge:
		return "Challenge"
	case OpCodeResponse:
		return "Response"
	case OpCodeSuccess:
		return "Success"
	case OpCodeFailure:
		return "Failure"
	case OpCodeChangePassword:
		return "ChangePassword"
	default:
		return "unknow OpCode " + strconv.Itoa(int(c))
	}
}
