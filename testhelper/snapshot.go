package testhelper

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"reflect"
)

type snapshotContext struct {
	pointer reflect.Value
	data    []byte
}

func (ctx *snapshotContext) Rollback() error {
	zero := reflect.New(ctx.pointer.Type()).Elem()
	dec := gob.NewDecoder(bytes.NewReader(ctx.data))
	if err := dec.DecodeValue(zero); err != nil {
		return fmt.Errorf("gob: decode failed: %w", err)
	}
	ctx.pointer.Set(zero)
	return nil
}

// Snapshot data, then rollback after test finished
func SnapshotT(t T, val interface{}) {
	snap, err := Snapshot(val)
	if err != nil {
		t.Errorf("Snapshot failed: %v", err)
	}
	t.Cleanup(func() {
		if err := snap.Rollback(); err != nil {
			t.Errorf("Rollback failed: %v", err)
		}
	})
}

// Snapshot will make a deep-copy
func Snapshot(val interface{}) (*snapshotContext, error) {
	buf := new(bytes.Buffer)
	enc := gob.NewEncoder(buf)
	if err := enc.Encode(val); err != nil {
		return nil, err
	}
	return &snapshotContext{
		pointer: reflect.ValueOf(val).Elem(),
		data:    buf.Bytes(),
	}, nil
}
