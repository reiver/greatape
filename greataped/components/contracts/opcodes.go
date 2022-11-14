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

	//SignupOperation
	SIGNUP_REQUEST = 0x48DB23BF
	SIGNUP_RESULT  = 0x83D062B4

	//VerifyOperation
	VERIFY_REQUEST = 0x8B78F7F6
	VERIFY_RESULT  = 0x2C8A8A49

	//LoginOperation
	LOGIN_REQUEST = 0xF480F151
	LOGIN_RESULT  = 0xBE819605
)

var OPCODES = Opcodes{
	0x00000000: "N/A",
	0x0541BD72: "ECHO",
	0xAB2FF7D4: "Echo",
	0x48DB23BF: "SIGNUP",
	0x83D062B4: "Signup",
	0x8B78F7F6: "VERIFY",
	0x2C8A8A49: "Verify",
	0xF480F151: "LOGIN",
	0xBE819605: "Login",
}
