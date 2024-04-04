package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"goserver/cryptography"
	"log"
)

type info struct {
	Name  string `json:"name"`
	Shift int    `json:"shift"`
}

func main() {
	app := fiber.New()

	app.Get("/:name", name)
	app.Post("/api", post)

	fmt.Println("server is running on port 5000")
	app.Listen("127.0.0.1:5000")
}
func name(c *fiber.Ctx) error {
	return c.SendString(c.Params("name"))
}

func post(c *fiber.Ctx) error {
	data := new(info)
	if err := c.BodyParser(data); err != nil {
		log.Fatal(err)
	}
	encryptedText := cryptography.Encrypt(data.Name, data.Shift)
	fmt.Println(encryptedText)
	decryptedText := cryptography.Decrypt(encryptedText, data.Shift)
	fmt.Println(decryptedText)
	return c.Send(c.Body())
}
