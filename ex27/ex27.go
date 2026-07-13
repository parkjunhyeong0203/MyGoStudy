package main

// dip 예제, 추상 관계로 역전시켜서 추상 모듈끼리 연결 = observer 패턴
import "fmt"

type Event interface {
	Register(EventListener)
}

type EventListener interface {
	OnFire()
}

type Mail struct {
	listener EventListener
}

func (m *Mail) Register(listener EventListener) {
	m.listener = listener
}
func (m *Mail) OnRecv() {
	m.listener.OnFire()
}

type Alarm struct {
}

func (a *Alarm) OnFire() {
	fmt.Println("메일이 왔습니다.")
}
func main() {
	mail := &Mail{}
	listener := &Alarm{}

	mail.Register(listener)
	mail.OnRecv()
}
