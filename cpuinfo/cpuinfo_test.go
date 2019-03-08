package cpuinfo

import (
	"testing"
)

func TestGet(t *testing.T) {
	_, err := Get()
	if err != nil {
		t.Error(err)
	}
}
