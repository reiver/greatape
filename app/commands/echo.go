package commands

import . "github.com/reiver/greatape/components/contracts"

func Echo(x IDispatcher, document IDocument) (IEchoResult, error) {
	return x.NewEchoResult(document), nil
}
