package utils

import (
	"testing"

	"github.com/jingshouyan/nvd-data-mirror/config"
)

func TestDownLoad(t *testing.T) {
	url := config.Cve11ModifiedMetaUrl

	absOutputFile := "sample.meta"

	n, err := get(url, absOutputFile)
	if err != nil {
		t.Error(err)
	}
	t.Log(n)
}
