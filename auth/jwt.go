package auth

import (
	"encoding/json"
	"fmt"

	structures "github.com/ParampreetWIL/CRUD_Go/structs"
	"github.com/golang-jwt/jwt/v5"
)

func generateJWT(user structures.User) string {
	var serialized_data, err = json.Marshal(user)

	if err != nil {
		fmt.Println(err)
	}

	jwt.NewWithClaims(jwt.SigningMethodHS512, jwt.MapClaims{
		"sub": serialized_data,
	})
	return ""
}
