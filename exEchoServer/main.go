package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/panjf2000/gnet/v2"
)

type echoServer struct { //먼저 서버부분
	gnet.BuiltinEventEngine

	eng       gnet.Engine
	addr      string
	multicore bool
}

func (es *echoServer) OnOpen(c gnet.Conn) (out []byte, action gnet.Action) { //클라가 접속할때 호출 (return이 두개)
	log.Printf("Client connected. address:%s", c.RemoteAddr().String())
	return nil, gnet.None
}

func (es *echoServer) OnClose(c gnet.Conn, err error) (action gnet.Action) {
	log.Printf("client disconnected. address:%s", c.RemoteAddr().String())
	return gnet.None
}

func (es *echoServer) OnBoot(eng gnet.Engine) gnet.Action {
	es.eng = eng
	log.Printf("echo server with multicore=%t is listening on %s\n", es.multicore, es.addr)
	return gnet.None
}

func (es *echoServer) OnTraffic(c gnet.Conn) gnet.Action {
	buf, _ := c.Next(-1) // 모든 데이터를 읽는다.
	c.Write(buf)         //읽은 데이터를 다시 전송한다.
	return gnet.None
}

func main() {
	var port int
	var multicore bool

	flag.IntVar(&port, "port", 9000, "--port 9000")
	flag.BoolVar(&multicore, "multicore", false, "--multicore true")
	flag.Parse()

	echo := &echoServer{
		addr:      fmt.Sprintf("tcp://:%d", port),
		multicore: multicore,
	}
	log.Fatal(gnet.Run(echo, echo.addr, gnet.WithMulticore(multicore)))
}
