package ecode

import (
	"fmt"
	"testing"
)

func TestNewResponse(t *testing.T) {
	response := NewResponse(OK, nil)
	fmt.Println(response)
}
