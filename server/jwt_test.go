package main

import (
	//"fmt"
	//"time"
	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"testing"
	"github.com/stretchr/testify/assert"
)

// because the encoded token is really just a random looking string, it is hard to test by itself; all the JWT functions are tested together.
func TestJWT (t *testing.T) {
	
	// the things we will store in our token for a player named "Chanel"
	name := "Chanel"
	userid := uuid.New()
	gameid := "1234"
	isHost := true
	
	// note: we are using the global constant from jwt.go here to encode the JWT token.
	encodedToken, err := newJWT(name, userid, gameid, isHost, []byte(signKey))
	
	// the encoding process shouldn't face any issues.
	assert.Equal(t, nil, err)
	
	// test the parse function!
	validToken, tokenIsValid := parseJWT(encodedToken, signKey)
	
	// the token should be valid, as it wasn't changed by the user.
	assert.Equal(t, true, tokenIsValid)
	
	// the token should be decoded to the contents that we put into it; e.g. same name, userid, gameid, host.
	assert.Equal(t, name, validToken.Claims.(jwt.MapClaims)["name"])
	assert.Equal(t, gameid, validToken.Claims.(jwt.MapClaims)["gameid"])
	assert.Equal(t, isHost, validToken.Claims.(jwt.MapClaims)["isHost"])
	
	// note: the process of putting the uuid into a JWT turns it into a string.
	assert.Equal(t, userid.String(), validToken.Claims.(jwt.MapClaims)["userid"])
	
	// test the claims function!
	validClaims, claimsAreValid := getValidClaims(encodedToken)
	
	// the claims should be valid, as they weren't changed by the user.
	assert.Equal(t, true, claimsAreValid)
	
	assert.Equal(t, name, validClaims["name"])
	assert.Equal(t, gameid, validClaims["gameid"])
	assert.Equal(t, isHost, validClaims["isHost"])
	
	// note: the process of putting the uuid into a JWT turns it into a string.
	assert.Equal(t, userid.String(), validClaims["userid"])
	
	// test the functions with a bad token!
	badEncodedToken := "modify" + encodedToken
	
	// try parsing JWT
	invalidToken, tokenIsInvalid := parseJWT(badEncodedToken, signKey)
	
	// false means invalid
	assert.Equal(t, false, tokenIsInvalid)
	
	// this should be the same as the invalid token: just an empty token!
	var emptyToken *jwt.Token
	
	assert.Equal(t, emptyToken, invalidToken)
	
	// try getting claims from bad token
	invalidClaims, claimsAreInvalid := getValidClaims(badEncodedToken)
	
	// false means invalid
	assert.Equal(t, false, claimsAreInvalid)
	
	// this should be the same as invalid claims: emtpy claims!
	var emptyClaims jwt.MapClaims
	
	assert.Equal(t, emptyClaims, invalidClaims)
	
}