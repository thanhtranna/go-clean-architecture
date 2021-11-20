package main

import (
	"fmt"
	"log"

	"golang-clean-architecture/config"
	"golang-clean-architecture/infrastructure/datastore"
	"golang-clean-architecture/infrastructure/router"
	"golang-clean-architecture/registry"

	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/labstack/echo/v4"
)

func main() {
	config.ReadConfig()

	db := datastore.NewDB()
	db.LogMode(true)
	defer db.Close()

	r := registry.NewRegistry(db)

	e := echo.New()
	e = router.NewRouter(e, r.NewAppController())

	fmt.Println("Server listen at http://localhost" + ":" + config.Config.Server.Address)
	if err := e.Start(":" + config.Config.Server.Address); err != nil {
		log.Fatalln(err)
	}
}
