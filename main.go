package main

import (
	"loadbalancer/routers"
	"log"

	"github.com/buaazp/fasthttprouter"
	"github.com/valyala/fasthttp"
)

var port = "1234"

func main() {
	router := fasthttprouter.New()
	routers.LoadBalancerRouters(router)

	server := &fasthttp.Server{Handler: router.Handler}
	go func() {
		if err := server.ListenAndServe(":" + port); err != nil {
			log.Printf("Error starting server: %s\n", err)
		}
	}()
	log.Println("Server listen on :" + port)
	select {}
}
