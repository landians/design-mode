package prototype

import (
	"fmt"
	"testing"
)

func Test_ProtoType(t *testing.T) {
	u1 := defaultUserFactory.Create()
	fmt.Printf("u1 = %v\n", u1)

	u2 := defaultUserFactory.Create()
	fmt.Printf("u2 = %v\n", u2)
}