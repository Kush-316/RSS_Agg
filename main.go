package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
	"github.com/kush-316/RSS_Agg/internal/database"
	_ "github.com/lib/pq"
)

type apiConfig struct {
	DB *database.Queries
}
func main(){
	fmt.Println("hello world")

	godotenv.Load(".env")
	portString := os.Getenv("PORT")
	if portString ==""{
		log.Fatal("PORT not found in the environment")
	}
	dbURL := os.Getenv("DB_URL")
	if dbURL ==""{
		log.Fatal("DB_URL not found in the environment")
	}

	conn,err1:= sql.Open("postgres", dbURL)
	if err1!= nil{
		log.Fatal("Can't connect to database:", err1)
	}

	queries:=database.New(conn)
	apiCfg := apiConfig{
		DB:queries,
	}
	router := chi.NewRouter()

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"https://*", "http://*"},
		AllowedMethods: []string{"GET","POST","PUT","DELETE","OPTIONS"},
		AllowedHeaders: []string{"*"},
		ExposedHeaders: []string{"Link"},
		AllowCredentials: false,
		MaxAge: 300,
	}))

	v1Router := chi.NewRouter()
	v1Router.Get("/healthz", handlerReadiness)
	v1Router.Get("/err", handlerErr)
	v1Router.Post("/users", apiCfg.handlerCreateUser)
	v1Router.Get("/users", apiCfg.middlewareAuth(apiCfg.handlerGetUser))

	v1Router.Post("/feeds", apiCfg.middlewareAuth(apiCfg.handlerCreateFeed))
	v1Router.Get("/feeds", apiCfg.handlerGetFeeds)

	router.Mount("/v1", v1Router)
	srv := &http.Server{
		Handler: router,
		Addr: ":" + portString,
	}

	log.Printf("server starting on port %v", portString)
	err := srv.ListenAndServe()
	//fmt.Println("Port:", portString)

	if err!=nil {
		log.Fatal(err)
	}
}