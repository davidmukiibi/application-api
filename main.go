package main

import (
	"log"
	"net/http"
	"os"

	"github.com/davidmukiibi/controllers"
	"github.com/davidmukiibi/routes"
	"github.com/davidmukiibi/services"

	"github.com/gorilla/handlers"
)

func Migrate() {
	db, _ := services.DbConnect()
	defer db.Close()
	if err := db.AutoMigrate(&controllers.User{}).Error; err != nil {
		log.Fatalln("Error migrating the database ", err.Error())
	} else {
		log.Println("Migration successful...")
	}
}

// init is going to have the DB connections and any one-time tasks
func init() {
	Migrate()
}

// Define HTTP request routes
func main() {
	router := routes.InitRoutes()
	log.Fatal(http.ListenAndServe(":3000", handlers.LoggingHandler(os.Stdout, router)))
}





























// var testCases = []struct {
// 	FirstName string
// 	Surname   string
// 	UserEmail string
// 	Password  string
// }{
// 	{
// 		"David",
// 		"Mukiibi",
// 		"david@gmail.com",
// 		"1234567890",
// 	},
// }

// var testCasesSlice = []string {
// 	"David",
// }

// func main() {
//     start := time.Now()

// 	for _, eachguy := range testCases {
// 		fmt.Println(eachguy.FirstName)
// 	}

//     elapsed := time.Since(start)
//     log.Printf("Map took %s", elapsed)
// }

// func main() {
//     start := time.Now()

// 	for _, name := range testCasesSlice {
// 					fmt.Println(string(name[:5]))
// 				}

//     elapsed := time.Since(start)
//     log.Printf("Slice took %s", elapsed)
// }