package main

import (
	"log"
	"net/http"

	"prodigy-program/db"
	"prodigy-program/repos"
	"prodigy-program/service"
	"prodigy-program/transport/user"
	"prodigy-program/transport/weekPlan"
)

func main() {
	db.InitDB()

	userRepo := repos.NewUserRepository(db.DB)
	userService := service.NewUserService(userRepo)

	weekPlanRepo := repos.NewWeekPlanRepository(db.DB)
	weekPlanService := service.NewWeekPlanService(weekPlanRepo)

	// Create the HTTP handler
	userHandler := user.NewHTTPHandler(userService)
	weekPlanHandler := weekPlan.NewWeekPlanHandler(weekPlanService)




	// Create a main router (mux) to combine handlers
	mux := http.NewServeMux()
	mux.Handle("/registerUser", userHandler)
	mux.Handle("/weekplan", weekPlanHandler)
	mux.Handle("/getweekplan", weekPlanHandler)

	// Start the server
	log.Println("Server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
