package conf

import (
	"fmt"
	"testing"
)

func TestNew(t *testing.T) {
	fmt.Println(Cf.GetString("name"))
}
