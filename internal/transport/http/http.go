package http

import (
	"fmt"
	"net"
	"net/http"
)

func RunHTTPServer(httpServer *http.Server, l net.Listener, listenErr chan error) {
	fmt.Println("HTTP server listening on ", l.Addr())
	if err := httpServer.Serve(l); err != nil && err != http.ErrServerClosed {
		listenErr <- err
	}
}
