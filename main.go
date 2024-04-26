package main

import (
	"fmt"
	"goserver/cryptography"
	"log"
	"github.com/gofiber/fiber/v2"
)

type info struct {
	Text  string
}

func main() {
	app := fiber.New()
	app.Get("/name", name)
	app.Post("/encode", encoding)
	app.Post("/decode",decoding)
	fmt.Println("server is running on port 5000")
	app.Listen("127.0.0.1:5000")
}
func name(c *fiber.Ctx) error {
	return c.SendString(c.Params("name"))
}
func encoding(c *fiber.Ctx) error {
	data := new(info)
	if err := c.BodyParser(data); err != nil {
		log.Fatal(err)
	}
	encryptedText,err := cryptography.Encode(data.Text)
	if err!=nil{
		panic("encoding failed")
	}
	data.Text=encryptedText
	return c.JSON(data)
}

func decoding(c *fiber.Ctx) error{
	data :=new(info)
	if err:=c.BodyParser(data);err!=nil{
		log.Fatal(err)
	}
	decryptedText,err := cryptography.Decode(data.Text)
	if err!=nil{
		panic("decoding failed")
	}
	data.Text=decryptedText
	return c.JSON(data)
}