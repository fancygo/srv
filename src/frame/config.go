package frame

//该文件主要是获取服务器启动的各个配置ip,port等

import (
	"encoding/json"
	"frame/logger"
	"io/ioutil"
	"os"
	"path"
)

type Server struct {
	IP   string `json:"ip"`
	Port string `json:"port"`
}
type Db struct {
	Host string `json:"host"`
	Port string `json:"port"`
	User string `json:"user"`
	Pswd string `json:"pswd"`
}

type ServerCfg struct {
	MainServer Server `json:"mainserver"`
	DbRpc      Server `json:"dbrpc"`
	DbServer   Server `json:"dbserver"`
	LogServer  Server `json:"logserver"`
	Db         Db     `json:"mysql"`
}

var server ServerCfg

func LoadConfig() {
	configName := path.Join(GetConfDir(), "config.json")
	file, err := os.Open(configName)
	if err != nil {
		logger.Error("error =", err)
		return
	}
	defer file.Close()
	data, err := ioutil.ReadAll(file)
	if err != nil {
		logger.Error("error =", err)
		return
	}
	err = json.Unmarshal(data, &server)
	if err != nil {
		logger.Error("error =", err)
		return
	}
}

func GetMainServerIP() string {
	return server.MainServer.IP
}
func GetMainServerPort() string {
	return server.MainServer.Port
}
func GetDbRpcIP() string {
	return server.DbRpc.IP
}
func GetDbRpcPort() string {
	return server.DbRpc.Port
}
func GetLogServerIP() string {
	return server.LogServer.IP
}
func GetLogServerPort() string {
	return server.LogServer.Port
}
func GetDbServerIP() string {
	return server.DbServer.IP
}
func GetDbServerPort() string {
	return server.DbServer.Port
}
func GetDbHost() string {
	return server.Db.Host
}
func GetDbPort() string {
	return server.Db.Port
}
func GetDbUser() string {
	return server.Db.User
}
func GetDbPswd() string {
	return server.Db.Pswd
}
