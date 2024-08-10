package main

import (
	"fmt"
	"net/http"

	"github.com/thaiminh2022/calculator_api/internal/handlers"
	"github.com/thaiminh2022/calculator_api/internal/middleware"
)

func main() {
	mux := http.NewServeMux()
	middlewareStack := middleware.CreateStack(
		middleware.Logging,
	)
    handlers.Handler(mux)

	fmt.Println("Serving at port 3000")
	http.ListenAndServe(":3000", middlewareStack(mux))
}
