package http

import (
	"net"
	"net/http"

	"github.com/rs/zerolog"
)

func RunHTTPServer(httpServer *http.Server, l net.Listener, listenErr chan error, netLogger zerolog.Logger) {
	netLogger.Info().Msgf("starting http server on %s", l.Addr())
	err := httpServer.Serve(l)
	if err != nil && err != http.ErrServerClosed {
		listenErr <- err
	}
}
