package main

import "dev08/internal/handlers"

func main() {

	handler := handlers.NewHandler()
	if err := handler.Run(); err != nil {
		panic(err)
	}
}
