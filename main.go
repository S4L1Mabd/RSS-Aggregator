package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
)

func main() {

	// Routers Configuration

	router := chi.NewRouter() //Creating new Router

	router.Use(cors.Handler(cors.Options{ // give permission to browser to reserve any request

		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}),
	)

	routerv1 := chi.NewRouter()

	routerv1.HandleFunc("/blida", handlerReader) // the fucntion handlerReader is excuting in /routerv1/blida with all type of request
	//  if we did routerv1.GET("/blida", handlerReader)  we can just do GET Request in localhost:8080/v1/blida

	routerv1.Get("/err", ErrorHandler)

	router.Mount("/v1", routerv1) /// Hadi m3ntha bli pour RbaTna Routerv1 ba Router fa :   /localhost:8080/v1

	// end of routers configuration

	godotenv.Load(".env")

	PortValue := os.Getenv("PORT")

	if PortValue == "" {

		log.Fatal("The Serveur is Down , No port Match ")
	}

	srv := &http.Server{ //Connectimg our router to http Serveur

		Handler: router,
		Addr:    ":" + PortValue,
	}

	msg := fmt.Sprintf("the Server is running on Port %s ........", PortValue)

	fmt.Printf(msg)

	err := srv.ListenAndServe()

	if err != nil {

		log.Fatal("The Serveur is Down , No port Match ")
	}

}