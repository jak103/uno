package main

import (
	//"fmt"
	"time"
    "strings"
	"github.com/dgrijalva/jwt-go"
    "github.com/jak103/uno/db"
    "github.com/jak103/uno/model"
	//"github.com/google/uuid"
)

// we can have this signkey be whatever we want.
const signKey = "^s@m&R@n&om,St)("

// example usage. Change the name of the function to main and type "go run jwt.go" in this directory to see it working.
// DO NOT commit after renaming this fucntion to main, as this will break the build.
/*
func exampleUsage() {
	// example of originally creating a token
	// a token like this should be set to their header. It is just a string, so it is easy enough to set to a cookie.
    createdToken, err := newJWT("Thomas", uuid.New())
    if err != nil {
        fmt.Println("Creating token failed")
    }

	// this is how we work with tokens that are passed back to us.

	// this is the simple, pretty looking way using the helper function I wrote.
	claims, validClaims := getValidClaims(createdToken)

	if validClaims {
		fmt.Println(claims["name"])
		fmt.Println(claims["userid"])
		fmt.Println(claims["iat"])
		fmt.Println(claims["exp"])
	}

	// techincally, you can just parse the token, but token.Claims.(jwt.MapClaims) looks ugly and is long to type
	// I think we will mostly use getValidClaims, but if for whatever reason you need the whole token, this is how you would do it and access the claims.
	token, valid := parseJWT(createdToken, signKey)

	if valid {
		fmt.Println(token.Claims.(jwt.MapClaims)["name"])
		fmt.Println(token.Claims.(jwt.MapClaims)["gameid"])
	}

}
*/

// function to create a new jwt based on a name, gameid, and a signKey
// note, when this is merged with the db branch, we may want to combine "name" and "userid" into one "player" object
func newJWT(name string, userid string) (string, error) {
	// Create the token
	token := jwt.New(jwt.SigningMethodHS256)

	// Set some claims
	token.Claims = jwt.MapClaims {
		// this 72 should make it expire in 72 hours, which is reasonable (a game of uno shouldn't take longer than 3 days)
		"exp": time.Now().Add(time.Hour * 72).Unix(),
		// this iat is the "iniated at time"
		"iat": time.Now().Unix(),
		// we store here the username, userid, and gameid
		"name":   name,
		"userid": userid,
	}
	// Sign and get the complete encoded token as a string
	tokenString, err := token.SignedString( []byte(signKey) )
	return tokenString, err
}

// function to parse jwt encoded string with a given signkey
func parseJWT(myToken string, signKey string) (*jwt.Token, bool) {

	// parse the token
	token, err := jwt.Parse(myToken, func(token *jwt.Token) (interface{}, error) {
		return []byte(signKey), nil
	})

	if err == nil && token.Valid {
		// no errors, and the token is valid
		// return the parsed token, and a flag that the token is valid and parsable
		return token, true
	} else {
		// either there was an error (couldn't parse), or the token is invalid/expired.
		// create an empty token, and return it with an "invalid" flag.
		var t *jwt.Token
		return t, false
	}
}

// function to parse jwt encoded claims with a given signkey; helper function for getting claims
func getValidClaims(myToken string) (jwt.MapClaims, bool) {

	// get the token, and see if it is valid
	token, valid := parseJWT(myToken, signKey)

	// set up an empty claims
	var claims jwt.MapClaims

	if valid {
		// if the token is valid, get the claims
		claims = token.Claims.(jwt.MapClaims)
	}

	// return the claims (empty if invalid), and a flag indicating if the claims (token) are valid
	return claims, valid

}

// function that gets JWT from auth header
func getValidClaimsFromHeader(authHeader string) (jwt.MapClaims, bool) {
    
	if authHeader == "" {
        // no authorization at all... obviously not authorized
        var c jwt.MapClaims
        return c, false
    }
    
    bearerAndToken := strings.Fields(authHeader) // we should get ["bearer" "encodedToken"]
    
    //authType := bearerAndToken[0] // I don't think we actually need this...
    
    encodedToken := bearerAndToken[1]
    
	claims, valid := getValidClaims(encodedToken)
    
	// return the claims (empty if invalid), and a flag indicating if the claims (token) are valid
	return claims, valid

}

// function that simply creates a JWT payload.
func makeJWTPayload(EncodedJWT string) map[string]interface{} {
	payload := make(map[string]interface{})
    payload["JWT"] = EncodedJWT

	return payload
}

func getPlayerFromHeader(authHeader string) (*model.Player, bool, error){
    database, err := db.GetDb()
	if err != nil {
		return nil, false, err
	}
    
    claims, validUser := getValidClaimsFromHeader(authHeader)
    
    if !validUser {
        return nil, false, err
    }
    
    player, err := database.LookupPlayer(claims["userid"].(string))

	if err != nil {
		return nil, false, err
	}
    
    return player, validUser, err
    
}