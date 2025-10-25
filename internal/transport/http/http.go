package http

import (
	"github.com/rs/zerolog"
	"net"
	"net/http"
)

func RunHTTPServer(httpServer *http.Server, l net.Listener, listenErr chan error, netLogger zerolog.Logger) {
	netLogger.Info().Msgf("starting http server on %s", l.Addr())
	if err := httpServer.Serve(l); err != nil && err != http.ErrServerClosed {
		listenErr <- err
	}
}
