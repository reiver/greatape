package services

import (
	"fmt"
	"strings"

	. "github.com/reiver/greatape/components/api/protobuf"
	. "github.com/reiver/greatape/components/constants"
	. "github.com/reiver/greatape/components/contracts"
	. "github.com/reiver/greatape/components/core"
	. "github.com/xeronith/diamante/contracts/service"
)

func SystemCallService(context IContext, input *SystemCallRequest) (result *SystemCallResult, err error) {
	source := "system_call"
	/* //////// */ Conductor.LogRemoteCall(context, INIT, source, input, result, err)
	defer func() { Conductor.LogRemoteCall(context, DONE, source, input, result, err) }()

	context.Logger().SysCall(fmt.Sprintf("SYSCALL: %s", input.Command))

	args := strings.Split(input.Command, " ")
	if len(args) < 1 {
		return nil, ERROR_NOT_IMPLEMENTED
	}

	result = context.ResultContainer().(*SystemCallResult)

	switch args[0] {
	case "reload":
		if len(args) < 2 {
			return nil, ERROR_NOT_IMPLEMENTED
		}

		componentName := args[1]
		if component := Conductor.GetSystemComponent(componentName); component == nil {
			return nil, ERROR_SYSTEM_COMPONENT_NOT_FOUND
		} else if err := component.Reload(); err != nil {
			return nil, err
		} else {
			return result, nil
		}
	default:
		if err := context.SystemCall(args); err != nil {
			return nil, err
		}

		return result, nil
	}
}
