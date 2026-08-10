package databuilder

import "testing"

func TestAutoMap(t *testing.T) {
	type TestStruct struct {
		Id   uint32
		Name string
	}

	rows := []TestStruct{
		{Id: 1, Name: "Alice"},
		{Id: 2, Name: "Bob"},
	}

	result := AutoMap(rows)

	if len(result) != 2 {
		t.Errorf("Expected map length 2, got %d", len(result))
	}

	if result[1].Name != "Alice" {
		t.Errorf("Expected Name 'Alice' for Id 1, got '%s'", result[1].Name)
	}

	if result[2].Name != "Bob" {
		t.Errorf("Expected Name 'Bob' for Id 2, got '%s'", result[2].Name)
	}
}
