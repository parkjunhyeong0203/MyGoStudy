package main

import (
	"bufio"
	"bytes"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIndexPage(t *testing.T) {
	/*assert := assert.New(t) //그냥 함수 느낌으로 테스트하는 방법
	res := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/", nil)

	mux := NewHandler()
	mux.ServeHTTP(res, req)

	assert.Equal(http.StatusOK, res.Code) */

	assert := assert.New(t)
	ts := httptest.NewServer(NewHandler()) //가상의 웹서버를 띄운다.
	defer ts.Close()

	res, err := http.Get(ts.URL)
	assert.NoError(err)
	assert.Equal(http.StatusOK, res.StatusCode)
	data, _ := io.ReadAll(res.Body)
	assert.Equal("Hello World", string(data))
}

func TestDecoHandler(t *testing.T) { //로그가 제대로 찍혔는지 확인

	assert := assert.New(t)
	ts := httptest.NewServer(NewHandler()) //가상의 웹서버를 띄운다.
	defer ts.Close()

	buf := &bytes.Buffer{}
	log.SetOutput(buf) //output을 모니터 말고 버퍼로 설정

	res, err := http.Get(ts.URL)
	assert.NoError(err)
	assert.Equal(http.StatusOK, res.StatusCode)
	data, _ := io.ReadAll(res.Body)
	assert.Equal("Hello World", string(data))

	r := bufio.NewReader(buf) //바이너리를 읽을수 있도록
	line, _, err := r.ReadLine()
	assert.NoError(err)
	assert.Contains(string(line), "[LOGGER1] Started")
}
