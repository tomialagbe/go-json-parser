package parser

import (
	"reflect"
	"testing"
)

func TestFromString(t *testing.T) {
	res, err := fromString(`{"a": 1, "b": true}`)
	if err != nil {
		t.Errorf("Error should be nil. Got %v", err)
	}

	if res == nil {
		t.Error("Expected a valid map. Got nil")
	}

	expected := map[string]interface{}{
		"a": 1,
		"b": true,
	}
	isEqual := reflect.DeepEqual(expected, res)
	if !isEqual {
		t.Errorf("Expected result to be %v. Got %v", expected, res)
	}
}
