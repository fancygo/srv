package conn

import (
    "sync"
    "frame/logger"
)

var (
    Conn map[uint32]*CliConn
    lock *sync.RWMutex
)

func init() {
    Conn = make(map[uint32]*CliConn)
    lock = new(sync.RWMutex)
}

func Add(cliConn *CliConn) {
    cid := cliConn.GetId()
    Conn[cid] = cliConn
    logger.Info("add conn cid = %d", cid)
}

func Del(cid uint32) {
    delete(Conn, cid)
    logger.Info("del conn cid = %d", cid)
}

func Broadcast(send_data []byte) {
    for _, v := range Conn {
        v.SendChanData(send_data)
    }
}
