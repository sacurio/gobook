package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/sacurio/gobook/controller"
	"github.com/sacurio/gobook/response"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Failed to load .env file")
	}

	initRouter()

}

func initRouter() {
	port := os.Getenv("app_port")
	if port == "" {
		port = "8000"
	}

	r := mux.NewRouter()

	fmt.Println("Listening in port: " + port)

	// Defines here the routers needed:
	r.HandleFunc("/api/test/ping", hola)

	http.ListenAndServe(fmt.Sprintf(":%s", port), r)
}

func hola(w http.ResponseWriter, r *http.Request) {
	response.RespondWithJSON(w, http.StatusOK, controller.Ping())
}
