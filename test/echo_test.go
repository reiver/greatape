package test

import (
	"errors"

	. "github.com/reiver/greatape/components/api/protobuf"
	. "github.com/reiver/greatape/components/contracts"
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
