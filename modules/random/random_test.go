package random

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"strconv"
)

func TestRandom(t *testing.T) {
	t.Parallel()

	min := 0
	max := 100

	for i := 0; i < 100000; i++ {
		value := Random(min, max)
		assert.True(t, value >= min && value <= max)
	}
}

func TestRandomInt(t *testing.T) {
	t.Parallel()

	min := 0
	max := 1000

	list := []int{}
	for i := min; i < max; i++ {
		list = append(list, i)
	}

	for i := 0; i < 100000; i++ {
		value := RandomInt(list)
		assert.Contains(t, list, value)
	}
}

func TestRandomString(t *testing.T) {
	t.Parallel()

	min := 0
	max := 1000

	list := []string{}
	for i := min; i < max; i++ {
		list = append(list, strconv.Itoa(i))
	}

	for i := 0; i < 100000; i++ {
		value := RandomString(list)
		assert.Contains(t, list, value)
	}
}

func TestUniqueId(t *testing.T) {
	t.Parallel()

	previouslySeen := map[string]bool{}

	for i := 0; i < 100; i++ {
		uniqueId := UniqueId()
		assert.Len(t, uniqueId, 6)
		assert.NotContains(t, previouslySeen, uniqueId)

		previouslySeen[uniqueId] = true
	}
}