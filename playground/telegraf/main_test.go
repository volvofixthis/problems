package benchtest

import (
	"math/rand"
	"strconv"
	"testing"
)

const (
	seed = 17
	// m2 cpu cache line
	cacheLine = 128
	dataSize  = cacheLine * 1000
)

var randKeys []int
var values []*testObj

func init() {
	r := rand.New(rand.NewSource(seed))
	for i := range dataSize {
		randKeys = append(randKeys, r.Intn(dataSize))
		values = append(values, &testObj{content: strconv.Itoa(i)})
	}
}

type testObj struct {
	content string
}

type HashMap struct {
	data []*testObj
}

func New() *HashMap {
	return &HashMap{data: make([]*testObj, 0)}
}

func (h *HashMap) Set(key int, value *testObj) {
	if len(h.data) <= key {
		if cap(h.data) > key {
			h.data = h.data[:cap(h.data)]
		} else {
			h.data = append(h.data, make([]*testObj, key-len(h.data)+1)...)
		}
	}
	h.data[key] = value
}

func (h *HashMap) Get(key int) (*testObj, bool) {
	if len(h.data) <= key {
		return nil, false
	}
	value := h.data[key]
	return value, value != nil
}

func BenchmarkHandmadeHashMap(b *testing.B) {
	h := New()

	// Populate the hashmap
	for i := 0; i < b.N; i++ {
		p := i % (dataSize - 1)
		h.Set(randKeys[p], values[p])
	}

	// Retrieve values
	for i := 0; i < b.N; i++ {
		_, _ = h.Get(i)
	}
}

func BenchmarkGoBuiltInMap(b *testing.B) {
	m := make(map[int]*testObj)

	// Populate the built-in map
	for i := 0; i < b.N; i++ {
		p := i % (dataSize - 1)
		m[randKeys[p]] = values[p]
	}

	// Retrieve values
	for i := 0; i < b.N; i++ {
		_, _ = m[i]
	}
}
