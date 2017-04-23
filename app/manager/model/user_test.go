package model

import (
	"testing"
)

func TestRandToken(t *testing.T) {
	token := RandToken()
	if len(token) != 172 {
		t.Error(token)
	}
}
