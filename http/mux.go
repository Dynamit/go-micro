package http

type Middleware func(Handler) Handler

type Mux struct {
	server *Server
	mw     []Middleware
}

// Use adds the provided Handler to the Mux's list of middleware.
func (m *Mux) Use(fn Middleware) {
	m.mw = append(m.mw, fn)
}

// build is a helper function for reducing the handler chain.
func (m *Mux) build(handler HandlerFunc) Handler {

	var route Handler

	length := len(m.mw)

	// loop through middleware, in reverse, and determine
	// the actual callback function resulting from it.
	for i, _ := range m.mw {
		r := m.mw[length-i-1]
		if i == 0 {
			route = r(HandlerFunc(handler))
		} else {
			route = r(route)
		}
	}

	return route

}

// Get calls the underlying Server.Get with the reduced handler chain.
func (m *Mux) Get(path string, handler HandlerFunc) {

	m.server.Get(path, m.build(handler))

}

// Head calls the underlying Server.Head with the reduced handler chain.
func (m *Mux) Head(path string, handler HandlerFunc) {

	m.server.Head(path, m.build(handler))

}

// Options calls the underlying Server.Options with the reduced handler chain.
func (m *Mux) Options(path string, handler HandlerFunc) {

	m.server.Options(path, m.build(handler))

}

// Post calls the underlying Server.Post with the reduced handler chain.
func (m *Mux) Post(path string, handler HandlerFunc) {

	m.server.Post(path, m.build(handler))

}

// Put calls the underlying Server.Put with the reduced handler chain.
func (m *Mux) Put(path string, handler HandlerFunc) {

	m.server.Put(path, m.build(handler))

}

// Patch calls the underlying Server.Patch with the reduced handler chain.
func (m *Mux) Patch(path string, handler HandlerFunc) {

	m.server.Patch(path, m.build(handler))

}

// Delete calls the underlying Server.Get with the reduced handler chain.
func (m *Mux) Delete(path string, handler HandlerFunc) {

	m.server.Delete(path, m.build(handler))

}
