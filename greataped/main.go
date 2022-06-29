package main

import (
	_ "github.com/reiver/greatape/greataped/arg"
	"github.com/reiver/greatape/greataped/cfg"

	"fmt"
	"net/http"
	"os"
)

func main() {
	fmt.Println("greateape daemon 🐒")

	{
		var addr string = fmt.Sprintf(":%d", cfg.HTTPPort)

		err := http.ListenAndServe(addr, nil)
		if nil != err {
			fmt.Fprintf(os.Stderr, "ERROR: problem with HTTP server: %s\n", err)
			os.Exit(1)
			return
		}
	}
}
