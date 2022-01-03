package main

import (
	"fmt"
	"random_password_generator/model"
	service "random_password_generator/service"
)

func main() {

	options := model.Option{Lenght: 10, IncludeNumbers: true, IncludeLowercase: true, IncludeUppercase: true}

	password, _ := service.GeneratePassword(options)

	fmt.Println(password)
}
