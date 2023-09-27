package controllers

import (
	"loadbalancer/config"
	"loadbalancer/models"

	"github.com/valyala/fasthttp"
	"golang.org/x/crypto/bcrypt"
)

func HashingBcrypt(password string) string {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		println("Error:", err)
		return ""
	}
	return string(hashedPassword)
}
func CreateUserCDB(ctx *fasthttp.RequestCtx) {
	print("CreateUserCDB")
	forwardedRequest := fasthttp.AcquireRequest()
	println(config.GetCredCDB() + "/_users")
	forwardedRequest.SetRequestURI(config.GetCredCDB() + "/_users")
	forwardedRequest.SetBody(ctx.Request.Body())
	forwardedRequest.Header.SetMethod("POST")
	forwardedResponse := fasthttp.AcquireResponse()
	client := &fasthttp.Client{}
	err := client.Do(forwardedRequest, forwardedResponse)
	if err != nil {
		res := models.DefaultResponse{Status: fasthttp.StatusInternalServerError, Messeges: "Request cant send forward : " + err.Error()}
		print(StructToJson(res))
		print(err.Error())
	}
	ctx.Response.Header.Set("Content-Type", "application/json")
	ctx.Response.SetStatusCode(forwardedResponse.StatusCode())
	ctx.Response.SetBody(forwardedResponse.Body())

	ctx.Response.SetStatusCode(forwardedResponse.StatusCode())
	ctx.Response.SetBody(forwardedResponse.Body())
	fasthttp.ReleaseRequest(forwardedRequest)
	fasthttp.ReleaseResponse(forwardedResponse)
}
