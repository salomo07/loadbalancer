package config

import (
	"encoding/base64"
)

var CDB_HOST = "YUhSMGNEb3ZMMnh2WTJGc2FHOXpkQT09"
var CDB_USER = "WVdSdGFXND0="
var CDB_PASS = "TVRJeg=="
var CDB_PORT = "TlRrNE5BPT0="
var CDB_CRED = ""

var REDIS_HOST = "WVhCdU1TMXJaWGt0Wm1sdVkyZ3RNelExTnpZdWRYQnpkR0Z6YUM1cGJ3PT0="
var REDIS_USER = "WkdWbVlYVnNkQT09"
var REDIS_PASS = "TUdFek9EZzJZMkl3TXpZME5EUm1aV0l3WXpVM01UY3dOV0UyWldKa04yST0="
var REDIS_PORT = "TXpRMU56WT0="
var REDIS_CRED = ""

func DecodedCredtial(encoded string) (string, string) {
	decodedText, err := base64.StdEncoding.DecodeString(encoded)
	if err != nil {
		panic(err)
	}
	return string(decodedText), ""
}
func GetCredRedis() string {
	if REDIS_CRED != "" {
		return REDIS_CRED
	}
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

func GetCredCDB() string {
	print(CDB_CRED)
	if CDB_CRED != "" {
		return CDB_CRED
	}
	for x := 0; x < 2; x++ {
		res, err := DecodedCredtial(CDB_USER)
		if err != "" {
			print(err)
		}
		CDB_USER = res
	}

	for x := 0; x < 2; x++ {

		res, err := DecodedCredtial(CDB_PASS)
		if err != "" {
			print(err)
		}
		CDB_PASS = res
	}
	for x := 0; x < 2; x++ {
		res, err := DecodedCredtial(CDB_HOST)
		if err != "" {
			print(err)
		}
		CDB_HOST = res
	}
	for x := 0; x < 2; x++ {
		res, err := DecodedCredtial(CDB_PORT)
		if err != "" {
			print(err)
		}
		CDB_PORT = res
	}
	CDB_CRED = "https://" + CDB_USER + ":" + CDB_PASS + "@" + CDB_HOST + ":" + CDB_PORT

	return CDB_CRED
}
