package collectionfilter

import "testing"

func TestFilter(t *testing.T) {
	filter := New(nil, nil, "optimism", "MiSwap")
	filter.Add("TEst")
	if !filter.Contains("Test") {
		t.Error("Expected filter contain 'Test'")
	}

	filter.Remove("Test")
	if filter.Contains("test") {
		t.Error("Expected filter not contain 'test'")
	}
}
