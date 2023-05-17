package main

import (
	"fmt"
	"os"

	"encoding/json"

	"github.com/golang-jwt/jwt"
)

func main() {

	// claims := decodeClaimsJSON("claims.json")

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, jwt.MapClaims{
		"foo":   "bar",
		"nbf":   "a",
		"hello": "test",
	})

	privatekey, err := os.ReadFile("private.key")
	if err != nil {
		fmt.Println("Could not read private key")
	}

	decoded, err := jwt.ParseRSAPrivateKeyFromPEM([]byte(privatekey))
	if err != nil {
		fmt.Println("Cannot parse RSA private key")
	}

	tokenString, err := token.SignedString(decoded)
	if err != nil {
		fmt.Println("Could not generate token")
	}
	fmt.Println(tokenString)

}

func decodeClaimsJSON(claimsfilepath string) map[string]interface{} {

	claims, err := os.ReadFile(claimsfilepath)
	if err != nil {
		fmt.Println("Claims not found")
	}
	var b map[string]interface{}
	json.Unmarshal([]byte(claims), &b)
	fmt.Println(b)

	return b

}
