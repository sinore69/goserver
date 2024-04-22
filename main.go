package main

import (
	"fmt"
	"goserver/cryptography"
	"log"
	"github.com/gofiber/fiber/v2"
)

type info struct {
	Name  string
	Shift int
	Pwd   string
}
type joke struct{
	Value string 
}

func main() {
	app := fiber.New()
	app.Get("/name", name)
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
	encryptedText := cryptography.Encode(data.Name)
	fmt.Println(encryptedText)
	decryptedText,err := cryptography.Decode(encryptedText)
	if err!=nil{
		panic("decoding failed")
	}
	fmt.Println(decryptedText)
	return c.Send(c.Body())
}
