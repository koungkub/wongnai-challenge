package main

import "github.com/koungkub/wongnai/internal/route"

func main() {

	r := route.New()

	r.Listen(3000)
}
