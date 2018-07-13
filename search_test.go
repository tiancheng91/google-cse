package cse

import (
	"testing"
)

func TestAgent(t *testing.T) {
	agent := New("008063188944472181627:xqha3yefaee", "zh_CN")

	token, err := agent.parseToken()
	if err != nil {
		t.Errorf("getToken() get error : %s", err)
	}
	if len(token) != 48 {
		t.Errorf("getToken() parse error : token is %v", token)
	}

	agent.refreshToken()

	result, err := agent.Query("abcd", 1, 20)
	if err != nil {
		t.Errorf("query() error : %v", err)
	}

	if len(result.Results) != 20 {
		t.Errorf("query() result size error : %v", len(result.Results))
	}
}
