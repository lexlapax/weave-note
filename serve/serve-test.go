package serve

import "testing"

func TestServe(t *testing.T) {
	want := "Hi, I'm weave-serve"

	if got := Serve(); got != want {
		t.Errorf("Serve() = %q, want %q", got, want)
	}
}
