package godemos

import "net/http"

func server(addr string, handler http.Handler, stop <-chan struct{}) error {
	s := http.Server{
		Addr: addr,
	}
}
