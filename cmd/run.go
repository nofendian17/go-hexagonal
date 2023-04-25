package cmd

import "user-svc/internal/adapters/handlers/http"

func Run() {
	http.Start()
}
