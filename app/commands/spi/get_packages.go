package spi

import (
	_ "embed"

	. "github.com/reiver/greatape/components/contracts"
)

//go:embed _packages.txt
var packages []byte

func GetPackages(x IDispatcher) (IGetPackagesResult, error) {
	return x.NewGetPackagesResult(packages), nil
}
