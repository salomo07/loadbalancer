package controllers

import (
	"encoding/json"
	"loadbalancer/models"
	"loadbalancer/services"
	"log"
)

func GetFastestServer(appid string, path string) (res string, isPathSaved bool, servers []models.Server) {
	var poolServers []models.Server
	json.Unmarshal([]byte(services.GetValueRedis(appid)), &poolServers)
	var fastestServer models.Server
	var lowestLatency int64 = -1
	for _, val := range poolServers {
		if val.Path == path {
			if lowestLatency == -1 || val.Latency < lowestLatency {
				lowestLatency = val.Latency
				fastestServer = val
				isPathSaved = true
			}
		}
	}
	if fastestServer.Address == "" {
		fastestServer = poolServers[0]
		isPathSaved = false
	}
	servers = poolServers
	return fastestServer.Address, isPathSaved, servers
}
func SaveResponseTime(appid string, servers []models.Server, newServer models.Server) {
	log.Println("sdf")
	log.Println(newServer)
}
func SaveServers(appid string, servers string) {
	services.SaveValueRedis(appid, servers)
}
