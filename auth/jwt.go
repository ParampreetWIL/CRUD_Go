package auth

import (
	"encoding/base64"
	"encoding/json"
	"fmt"

	structures "github.com/ParampreetWIL/CRUD_Go/structs"
	"github.com/golang-jwt/jwt/v5"
)

func GenerateJWT(user structures.User, secret_key string) (string, error) {
	var serialized_data, err = json.Marshal(user)

	if err != nil {
		fmt.Println(err)
	}

	claims := jwt.NewWithClaims(jwt.SigningMethodHS512, jwt.MapClaims{
		"sub": serialized_data,
		"iss": "CRUD_Go",
		"aud": "user",
	})

	tokenString, err := claims.SignedString([]byte(secret_key))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func DecryptJWT(tokenString string, secret_key string) (structures.User, error) {
	var user structures.User

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret_key), nil
	})

	if err != nil {
		return user, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		serializedData, ok := claims["sub"].(string)
		if !ok {
			return user, fmt.Errorf("invalid sub claim")
		}

		fmt.Println(base64.StdEncoding.DecodeString(serializedData))
		base64_decoded, err := base64.StdEncoding.DecodeString(serializedData)

		if err != nil {
			fmt.Println(err)
		}

		err = json.Unmarshal([]byte(base64_decoded), &user)
		if err != nil {
			return user, err
		}
		return user, nil
	} else {
		fmt.Println("Hello")
		return user, fmt.Errorf("invalid token")
	}
}
