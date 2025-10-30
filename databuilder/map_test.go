package databuilder

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMap(t *testing.T) {
	type testStruct struct {
		ID   int
		Name string
	}

	lst := []*testStruct{
		{ID: 1, Name: "Alice"},
		{ID: 2, Name: "Bob"},
		{ID: 3, Name: "Charlie"},
	}
	idGetter := func(t *testStruct) int {
		return t.ID
	}
	mapped, err := Map(lst, idGetter)
	assert.NoError(t, err)
	assert.Equal(t, 3, len(mapped))
	for _, data := range lst {
		assert.Equal(t, data, mapped[data.ID])
	}
}

func TestMapList(t *testing.T) {
	type testStruct struct {
		ID   int
		Name string
	}

	lst := []*testStruct{
		{ID: 1, Name: "Alice"},
		{ID: 2, Name: "Bob"},
		{ID: 3, Name: "Charlie"},
		{ID: 1, Name: "Alice2"},
		{ID: 2, Name: "Bob2"},
		{ID: 3, Name: "Charlie2"},
	}
	idGetter := func(t *testStruct) int {
		return t.ID
	}
	mapped, err := MapList(lst, idGetter)
	assert.NoError(t, err)
	assert.Equal(t, 3, len(mapped))
	assert.Equal(t, 2, len(mapped[1]))
	assert.Equal(t, 2, len(mapped[2]))
	assert.Equal(t, 2, len(mapped[3]))
}
