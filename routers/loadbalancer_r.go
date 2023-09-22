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
	// @Summary Get a list of items
	// @Description Get a list server of application
	// @ID get-item-list
	// @Accept  json
	// @Produce  json
	// @Success 200 {object} []models.Server
	// @Router /servers/:appid/ [get]
	router.GET("/servers/:appid", GetServerPool)    //Menampilkan pool server berdasarkan appid
	router.POST("/servers/:appid", AddServerHandle) //Menyimpan pool server berdasarkan appid

	router.HandleMethodNotAllowed = false
	router.MethodNotAllowed = func(ctx *fasthttp.RequestCtx) {
		ctx.SetStatusCode(fasthttp.StatusMethodNotAllowed)
		jsonResponse := `{"messege": "Your method is not allowed"}`
		fmt.Fprintf(ctx, jsonResponse)
	}
	router.NotFound = func(ctx *fasthttp.RequestCtx) {
		ctx.SetStatusCode(fasthttp.StatusNotFound)
		jsonResponse := `{"messege": "API is not found"}`
		fmt.Fprintf(ctx, jsonResponse)
	}
}
func AddServerHandle(ctx *fasthttp.RequestCtx) {
	appid := ctx.UserValue("appid").(string)
	ctx.Response.Header.Set("Content-Type", "application/json")
	if appid == "" {
		ctx.SetStatusCode(fasthttp.StatusNotFound)
		fmt.Fprintf(ctx, `{"messege": "AppId is not found"}`)
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
		fmt.Fprintf(ctx, `{"messege": "AppId is not found"}`)
	} else {
		fmt.Fprintf(ctx, res)
	}

}

func SendToNextServer(ctx *fasthttp.RequestCtx) {
	appid := ctx.UserValue("appid").(string)
	path := strings.Replace(string(ctx.Path()), "/lb/"+appid+"/", "", -1)
	fastestServer, isPathSaved, poolServers := controllers.GetFastestServer(appid, path)
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
}
