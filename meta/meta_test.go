package meta

import (
	"testing"
)

func TestLoadFromFile(t *testing.T) {
	meta := LoadFromFile("sample.meta")

	if "A448ED7DC59ED515375B454B9F3348FABB4A9241728B36A4CD1E8E3273E6CA14" != meta.Sha256 {
		t.Error("sha256 not match")
	}
}
