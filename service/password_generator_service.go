package service

import (
	"math/rand"
	constant "random_password_generator/constant"
	model "random_password_generator/model"
	"time"
)

func GeneratePassword(option model.Option) (string, error) {

	randomSource := rand.NewSource(time.Now().UnixMicro())
	passwordSlice := make([]byte, option.Lenght)
	charset := createCharset(option)

	for i := range passwordSlice {
		randomNumber := randomSource.Int63()
		charsetLen := len(charset)
		passwordSlice[i] = charset[randomNumber%int64(charsetLen)]
	}

	return string(passwordSlice), nil
}

func createCharset(option model.Option) string {
	var charset string = ""

	if option.IncludeLowercase {
		charset += constant.CharsetLowercase
	}
	if option.IncludeUppercase {
		charset += constant.CharsetUppercase
	}
	if option.IncludeSpecials {
		charset += constant.CharsetSpecialLetters
	}
	if option.IncludeNumbers {
		charset += constant.CharsetNumbers
	}
	return charset
}
