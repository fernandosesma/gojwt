package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
)

func main() {

	env := "debug"
	keypath, err := FetchKeypath(env)
	if err != nil {
		fmt.Println("Cannot get private key, check environment name")
	}

	claims := WriteClaimsFromJSON("claims.json")

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)

	privatekey, err := os.ReadFile(keypath)
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

func WriteClaimsFromJSON(JsonFilePath string) *jwt.MapClaims {

	data, err := os.ReadFile(JsonFilePath)
	if err != nil {
		fmt.Println("Cannot parse JSON file")
	}

	// Map container to decode JSON structure into
	c := make(map[string]interface{})

	// Unmarshal JSON and set standard claims
	e := json.Unmarshal(data, &c)
	SetStandardClaims(c)

	if e != nil {
		panic(e)
	}

	claims := jwt.MapClaims(c)
	return &claims

}

func SetStandardClaims(DecodedJsonStruct map[string]interface{}) {
	DecodedJsonStruct["jti"] = uuid.NewString()
	DecodedJsonStruct["iat"] = time.Now().Unix()
	DecodedJsonStruct["exp"] = time.Now().Add(time.Minute * 20).Unix()
}

func FetchKeypath(env string) (string, error) {
	absFilepath := filepath.Join("secrets", env, "private.key")
	return filepath.Abs(absFilepath)
}
