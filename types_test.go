package main

import "testing"

func TestTypeFound(t *testing.T) {
	mapping := GetTypeMapping()
	mapped, ok := mapping.GetType("jPG")
	if !ok {
		t.Errorf("Did not find jpg file type")
	}
	if "pic" != mapped {
		t.Errorf("Expected type to be of type pic")
	}
}
