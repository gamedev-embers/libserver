package loot

import (
	"fmt"
	"log/slog"
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
)

type CustomItem struct {
	ID     int
	Weight uint32
}

func (c CustomItem) GetWeight() int64 {
	return int64(c.Weight)
}

func make_items(count int) []*CustomItem {
	rs := make([]*CustomItem, count)
	for i := 0; i < count; i++ {
		num := i + 1
		rs[i] = &CustomItem{ID: num, Weight: uint32(num)}
	}
	return rs
}

func TestBag(t *testing.T) {
	inputs := []CustomItem{
		{ID: 1, Weight: 1},
		{ID: 2, Weight: 3},
		{ID: 3, Weight: 5},
		{ID: 4, Weight: 7},
		{ID: 5, Weight: 10},
		{ID: 6, Weight: 30},
		{ID: 7, Weight: 50},
	}

	r := rand.New(rand.NewSource(0))
	bag := New(inputs)
	t.Run("DropMany", func(t *testing.T) {
		assert := assert.New(t)
		dumps := bag.Dump()
		for _, row := range bag.Dump() {
			fmt.Println(row)
		}
		count := 5
		drops, err := bag.DropMany(r, count)
		assert.NoError(err)
		assert.Equal(count, len(drops))
		assert.Equal(dumps, bag.Dump())
		for _, item := range drops {
			t.Logf("%3v: %v", item.ID, item.Weight)
		}
	})

	t.Run("DropOne", func(t *testing.T) {
		for _, times := range []int{1, 10, 100, 1000} {
			name := fmt.Sprintf("loops=%d", times)
			t.Run(name, func(t *testing.T) {
				assert := assert.New(t)
				dumps, err := bag.DryRun(r, 1000)
				assert.NoError(err)

				slog.Info("Dumps:" + name)
				for _, row := range dumps {
					slog.Info("  " + row.String())
				}
			})
		}
		t.Fatalf("Done. Read the report above.")
	})
}

func TestBag_DropMany(t *testing.T) {
	assert := assert.New(t)
	inputs := make_items(10)

	bag := New(inputs)
	r := rand.New(rand.NewSource(0))
	for i := 0; i < 100; i++ {
		drops, err := bag.DropMany(r, 10)
		assert.NoError(err)
		assert.Equal(10, len(drops))

		drops2, err := bag.DropMany(r, 11)
		assert.NoError(err)
		assert.Equal(10, len(drops2))
	}
}
