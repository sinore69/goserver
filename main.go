package main

import (
	"encoding/json"
	"fmt"
	"goserver/cryptography"
	"io"
	"log"
	"net/http"

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
	app.Get("/jokes", jokes)
	fmt.Println("server is running on port 5000")
	app.Listen("127.0.0.1:5000")
}
func name(c *fiber.Ctx) error {
	return c.SendString(c.Params("name"))
}
func jokes(c *fiber.Ctx) error {
	res,err:=http.Get("https://api.chucknorris.io/jokes/random")
	if err!=nil{
		log.Fatal("crash")
	}
	body,err:=io.ReadAll(res.Body)
	if err!=nil{
		log.Fatal("crash")
	}
	var data joke
	err=json.Unmarshal(body,&data)
	if err!=nil{
		log.Fatal(err.Error())
	}
	return c.SendString(data.Value)
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
