package http

import (
	"fmt"
	"net/http"
	"strings"
)

type HttpServer struct {
	server *http.Server
	mux    *http.ServeMux
}

type customRouter struct {
	mux *http.ServeMux
}

var routes = map[string]func(http.ResponseWriter, *http.Request){
	"/hello": HandleHello,
	"/test":  HandleTest,
}

func GetRoutes() []string {
	routesSlice := make([]string, len(routes))
	i := 0
	for route := range routes {
		routesSlice[i] = route
		i++
	}
	return routesSlice
}

func NewHttpServer(address string) (*HttpServer, error) {
	mux := http.NewServeMux()
	for route, handler := range routes {
		mux.HandleFunc(route, handler)
	}

	mux.Handle("/", &customRouter{mux})

	server := &HttpServer{
		server: &http.Server{
			Addr:    address,
			Handler: mux,
		},
		mux: mux,
	}
	return server, nil
}

func (s *HttpServer) Start() {
	defer s.Stop()

	fmt.Printf("Server is starting on %s\n", s.server.Addr)
	server, err := NewHttpServer("localhost:8080")
	if err != nil {
		panic(err)
	}

	if err := server.server.ListenAndServe(); err != nil {
		panic(err)
	}
}

func (cr *customRouter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		HandleRoot(w, r)
	} else {
		HandleDynamic(w, r)
	}
}

func HandleRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "[HandleRoot] Hello! Welcome to Go Networking example")
}

func HandleDynamic(w http.ResponseWriter, r *http.Request) {
	dynamicID := strings.TrimPrefix(r.URL.Path, "/")
	fmt.Fprintf(w, "%s", dynamicID)
}

func HandleHello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "[HandleHello] Hello from /hello!]")
}

func HandleTest(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "[HandleTest] Hello from test!]")
}

func (s *HttpServer) Stop() {
	fmt.Printf("Server is stopping on %s\n", s.server.Addr)
	s.server.Close()
}
