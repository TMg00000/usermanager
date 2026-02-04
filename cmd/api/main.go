package main

import (
	"context"
	"log"
	"net/http"
	"time"
	"usermanager/internal/configs"
	"usermanager/internal/http/handler"
	"usermanager/internal/repository"
	"usermanager/internal/repository/connection/mongoconnection"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

func main() {
	errConfigs := configs.StartConfigs()
	returnFatalError(errConfigs)

	col := StartDataBase()
	defer func() {
		if err := col.Database().Client().Disconnect(context.Background()); err != nil {
			log.Fatal(err)
		}
	}()

	repository.UniqueKeyEmail(col)

	UserRepo := repository.NewUserManagerRepository(col)
	controller := handler.UsersManagerServices{
		Services: UserRepo,
	}

	r := mux.NewRouter()
	r.HandleFunc("/api/user/signup", controller.RegisterNewUser).Methods("POST")
	r.HandleFunc("/api/user/login", controller.LoginUser).Methods("POST")
	r.HandleFunc("/api/user/allusers", controller.AllUsers).Methods("GET")

	log.Println("Server is running on port 9437")
	log.Println("Server is running in address http://localhost:9437")
	returnFatalError(http.ListenAndServe(":9437", r))
}

func StartDataBase() *mongo.Collection {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongoconnection.NewMongoConnection(ctx)
	returnFatalError(err)

	return client.Database(configs.Env.MongoDB).Collection(configs.Env.MongoCol)
}

func returnFatalError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
