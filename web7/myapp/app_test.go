package myapp

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIndex(t *testing.T) {
	assert := assert.New(t)
	ts := httptest.NewServer(NewHandler()) //가상의 웹서버를 띄운다.
	defer ts.Close()

	res, err := http.Get(ts.URL)
	assert.NoError(err)
	assert.Equal(http.StatusOK, res.StatusCode)
}
