package hotconfig

import (
	"fmt"
	"sync/atomic"
)

var (
	ErrMissingLoader = fmt.Errorf("missing loader")
	ErrInvliadConfig = fmt.Errorf("invalid config")
)

type loader[T any] func() (*T, error)

type HotConfig[T any] struct {
	loader loader[T]
	value  atomic.Pointer[T]
}

func New[T any](loader loader[T]) (*HotConfig[T], error) {
	obj := HotConfig[T]{loader: loader}
	if err := obj.init(); err != nil {
		return nil, err
	}
	return &obj, nil
}

func NewOrPanic[T any](loader loader[T]) *HotConfig[T] {
	obj, err := New(loader)
	if err != nil {
		panic(err)
	}
	return obj
}

func (c *HotConfig[T]) init() error {
	if c.loader == nil {
		return ErrMissingLoader
	}
	v, err := c.loader()
	if err != nil {
		return err
	}
	c.value.Store(v)
	if v := c.Load(); v == nil {
		return ErrInvliadConfig
	}
	return nil
}

func (c *HotConfig[T]) Load() *T {
	return c.value.Load()
}

// hot reload safety
func (c *HotConfig[T]) Reload() error {
	if c.loader == nil {
		return ErrMissingLoader
	}
	v, err := c.loader()
	if err != nil {
		return err
	}
	c.value.Store(v)
	return nil
}

func (c *HotConfig[T]) SetLoader(loader loader[T]) {
	if loader == nil {
		panic(fmt.Errorf("invalid loader"))
	}
	c.loader = loader
}
