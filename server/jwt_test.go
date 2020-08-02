package main

import (
	//"fmt"
	//"time"
	"github.com/dgrijalva/jwt-go"
	//"github.com/google/uuid"
	"testing"
	"github.com/stretchr/testify/assert"
)

// because the encoded token is really just a random looking string, it is hard to test by itself; all the JWT functions are tested together.
func TestJWT (t *testing.T) {
    
	// the things we will store in our token for a player named "Chanel"
	name := "Chanel"
	userid := "someUniqueID"
    
    // this should be the same as the invalid token: just an empty token!
	var emptyToken *jwt.Token
    // this should be the same as invalid claims: emtpy claims!
	var emptyClaims jwt.MapClaims
	
	// note: we are using the global constant from jwt.go here to encode the JWT token.
	encodedToken, err := newJWT(name, userid)
	
	// the encoding process shouldn't face any issues.
	assert.Equal(t, nil, err)
	
	// test the parse function!
	validToken, tokenIsValid := parseJWT(encodedToken, signKey)
	
	// the token should be valid, as it wasn't changed by the user.
	assert.Equal(t, true, tokenIsValid)
	
	// the token should be decoded to the contents that we put into it; e.g. same name, userid, gameid, host.
	assert.Equal(t, name, validToken.Claims.(jwt.MapClaims)["name"])
	assert.Equal(t, userid, validToken.Claims.(jwt.MapClaims)["userid"])
	
	// test the claims function!
	validClaims, claimsAreValid := getValidClaims(encodedToken)
	
	// the claims should be valid, as they weren't changed by the user.
	assert.Equal(t, true, claimsAreValid)
	
	assert.Equal(t, name, validClaims["name"])
	assert.Equal(t, userid, validClaims["userid"])
	
    // this should also work with an authHeader bearing the token
    
    // first, blank auth header should be bad
    authHeader := ""
    
    emptyHeaderClaims, emptyHeaderIsInvalid := getValidClaimsFromHeader(authHeader)
    
    assert.Equal(t, false, emptyHeaderIsInvalid)
	assert.Equal(t, emptyClaims, emptyHeaderClaims)
    
    // simulate a good authHeader
    authHeader = "bearer " + encodedToken
    
    validHeaderClaims, headerIsValid := getValidClaimsFromHeader(authHeader)
    
    assert.Equal(t, true, headerIsValid)
    
    assert.Equal(t, name, validHeaderClaims["name"])
    
	// test the functions with a bad token!
	badEncodedToken := "modify" + encodedToken
	
	// try parsing JWT
	invalidToken, tokenIsInvalid := parseJWT(badEncodedToken, signKey)
	
	// false means invalid
	assert.Equal(t, false, tokenIsInvalid)
	// should be empty when invalid
	assert.Equal(t, emptyToken, invalidToken)
	
	// try getting claims from bad token
	invalidClaims, claimsAreInvalid := getValidClaims(badEncodedToken)
	
	// false means invalid
	assert.Equal(t, false, claimsAreInvalid)
	// should be empty when invalid
	assert.Equal(t, emptyClaims, invalidClaims)
    
    // bad auth header simulation
    authHeader = "bearer " + badEncodedToken
    
    // trying to get claims from bad token in header
    badHeaderClaims, badHeaderIsInvalid := getValidClaimsFromHeader(authHeader)
    
    // false is invalid
    assert.Equal(t, false, badHeaderIsInvalid)
    // should be empty when invalid
	assert.Equal(t, emptyClaims, badHeaderClaims)
	
}

func TestMakeJWTPayload(t *testing.T){
    token := "test"
    payload := makeJWTPayload(token)
    
    assert.Equal(t, token, payload["JWT"])
}