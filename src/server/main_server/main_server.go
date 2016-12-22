package main

import (
	"frame"
    "frame/logger"
    "frame/def"
    "server/main_server/conn"
	module "server/main_server/game"
    "server/main_server/dbdata"
    "golang.org/x/net/websocket"
    "net/http"
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

    //初始化主服务
	wsAddr := frame.GetMainServerIP() + ":" + frame.GetMainServerPort()
	logger.Info("wsAddr = %+v", wsAddr)
    http.Handle("/", websocket.Handler(cliHandler))
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
	go cliConn.HandleClient(waitGroup)
    waitGroup.Wait()
}
