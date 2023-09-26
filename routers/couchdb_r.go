package routers

import (
	"fmt"
	"loadbalancer/controllers"
	"loadbalancer/models"

	"github.com/buaazp/fasthttprouter"
	"github.com/valyala/fasthttp"
)

func CouchDBRouters(router *fasthttprouter.Router) {
	router.PUT("/cdb/create", SendToNextServer)
	router.POST("/lb/:appid/*path", SendToNextServer)

	router.GET("/servers/:appid", GetServerPool)    //Menampilkan pool server berdasarkan appid
	router.POST("/servers/:appid", AddServerHandle) //Menyimpan pool server berdasarkan appid

	router.HandleMethodNotAllowed = false
	router.MethodNotAllowed = func(ctx *fasthttp.RequestCtx) {
		ctx.SetStatusCode(fasthttp.StatusMethodNotAllowed)
		res := models.DefaultResponse{Messeges: "Your method is not allowed", Status: fasthttp.StatusMethodNotAllowed}
		fmt.Fprintf(ctx, controllers.StructToJson(res))
	}
	router.NotFound = func(ctx *fasthttp.RequestCtx) {
		ctx.SetStatusCode(fasthttp.StatusNotFound)
		res := models.DefaultResponse{Messeges: "API is not found", Status: fasthttp.StatusNotFound}
		fmt.Fprintf(ctx, controllers.StructToJson(res))
	}
}
