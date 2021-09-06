package blockrand_test

import (
	"fmt"
	"testing"

	"github.com/tyeolrik/rngset"
)

func TestHello(t *testing.T) {
	well := rngset.NewWELL512a([16]uint32{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16})
	fmt.Println(well.NextFloat64())
	fmt.Println(well.NextFloat64())
	fmt.Println(well.NextFloat64())
	fmt.Println(well.NextFloat64())
	fmt.Println(well.NextFloat64())
}
