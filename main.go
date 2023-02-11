package main

import (
	"image"
	"image/color"
	"image/png"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func MainHandler(ctx *fiber.Ctx) error {
	width, err := strconv.Atoi(ctx.Params("width"))
	if err != nil {
		return ctx.SendStatus(http.StatusMethodNotAllowed)
	}

	height, err := strconv.Atoi(ctx.Params("height"))
	if err != nil {
		return ctx.SendStatus(http.StatusMethodNotAllowed)
	}

	img := image.NewRGBA(image.Rect(0, 0, width, height))

	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			img.Set(x, y, color.RGBA{
				uint8(x),
				uint8(y),
				0,
				255,
			})
		}
	}

	err = png.Encode(ctx, img)
	if err != nil {
		return ctx.SendStatus(http.StatusInternalServerError)
	}
	return nil
}

func generateImageFile() {
	width := 256
	height := 256
	img := image.NewRGBA(image.Rect(0, 0, width, height))

	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			img.Set(x, y, color.RGBA{
				uint8(x),
				uint8(y),
				0,
				255,
			})
		}
	}

	file, err := os.Create("image.png")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	err = png.Encode(file, img)
	if err != nil {
		panic(err)
	}

}

func main() {
	app := fiber.New()
	app.Get("/image/:width/:height", MainHandler)
	err := app.Listen(":8080")
	if err != nil {
		log.Fatalf("listen server error")
	}
}
