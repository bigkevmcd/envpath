package envpath

import (
	"os"
	"testing"
)

func TestAppend(t *testing.T) {
	appendTests := []struct {
		name    string
		envVar  string
		newPath string
		wantVar string
	}{
		{"empty path", "", "/usr/bin", "/usr/bin"},
		{"existing path", "/usr/local/bin", "/usr/bin", "/usr/local/bin" + string(os.PathListSeparator) + "/usr/bin"},
		{"new path exists in path already", "/usr/bin:/usr/local/bin", "/usr/bin", "/usr/local/bin" + string(os.PathListSeparator) + "/usr/bin"},
	}

	for _, tt := range appendTests {
		t.Run(tt.name, func(t *testing.T) {
			got := Append(tt.envVar, tt.newPath)
			if got != tt.wantVar {
				t.Fatalf("got %s, want %s", got, tt.wantVar)
			}
		})
	}
}

func TestPrepend(t *testing.T) {
	prependTests := []struct {
		name    string
		envVar  string
		newPath string
		wantVar string
	}{
		{"empty path", "", "/usr/bin", "/usr/bin"},
		{"existing path", "/usr/local/bin", "/usr/bin", "/usr/bin" + string(os.PathListSeparator) + "/usr/local/bin"},
		{"new path exists in path already", "/usr/local/bin:/usr/bin", "/usr/bin", "/usr/bin" + string(os.PathListSeparator) + "/usr/local/bin"},
	}

	for _, tt := range prependTests {
		t.Run(tt.name, func(t *testing.T) {
			got := Prepend(tt.envVar, tt.newPath)
			if got != tt.wantVar {
				t.Fatalf("got %s, want %s", got, tt.wantVar)
			}
		})
	}
}
