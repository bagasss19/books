package main

import (
	"log"
	"net/http"
	"os"
	"runtime"

	"github.com/books/config"
	"github.com/books/docs"
	bookController "github.com/books/service/books/controller"
	bookFeature "github.com/books/service/books/feature"
	bookRepo "github.com/books/service/books/repository"
	"github.com/books/utils"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
	httpSwagger "github.com/swaggo/http-swagger/v2"
)

func readEnvironmentFile() {
	//Environment file Load --------------------------------
	err := godotenv.Load(".secret.env")

	if err != nil {
		log.Println("Error loading .env file")
	}

	mode := os.Getenv("MODE")

	switch mode {
	case "development":
		docs.SwaggerInfo.Host = "localhost:3000"
		docs.SwaggerInfo.Schemes = []string{"http"}
	case "staging":
		docs.SwaggerInfo.Host = "books.staging"
		docs.SwaggerInfo.Schemes = []string{"https"}
	case "production":
		docs.SwaggerInfo.Host = "books.prod"
		docs.SwaggerInfo.Schemes = []string{"https"}
	}
}

// @title           Books API
// @version         1.0
// @description     This is a collection of books API.

// @host      localhost:9000
// @BasePath  /api
func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	var (
		port = 3000
	)

	forever := make(chan struct{})

	go func() {
		readEnvironmentFile()
		DBConn, err := config.DBConnect()
		if err != nil {
			log.Fatalf("Database connection error: %s", err)
		}

		// Init Repository
		bookRepository := bookRepo.NewBookRepository(DBConn)

		// Init Feature
		bookFeat := bookFeature.NewBookFeature(bookRepository)

		// Init Controller
		bookCont := bookController.NewBookController(&utils.App{}, bookFeat)

		r := chi.NewRouter()
		cr := cors.New(cors.Options{
			AllowedOrigins: []string{"*"},
			AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
			AllowedHeaders: []string{
				"Accept",
				"Authorization",
				"Content-Type",
				"X-CSRF-Token",
				"X-SIGNATURE",
				"X-TIMESTAMPT",
				"X-CHANNEL",
				"X-PLAYER",
				"Access-Control-Allow-Headers",
				"X-Requested-With",
				"application/json",
				"Cache-Control",
				"Token",
				"X-Token",
			},
			ExposedHeaders:   []string{"Link"},
			AllowCredentials: true,
			MaxAge:           300,
		})
		r.Use(cr.Handler)
		r.Use(middleware.Logger)
		r.Get("/", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("welcome"))
		})

		r.Route("/api", func(r chi.Router) {
			r.Route("/books", func(r chi.Router) {
				r.Get("/", bookCont.GetAllBook)
				r.Get("/{id}", bookCont.GetOneBook)
				r.Post("/", bookCont.CreateBook)
				r.Put("/{id}", bookCont.UpdateBook)
				r.Delete("/{id}", bookCont.DeleteBook)
			})
		})

		r.Get("/swagger/*", httpSwagger.Handler(
			httpSwagger.URL("http://localhost:3000/swagger/doc.json"), //The url pointing to API definition
		))

		log.Printf("Starting up on http://localhost:%d", port)
		http.ListenAndServe(":3000", r)
	}()
	<-forever

}
