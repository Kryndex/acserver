package main

import (
	"testing"
)

func assertVersion(t *testing.T, a, b string, result int) {
	if r := Version(a).compareTo(Version(b)); r != result {
		t.Fatalf("Unexpected version comparison result. Found %d, expected %d", r, result)
	}
}

func TestCompareVersion(t *testing.T) {
	assertVersion(t, "1.12", "1.12", 0)
	assertVersion(t, "1.0.0", "1", 0)
	assertVersion(t, "1", "1.0.0", 0)
	assertVersion(t, "1.05.00.0156", "1.0.221.9289", 1)
	assertVersion(t, "1", "1.0.1", -1)
	assertVersion(t, "1.0.1", "1", 1)
	assertVersion(t, "1.0.1", "1.0.2", -1)
	assertVersion(t, "1.0.2", "1.0.3", -1)
	assertVersion(t, "1.0.3", "1.1", -1)
	assertVersion(t, "1.1", "1.1.1", -1)
	assertVersion(t, "1.1.1", "1.1.2", -1)
	assertVersion(t, "1.1.2", "1.2", -1)
	assertVersion(t, "1", "1", 0)
	assertVersion(t, "20161105.091808-vd37650e", "20161105.004945-vd37650e", 1)
	assertVersion(t, "20161105.091808-42A", "20161105.091808-42A", 0)
	assertVersion(t, "20161105.091808-42A", "20161105.091808-42B", -1)
	assertVersion(t, "20161105.091808-42A", "20161105-42A", 1)
	assertVersion(t, "20161105.091808-42A", "20161105", 1)
	assertVersion(t, "20161105.091808", "20161105-42A", 1)
	assertVersion(t, "20161105.091808", "20161105-42A", 1)
}
