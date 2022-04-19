package main

import (
	"fmt"
	"github.com/FallenStarrr/go-projects/go-fiber-crm-basic/lead"
	"github.com/FallenStarrr/go-projects/go-fiber-crm-basic/database"
	"github.com/gofiber/fiber"
	"gorm.io/gorm"
	_ "gihub.com/jinzhu/gorm/dialects/sqllite"
 
)

func setupRoutes(app *fiber.App) {
	app.Get("/api/v1/lead", lead.GetLeads)
	app.Get("/api/v1/lead/:id", lead.GetLead)
	app.Post("/api/v1/lead", lead.NewLead)
	app.Delete("/api/v1/lead/:id", lead.DeleteLead)
}

func initDatabase() {
		var err error
		database.DBConn, err = gorm.Open("sqllite3", "leads.db")
		if err != nil {
			panic("failsed to connect database")
		}

		fmt.Println("Connection opebed to database")
		database.DBConn.AutoMigrate(&lead.Lead())
		fmt.Println("DB migrated")
}

func main() {
 app := fiber.New()
 initDatabase()
 setupRoutes(app)
 app.Listen(3000)
 defer database.DBConn.Close()

}