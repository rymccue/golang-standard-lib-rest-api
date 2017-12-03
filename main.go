package main

import (
	"log"
	"net/http"
	"os"

	"github.com/rymccue/golang-standard-lib-rest-api/controllers"
	"github.com/rymccue/golang-standard-lib-rest-api/routes"
	"github.com/rymccue/golang-standard-lib-rest-api/utils/caching"
	"github.com/rymccue/golang-standard-lib-rest-api/utils/database"
)

func main() {
	db, err := database.Connect(os.Getenv("PGUSER"), os.Getenv("PGPASS"), os.Getenv("PGDB"), os.Getenv("PGHOST"), os.Getenv("PGPORT"))
	if err != nil {
		log.Fatal(err)
	}
	cache := &caching.Redis{
		Client: caching.Connect(os.Getenv("REDIS_ADDR"), os.Getenv("REDIS_PASSWORD"), 0),
	}

	userController := controllers.NewUserController(db, cache)
	jobController := controllers.NewJobController(db, cache)

	mux := http.NewServeMux()
	routes.CreateRoutes(mux, userController, jobController)

	if err := http.ListenAndServe(":8000", mux); err != nil {
		log.Fatal(err)
	}
}
