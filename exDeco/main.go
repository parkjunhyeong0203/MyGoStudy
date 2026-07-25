package main

import (
	"fmt"

	"github.com/tuckersGo/goWeb/web9/cipher"
	"github.com/tuckersGo/goWeb/web9/lzw"
)

type Compopnent interface {
	Operator(string)
}

var sentData string

type SendComponent struct {
}

func (self *SendComponent) Operator(data string) { // 보내는 함수
	sentData = data
}

type ZipComponent struct {
	com Compopnent
}

func (self *ZipComponent) Operator(data string) { // 압축 함수
	data1, err := lzw.Write([]byte(data))
	if err != nil {
		panic(err)
	}
	self.com.Operator(string(data1))
}

type EncryptComponent struct {
	key string
	com Compopnent
}

func (self *EncryptComponent) Operator(data string) { // 암호화 함수
	data1, err := cipher.Encrypt([]byte(data), self.key)
	if err != nil {
		panic(err)
	}
	self.com.Operator(string(data1))
}

func main() { //암호화 -> 압축 -> 보내기, decorate pattern 그냥 원본 안건드리고 래핑시켜서 기능 추가하는 느낌?
	//복구 시킬려면 압축헤제 -> 복호화 -> 받기
	sender := &EncryptComponent{key: "abcde",
		com: &ZipComponent{
			com: &SendComponent{},
		},
	}
	sender.Operator("Hello World")
	fmt.Println(sentData)
}
