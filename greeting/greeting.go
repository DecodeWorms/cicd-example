package greeting

import (
	"fmt"

	"github.com/DecodeWorms/cicd-example/models"
)

func Greeting() string {
	us := models.User{
		Name:    "Bola",
		Age:     22,
		Address: "Adewole Estate",
	}
	return fmt.Sprintf("Hello Mr %s of Age %d and at the Address %s", us.Name, us.Age, us.Address)
}
