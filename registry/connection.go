// ******************************************************************************
// Matronator © 2024.                                                          *
// ******************************************************************************

package registry

import (
	"io"
	"net/http"
	"net/url"
	"strconv"
)

const APIURL = `https://www.mtrgen.com/api`

func CreateUser(username string, password string) []byte {
	apiURL := APIURL + "/signup"

	res, err := http.PostForm(apiURL, url.Values{"username": {username}, "password": {password}})

	if err != nil {
		panic(err)
	}

	defer res.Body.Close()

	body, _ := io.ReadAll(res.Body)

	return body
}

func Login(username string, password string, duration ...int) []byte {
	dur := 24
	if len(duration) > 0 {
		dur = duration[0]
	}

	apiURL := APIURL + "/login"

	res, err := http.PostForm(apiURL, url.Values{"username": {username}, "password": {password}, "duration": {strconv.Itoa(dur)}})

	if err != nil {
		panic(err)
	}

	defer res.Body.Close()

	body, _ := io.ReadAll(res.Body)
	return body
	// var object map[string]interface{}
	//
	// _ = json.Unmarshal(body, &object)
	//
	// return object
}
