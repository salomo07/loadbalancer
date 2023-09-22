package config

import (
	"encoding/base64"
)

var REDIS_HOST = "WVhCdU1TMXJaWGt0Wm1sdVkyZ3RNelExTnpZdWRYQnpkR0Z6YUM1cGJ3PT0="
var REDIS_USER = "WkdWbVlYVnNkQT09"
var REDIS_PASS = "TUdFek9EZzJZMkl3TXpZME5EUm1aV0l3WXpVM01UY3dOV0UyWldKa04yST0="
var REDIS_PORT = "TXpRMU56WT0="
var REDIS_CRED = ""

func main() {
	GetCredRedis()
	println(REDIS_CRED)
}
func DecodedCredtial(encoded string) (string, string) {
	decodedText, err := base64.StdEncoding.DecodeString(encoded)
	if err != nil {
		panic(err)
	}
	return string(decodedText), ""
}
func GetCredRedis() string {
	for x := 0; x < 2; x++ {
		res, err := DecodedCredtial(REDIS_USER)
		if err != "" {
			print(err)
		}
		REDIS_USER = res
	}
	for x := 0; x < 2; x++ {
		res, err := DecodedCredtial(REDIS_PASS)
		if err != "" {
			print(err)
		}
		REDIS_PASS = res
	}
	for x := 0; x < 2; x++ {
		res, err := DecodedCredtial(REDIS_HOST)
		if err != "" {
			print(err)
		}
		REDIS_HOST = res
	}

	for x := 0; x < 2; x++ {
		res, err := DecodedCredtial(REDIS_PORT)
		if err != "" {
			print(err)
		}
		REDIS_PORT = res
	}
	REDIS_CRED = "redis://" + REDIS_USER + ":" + REDIS_PASS + "@" + REDIS_HOST + ":" + REDIS_PORT
	return REDIS_CRED
}
