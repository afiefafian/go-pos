package helper

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

func HashAndSalt(pwd string) string {
	//salt := viper.GetString(`salt`)
	hash, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.MinCost)
	if err != nil {
		log.Println(err)
	}

	return string(hash)
}
