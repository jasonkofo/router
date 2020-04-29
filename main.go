package main

import (
	"fmt"

	service "github.com/jasonkofo/router/service"
)

func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	s, err := service.New()
	if err != nil {
		panic(err)
	}
	s.ListenAndServe()
}
