package IdGenerator

import (
	"fmt"
	"testing"
)

func TestDistributedSystemIdGenerator(t *testing.T) {
	id, err := DistributedSystemIdGenerator()
	if err != nil {
		panic(err)
	}
	fmt.Println(id)
}
