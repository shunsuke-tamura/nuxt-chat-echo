package main

import (
	"nuxt-echo-chat/backend/route"
	"os"
)

func main() {
	route := route.NewRoute()
	route.Start(":" + os.Getenv("PORT"))
}