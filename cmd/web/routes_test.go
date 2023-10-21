package main

import (
	"testing"

	"github.com/ManojGowda27/bnb-booking/internal/config"
	"github.com/go-chi/chi"
)

func TestRoutes(t *testing.T) {
	// Adding test cases for routes

	var app config.AppConfig

	mux := routes(&app)

	switch v := mux.(type) {
	case *chi.Mux:
		// do nothing; test passed

	default:
		t.Errorf("type is not *chi.Mux, type is %T", v)
	}
}
