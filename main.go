package main

import (
	"sky-take-out/resources/commonParams"
	"sky-take-out/server/routes"
)

func main() {
	r := routes.SetupRouter()
	r.Run(":" + commonParams.ServerPort)
}
