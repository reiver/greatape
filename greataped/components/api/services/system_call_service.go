package services

import (
	"fmt"
	"strings"

	. "github.com/xeronith/diamante/contracts/service"
	. "rail.town/infrastructure/components/api/protobuf"
	. "rail.town/infrastructure/components/constants"
	. "rail.town/infrastructure/components/contracts"
	"rail.town/infrastructure/components/core"
)

func SystemCallService(context IContext, input *SystemCallRequest) (result *SystemCallResult, err error) {
	core.Conductor.LogRemoteCall(context, INITIALIZE, "system_call", input, result, err)
	defer func() { core.Conductor.LogRemoteCall(context, FINALIZE, "system_call", input, result, err) }()

	context.Logger().SysCall(fmt.Sprintf("SYSCALL: %s", input.Command))

	args := strings.Split(input.Command, " ")
	if len(args) < 1 {
		return nil, ERROR_NOT_IMPLEMENTED
	}

	switch args[0] {
	case "reload":
		if len(args) < 2 {
			return nil, ERROR_NOT_IMPLEMENTED
		}

		componentName := args[1]
		if component := core.Conductor.GetSystemComponent(componentName); component == nil {
			return nil, ERROR_SYSTEM_COMPONENT_NOT_FOUND
		} else if err := component.Reload(); err != nil {
			return nil, err
		} else {
			return context.ResultContainer().(*SystemCallResult), nil
		}
	default:
		if err := context.SystemCall(args); err != nil {
			return nil, err
		}

		return context.ResultContainer().(*SystemCallResult), nil
	}
}
