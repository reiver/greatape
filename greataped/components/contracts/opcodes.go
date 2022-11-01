package contracts

import . "github.com/xeronith/diamante/contracts/server"

// noinspection GoSnakeCaseUsage
const (
	//SystemCallOperation
	SYSTEM_CALL_REQUEST = 0x00001000
	SYSTEM_CALL_RESULT  = 0xF0001000

	//EchoOperation
	ECHO_REQUEST = 0x0541BD72
	ECHO_RESULT  = 0xAB2FF7D4
)

var OPCODES = Opcodes{
	0x00000000: "N/A",
	0x0541BD72: "ECHO",
	0xAB2FF7D4: "Echo",
}
