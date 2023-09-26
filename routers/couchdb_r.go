package routers

import (
	"fmt"
	"loadbalancer/controllers"
	"loadbalancer/models"

	"github.com/buaazp/fasthttprouter"
	"github.com/valyala/fasthttp"
)

func CouchDBRouters(router *fasthttprouter.Router) {
	router.POST("/cdb/create_user", controllers.CreateUserCDB)

	router.GET("/gateway/register", GetServerPool)  //Menampilkan pool server berdasarkan appid
	router.POST("/gateway/:appid", AddServerHandle) //Menyimpan pool server berdasarkan appid

	router.HandleMethodNotAllowed = false
	router.MethodNotAllowed = func(ctx *fasthttp.RequestCtx) {
		ctx.SetStatusCode(fasthttp.StatusMethodNotAllowed)
		ctx.Response.Header.Set("Content-Type", "application/json")
		res := models.DefaultResponse{Messeges: "Your method is not allowed", Status: fasthttp.StatusMethodNotAllowed}
		fmt.Fprintf(ctx, controllers.StructToJson(res))
	}
	router.NotFound = func(ctx *fasthttp.RequestCtx) {
		ctx.SetStatusCode(fasthttp.StatusNotFound)
		ctx.Response.Header.Set("Content-Type", "application/json")
		res := models.DefaultResponse{Messeges: "API is not found", Status: fasthttp.StatusNotFound}
		fmt.Fprintf(ctx, controllers.StructToJson(res))
	}
}
