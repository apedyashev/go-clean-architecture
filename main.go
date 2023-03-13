package main

import (
	"fmt"
	"log"
	"posts/infrastructure/datastore"
	"posts/infrastructure/router"
	"posts/registry"

	"github.com/labstack/echo"
)

func main() {
	db := datastore.NewDB()

	r := registry.NewRegistry(db)

	// Create controllers and routers to hanlde the requests.
	e := echo.New()
	e = router.NewRouter(e, r.NewAppController())

	fmt.Println("Server listen at http://localhost:3005")
	if err := e.Start(":3005"); err != nil {
		log.Fatalln(err)
	}
}
