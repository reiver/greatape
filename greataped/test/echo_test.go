package test

import (
	"errors"

	. "rail.town/infrastructure/components/api/protobuf"
	. "rail.town/infrastructure/components/contracts"
)

func echo(api IApi) error {

	// Echo
	{
		input := &EchoRequest{
			Document: &Document{
				Content: "{}",
			},
		}

		output, err := api.Echo(input)
		if err != nil {
			return err
		}

		if output.Document.Content != input.Document.Content {
			return errors.New("echo_failed")
		}
	}

	return nil
}
