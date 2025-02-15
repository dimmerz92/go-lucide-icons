package internal

import (
	"flag"
	"fmt"
	"net/http"
	"os"

	"github.com/a-h/templ"
)

func TestServer(args []string) {
	cmd := flag.NewFlagSet("test", flag.ExitOnError)
	port := cmd.Int("p", 8000, "use a port between 1024 and 65535")
	cmd.Parse(args)

	if *port < 1024 || *port > 65535 {
		fmt.Printf(
			"port: %d out of range, use a port between 1024 and 65535",
			*port,
		)
		os.Exit(1)
	}

	http.Handle("GET /", templ.Handler(index()))
	fmt.Printf("Listening on :%d", *port)
	http.ListenAndServe(fmt.Sprintf(":%d", *port), nil)
}
