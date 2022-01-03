# random_password_generator

Create a new password of desired lenght.

# Password Create Options
  - Lenght
  - Include Lowercase
  - Include Uppercase
  - Include Specials
  - Include Numbers
  
# Usage
  ```
  import (
	"fmt"
	"random_password_generator/model"
	service "random_password_generator/service"
	)
  
	func main() {

		options := model.Option{
				Lenght: 10,
				IncludeNumbers: true, 
				IncludeSpecials: true,
				IncludeLowercase: true,
				IncludeUppercase: true
			   }

		password, _ := service.GeneratePassword(options)

		fmt.Println(password)
	}

 	 Output : +segW@Qw8e
  ```
