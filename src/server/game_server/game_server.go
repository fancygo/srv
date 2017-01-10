package main

//服务主程序

import (
	"frame"
	"frame/def"
	"frame/logger"
	"golang.org/x/net/websocket"
	"net/http"
	"server/main_server/conn"
	"server/main_server/dbdata"
	module "server/main_server/game"
	"sync"
)

func main() {
	logger.Info("-----------------------Main server start-----------------------")
	//初始化db
	if !dbdata.Init() {
		return
	}
	dbNormalCraft := dbdata.LoadNormalCraft()
	dbGoodCraft := dbdata.LoadGoodCraft()

	//初始化mgr
	module.Init()

	//初始化mgr数据
	for _, v := range dbNormalCraft {
		module.CraftApi.InitNormalCraft(v.Id, v.Author, v.Rect, v.Data, v.Praise)
		module.RankApi.AddRankData(v.Id, v.Praise, def.CraftNormal)
	}
	for _, v := range dbGoodCraft {
		module.CraftApi.InitGoodCraft(v.Id, v.Author, v.Rect, v.Data, v.Praise)
		module.RankApi.AddRankData(v.Id, v.Praise, def.CraftGood)
	}
	module.RankApi.DoSort()

	//初始化主服务, 使用websocket, 之前做过一版tcp连接, 也是因为客户端不好处理, 改成websocket
	wsAddr := frame.GetMainServerIP() + ":" + frame.GetMainServerPort()
	logger.Info("wsAddr = %+v", wsAddr)
	http.Handle("/", websocket.Handler(cliHandler))
	//监听
	err := http.ListenAndServe(wsAddr, nil)
	if err != nil {
		logger.Error("websocket handle err =", err)
		return
	}
}

func cliHandler(conn_ws *websocket.Conn) {
	logger.Notice("connWs =", conn_ws.RemoteAddr().String())
	cliConn := conn.NewCliConn(conn_ws)
	conn.Add(cliConn)
	waitGroup := &sync.WaitGroup{}
	waitGroup.Add(1)
	//这边每个连接创建一个线程进行消息处理
	go cliConn.HandleClient(waitGroup)
	waitGroup.Wait()
}