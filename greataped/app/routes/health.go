package routes

import (
	"contracts"
	"server/route"
)

var Health = route.New(contracts.HttpGet, "/health", func(x contracts.IContext) error {
	return x.Nothing()
})
