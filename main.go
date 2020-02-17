package main

import(
	"github.com/ebot1234/Go-Webserver/dmv"
	"github.com/ebot1234/Go-Webserver/web"
	"log"
	"math/rand"
	"time"
)

const DatabasePath = "./DMV_DB.db"
const httpPort = 8080

//Main entry point
func main(){
	rand.Seed(time.Now().UnixNano())

	dmv, err := dmv.NewDMV(DatabasePath)
	if err != nil {
		log.Fatalln("Error on startup:", err)
	}

	web := web.NewWeb(dmv)
	go web.ServeWebInterface(httpPort)

	//Maybe run a thread here later
}