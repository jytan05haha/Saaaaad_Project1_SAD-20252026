package main

import (
	"backend/pkg/handler"
	"net/http"
)

// Handler is the serverless entrypoint required by Vercel. It delegates to the reusable handler package.
func Handler(w http.ResponseWriter, r *http.Request) {
	handler.Handler(w, r)
}

func init() {
	// Attempt to connect to DB when running on platforms that provide DATABASE_URL.
	_ = handler.ConnectDB()
}
