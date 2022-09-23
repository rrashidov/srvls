package model

import "testing"

func TestGenerateId(t *testing.T) {
	id1 := GenerateId("some string to hash")
	id2 := GenerateId("another string to hash")

	if id1 == id2 {
		t.Errorf("Different ids should be generated for different inputs")
	}
}
