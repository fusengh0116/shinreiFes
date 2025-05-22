package main

import "net/http"

type Ping struct {
}

func (p *Ping) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	_, err := w.Write([]byte("{\"message\":\"pong\"}"))
	if err != nil {
		return
	}
}

func NewPing() *Ping {
	return &Ping{}
}
