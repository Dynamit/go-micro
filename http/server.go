package http

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"golang.org/x/net/context"

	"app/pkg/dttp/mux"
)

type Server struct {
	mux *httprouter.Router
}

type Handler interface {
	ServeHTTP(context.Context, http.ResponseWriter, *http.Request)
}

type HandlerFunc func(context.Context, http.ResponseWriter, *http.Request)

func (h HandlerFunc) ServeHTTP(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	h(ctx, w, r)
}

// NewServer creates a new Server.
func NewServer() *Server {
	return &Server{
		mux: httprouter.New(),
	}
}

// NewMux provides the ability to create simple middleware groups.
func (s *Server) NewMux() *Mux {
	return &Mux{
		server: s,
	}
}

// Serve calls http.ListenAndServe with the Server's mux.
func (s *Server) Serve(addr string) error {

	return http.ListenAndServe(addr, s.mux)

}

// ServeTLS calls http.ListenAndServeTLS with the Server's mux.
func (s *Server) ServeTLS(addr string, certFile string, keyFile string) error {

	return http.ListenAndServeTLS(addr, certFile, keyFile, s.mux)

}

// Get registers a new request handle with context.
func (s *Server) Get(route string, handler Handler) {

	s.mux.GET(route, s.contextualize(handler))

}

// Head registers a new request handle with context.
func (s *Server) Head(route string, handler Handler) {

	s.mux.HEAD(route, s.contextualize(handler))

}

// Options registers a new request handle with context.
func (s *Server) Options(route string, handler Handler) {

	s.mux.OPTIONS(route, s.contextualize(handler))

}

// Post registers a new request handle with context.
func (s *Server) Post(route string, handler Handler) {

	s.mux.POST(route, s.contextualize(handler))

}

// Put registers a new request handle with context.
func (s *Server) Put(route string, handler Handler) {

	s.mux.PUT(route, s.contextualize(handler))

}

// Patch registers a new request handle with context.
func (s *Server) Patch(route string, handler Handler) {

	s.mux.PATCH(route, s.contextualize(handler))

}

// Delete registers a new request handle with context.
func (s *Server) Delete(route string, handler Handler) {

	s.mux.DELETE(route, s.contextualize(handler))

}

// contextualize is a helper function for creating a handler with context.
func (s *Server) contextualize(handler Handler) httprouter.Handle {

	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

		ctx := context.Background()

		ctx = mux.NewContext(ctx, ps)

		handler.ServeHTTP(ctx, w, r)

	}

}
