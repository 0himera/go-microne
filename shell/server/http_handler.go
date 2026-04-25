package main

import (
	"net/http"

	"github.com/0himera/go-microne/internal/infra"
	"github.com/0himera/go-microne/internal/protocol/httpapi"
)

func buildHTTPHandler() http.Handler {
	idGen:= infra.UUIDGenerator{}

	mux:= http.NewServeMux()

	var handler http.Handler = mux
	handler = httpapi.RequestIDMiddleware(idGen)(handler)

	return handler
}

