package main

import "net/http"

func (server *Server) HandleGetIndex(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Server works!"))
}
