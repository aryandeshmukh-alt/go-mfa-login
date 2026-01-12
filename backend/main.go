package main

import (
	"log"
	"net/http"

	"go-mfa-login/handlers"
)

func withCORS(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		h.ServeHTTP(w, r)
	})
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/register", handlers.Register)
	mux.HandleFunc("/login", handlers.Login)
	mux.HandleFunc("/verify-otp", handlers.VerifyOTP)

	log.Println("Backend running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", withCORS(mux)))
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

}
