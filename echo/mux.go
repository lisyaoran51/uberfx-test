package echo

import "net/http"

// NewServeMux builds a ServeMux that will route requests
// to the given EchoHandler.
// func NewServeMux(route1, route2 Route) *http.ServeMux {
// 	mux := http.NewServeMux()
// 	mux.Handle(route1.Pattern(), route1)
// 	mux.Handle(route2.Pattern(), route2)
// 	return mux
// }

func NewServeMux(routes []Route) *http.ServeMux {
	mux := http.NewServeMux()
	for _, route := range routes {
		mux.Handle(route.Pattern(), route)
	}
	return mux
}
