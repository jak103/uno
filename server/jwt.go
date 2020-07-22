package main

import (
	"fmt"
	"time"
	"github.com/dgrijalva/jwt-go"
)

const signKey = "^s@m&R@n&om,St)("

func exampleUsage() {
    createdToken, err := newJWT("Thomas", 1234, []byte(signKey))
    if err != nil {
        fmt.Println("Creating token failed")
    }
    
	token, valid := parseJWT(createdToken, signKey)
	
	if valid {
		fmt.Println(token.Claims.(jwt.MapClaims)["name"])
		fmt.Println(token.Claims.(jwt.MapClaims)["gameid"])
	}
	
}

func newJWT(name string, gameid int, signKey []byte) (string, error) {
    // Create the token
    token := jwt.New(jwt.SigningMethodHS256)
	
    // Set some claims
    token.Claims = jwt.MapClaims{
        "exp": time.Now().Add(time.Hour * 72).Unix(),
        "iat": time.Now().Unix(),
		"name": name,
		"gameid": gameid,
    }
    // Sign and get the complete encoded token as a string
    tokenString, err := token.SignedString(signKey)
    return tokenString, err
}

func parseJWT(myToken string, signKey string) (*jwt.Token, bool) {
    token, err := jwt.Parse(myToken, func(token *jwt.Token) (interface{}, error) {
        return []byte(signKey), nil
    })

    if err == nil && token.Valid {
		return token, true
    } else {
		var i *jwt.Token
		return i, false
    }
}