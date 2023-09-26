package routers

import (
	"fmt"
	"loadbalancer/controllers"
	"loadbalancer/models"
	"loadbalancer/services"
	"strings"
	"time"

	"github.com/buaazp/fasthttprouter"
	"github.com/valyala/fasthttp"
)

func LoadBalancerRouters(router *fasthttprouter.Router) {
	router.GET("/lb/:appid/*path", SendToNextServer)
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
func AddServerHandle(ctx *fasthttp.RequestCtx) {
	appid := ctx.UserValue("appid").(string)
	ctx.Response.Header.Set("Content-Type", "application/json")
	if appid == "" {
		ctx.SetStatusCode(fasthttp.StatusNotFound)
		resJson := models.DefaultResponse{Messeges: "AppId is not found", Status: fasthttp.StatusNotFound}
		fmt.Fprintf(ctx, controllers.StructToJson(resJson))
	} else {
		controllers.SaveServers(appid, string(ctx.Request.Body()))
		fmt.Fprintf(ctx, appid+" was saved")
	}
}
func GetServerPool(ctx *fasthttp.RequestCtx) {
	appid := ctx.UserValue("appid").(string)
	ctx.Response.Header.Set("Content-Type", "application/json")
	res := services.GetValueRedis(appid)
	if res == "" {
		ctx.SetStatusCode(fasthttp.StatusNotFound)
		resJson := models.DefaultResponse{Messeges: "AppId is not found", Status: fasthttp.StatusNotFound}
		fmt.Fprintf(ctx, controllers.StructToJson(resJson))
	} else {
		fmt.Fprintf(ctx, res)
	}
}

func SendToNextServer(ctx *fasthttp.RequestCtx) {
	appid := ctx.UserValue("appid").(string)
	path := strings.Replace(string(ctx.Path()), "/lb/"+appid+"/", "", -1)
	fastestServer, isPathSaved, poolServers := controllers.GetFastestServer(appid, path)
	if fastestServer == "" {
		res := models.DefaultResponse{Status: fasthttp.StatusGatewayTimeout, Messeges: "Request cant send forward : Next Server Address is not found"}
		fmt.Fprintf(ctx, controllers.StructToJson(res))
	}
	print(string(ctx.Method()) + " : " + fastestServer + path)

	client := &fasthttp.Client{}
	forwardedRequest := fasthttp.AcquireRequest()
	forwardedRequest.SetRequestURI(fastestServer + path)
	forwardedRequest.SetBody(ctx.Request.Body())
	forwardedRequest.Header.SetMethod(string(ctx.Method()))

	forwardedResponse := fasthttp.AcquireResponse()
	startTime := time.Now()
	err := client.Do(forwardedRequest, forwardedResponse)

	ctx.Response.Header.Set("Content-Type", "application/json")
	ctx.Response.SetStatusCode(forwardedResponse.StatusCode())
	ctx.Response.SetBody(forwardedResponse.Body())
	endTime := time.Now()

	var responseTime int64
	responseTime = endTime.Sub(startTime).Nanoseconds()
	if err != nil {
		print("\033[31m" + "\n\nError : " + err.Error() + "\n\n" + "\033[0m")
		responseTime = 999999999999999999
	}

	if isPathSaved == false {
		isOnline := false
		if err == nil {
			isOnline = true
		} else {
			isOnline = false
		}
		go controllers.SaveResponseTime(appid, poolServers, models.Server{Address: fastestServer, Path: path, Online: isOnline, Latency: responseTime, LastUpdate: time.Now().Format("2006-01-02 15:04:05")})
	}

	fasthttp.ReleaseRequest(forwardedRequest)
	fasthttp.ReleaseResponse(forwardedResponse)
	res := models.DefaultResponse{Status: fasthttp.StatusInternalServerError, Messeges: "Request cant send forward : " + err.Error()}
	fmt.Fprintf(ctx, controllers.StructToJson(res))
}
