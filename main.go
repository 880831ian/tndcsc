package main

import (
	"tndcsc/routers"
)

func main() {
	r := routers.SetupRouter()
	r.Run(":8080")
}
