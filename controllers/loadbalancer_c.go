package controllers

import (
	"encoding/json"
	"loadbalancer/models"
	"loadbalancer/services"
	"log"
)

func JsonToStruct(jsonStr string, dynamic any) interface{} {
	json.Unmarshal([]byte(jsonStr), &dynamic)
	return dynamic
}
func StructToJson(v any) string {
	res, err := json.Marshal(v)
	if err != nil {
		println("Fail to convert to JSON")
	}
	return string(res)
}
func GetFastestServer(appid string, path string) (res string, isPathSaved bool, servers []models.Server) {
	var poolServers []models.Server
	JsonToStruct(services.GetValueRedis(appid), &poolServers)
	if len(poolServers) == 0 {
		print("No server found")
		return "", false, []models.Server{}
	}
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
	log.Println("SaveResponseTime")
	log.Println(newServer)
}
func SaveServers(appid string, servers string) {
	services.SaveValueRedis(appid, servers)
}
