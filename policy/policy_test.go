package policy

import (
	"math/rand"
	"testing"
	"time"
)

func Test_Policy(t *testing.T) {
	size := 20
	data := make([]int, size)

	r := rand.New(rand.NewSource(time.Now().Unix()))

	for i := range data {
		data[i] = r.Intn(100)
	}
	t.Logf("UnsortedData \t= %v\n", data)

	fnMakeCopy := func() []int{
		copies := make([]int, size)
		for i, v := range data {
			copies[i] = v
		}
		return copies
	}

	fnTestPolicy := func(policy ISortPolicy) {
		sorted := policy.Sort(fnMakeCopy())
		t.Logf("%s \t= %v", policy.Name(), sorted)
	}

	fnTestPolicy(newBubbleSortPolicy())
	fnTestPolicy(newSelectSortPolicy())
}

