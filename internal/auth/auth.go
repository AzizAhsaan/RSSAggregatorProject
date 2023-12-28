package auth

import (
	"net/http"
	"strings"
	"errors"
)

// GetApikey extracts an api key from the headers of an http request
// example :
//Authorization : APIKEY {insert apikey here}
func GetAPIKey(headers http.Header) (string,error){
	val := headers.Get("Authorization")
	if val == ""{
		return "", errors.New("No authntication info found")
	}
	vals := strings.Split(val," ")
	if len(vals) != 2{
		return "", errors.New("Malformed auth header")
	}
	if vals[0] != "ApiKey"{
		return "", errors.New("Malformed first part of auth header")
	}
	return vals[1],nil
}