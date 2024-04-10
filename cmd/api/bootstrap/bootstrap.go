package bootstrap

import "github.com/gmd3v/gohex/internal/platform/server"

const (
	port    = 8080
	address = "localhost"
)

func Start() error {
	srv := server.New(address, port)
	return srv.Run()
}
