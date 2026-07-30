package main

import (
	"flag"
	"fmt"
	"log"
	"sync"

	"github.com/panjf2000/gnet/v2"
)

type chatServer struct {
	gnet.BuiltinEventEngine

	cliMap sync.Map
}

func (cs *chatServer) OnOpen(c gnet.Conn) (out []byte, action gnet.Action) {
	log.Printf("client connected. address:%s", c.RemoteAddr().String())
	cs.cliMap.Store(c, true) // sync.Map에서는 store 사용
	return nil, gnet.None
}

func (cs *chatServer) OnClose(c gnet.Conn, err error) gnet.Action {
	log.Printf("client connected. address:%s", c.RemoteAddr().String())
	if _, ok := cs.cliMap.LoadAndDelete(c); ok {
		log.Printf("connection removed")
	}
	return gnet.None
}

func (cs *chatServer) OnBoot(eng gnet.Engine) gnet.Action {
	log.Print("chat server os listening\n")
	return gnet.None
}

func (cs *chatServer) OnTraffic(c gnet.Conn) gnet.Action {
	buf, _ := c.Next(-1)

	cs.cliMap.Range(func(key any, value any) bool {
		if conn, ok := key.(gnet.Conn); ok { // 인처페이스 타입인 key를 gnet.Conn으로 타입 변환
			conn.AsyncWrite(buf, nil)
		}
		return true
	})
	return gnet.None
}

func main() {
	var port int
	var multicore bool

	flag.IntVar(&port, "port", 9000, "--port 9000")
	flag.BoolVar(&multicore, "multicore", false, "--multicore true")
	flag.Parse()

	chat := &chatServer{}
	log.Fatal(gnet.Run(chat, fmt.Sprintf("tcp://:%d", port), gnet.WithMulticore(multicore)))
}
