package hotconfig

import (
	"fmt"
	"testing"

	"github.com/gamedev-embers/libserver/hotconfig/blocks"
	"github.com/stretchr/testify/assert"
)

type ConfigA struct {
	Gate *blocks.Server `json:"gate" toml:"gate"`
}

func TestHotObject(t *testing.T) {
	loaderFunc := func() func() (*ConfigA, error) {
		countor := 0
		return func() (*ConfigA, error) {
			countor++
			return &ConfigA{
				Gate: &blocks.Server{
					Host: "host",
					Ports: map[string]int{
						"tcp": countor,
					},
				},
			}, nil
		}
	}

	t.Run("Load", func(t *testing.T) {
		assert := assert.New(t)
		loader := loaderFunc()
		obj := NewOrPanic(loader)
		assert.Equal(1, obj.Load().Gate.Ports["tcp"])
		assert.Equal(1, obj.Load().Gate.Ports["tcp"])
		assert.Equal(1, obj.Load().Gate.Ports["tcp"])
	})

	t.Run("Reload", func(t *testing.T) {
		assert := assert.New(t)
		loader := loaderFunc()
		obj := NewOrPanic(loader)
		v := obj.Load()
		assert.Equal(1, v.Gate.Ports["tcp"])

		// first reload
		obj.Reload()
		v2 := obj.Load()
		assert.False(v == v2)
		assert.Equal(1, v.Gate.Ports["tcp"])
		assert.Equal(2, v2.Gate.Ports["tcp"])

		// second reload
		obj.Reload()
		v3 := obj.Load()
		assert.False(v == v2)
		assert.False(v == v3 || v2 == v3)
		assert.Equal(1, v.Gate.Ports["tcp"])
		assert.Equal(2, v2.Gate.Ports["tcp"])
		assert.Equal(3, v3.Gate.Ports["tcp"])
	})
}

func TestHotObject_Errors(t *testing.T) {
	ErrForTest := fmt.Errorf("error")

	loaderGood := func() (*ConfigA, error) {
		return &ConfigA{}, nil
	}
	loaderBad := func() (*ConfigA, error) {
		return nil, ErrForTest
	}

	t.Run("New", func(t *testing.T) {
		assert := assert.New(t)
		assert.PanicsWithError(ErrForTest.Error(), func() {
			obj := NewOrPanic(loaderBad)
			_ = obj
		})
		// assert.Equal(ErrForTest, err)
	})

	t.Run("SetLoader", func(t *testing.T) {
		assert := assert.New(t)
		obj := NewOrPanic(loaderGood)
		v := obj.Load()
		assert.NotNil(v)
	})

}
