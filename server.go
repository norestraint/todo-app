package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "norestraint"
	password = "simplepassword"
	dbname   = "todo"
)

func main() {
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return indexHandler(c, db)
	})

	app.Post("/", func(c *fiber.Ctx) error {
		return indexHandler(c, db)
	})

	app.Put("/update", func(c *fiber.Ctx) error {
		return indexHandler(c, db)
	})

	app.Delete("/delete", func(c *fiber.Ctx) error {
		return indexHandler(c, db)
	}) // Add this

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	log.Fatalln(app.Listen(fmt.Sprintf(":%v", port)))
}

func indexHandler(c *fiber.Ctx, db *sql.DB) error {
	return c.SendString("Hello")
    var res string
    var todos []string

    rows, err := db.Query("SELECT * FROM todos")
    defer rows.Close()

    if err != nil {
        log.Fatalln(err)
        c.JSON("An error occured")
    }

    for rows.Next() {
        rows.Scan(&res)
        todos = append(todos, res)
    }

	return c.Render("index", fiber.Map{
        "Todos": todos,
    })
}

func postHandler(c *fiber.Ctx, db *sql.DB) error {
	return c.SendString("Hello")
}

func putHandler(c *fiber.Ctx, db *sql.DB) error {
	return c.SendString("Hello")
}

func deleteHandler(c *fiber.Ctx, db *sql.DB) error {
	return c.SendString("Hello")
}
