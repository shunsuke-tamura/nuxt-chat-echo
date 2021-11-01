package main

import (
	"nuxt-echo-chat/backend/route"
	"nuxt-echo-chat/backend/lib"

	"os"
)

func main() {
	lib.DBInit()
	route := route.NewRoute()
	route.Start(":" + os.Getenv("PORT"))
}