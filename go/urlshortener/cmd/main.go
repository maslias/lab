package main

import (
	"log"

	"github.com/maslias/urlshortener/app"
	"github.com/maslias/urlshortener/configs"
)

func main() {
	app := app.NewApp(configs.Envs.APP_ADDR)
	if err := app.Run(); err != nil {
		log.Fatal(err)
	}
}

