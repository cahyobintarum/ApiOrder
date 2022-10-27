package main

import (
	"apiorder/database"
	_ "apiorder/docs"
	"apiorder/routers"
)

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /
func main() {
	database.StartDB()
	var port = ":8080"
	routers.StartServer().Run(port)
}
