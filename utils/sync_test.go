package utils

import (
	"path/filepath"
	"testing"

	"github.com/jingshouyan/nvd-data-mirror/config"
)

func TestSync(t *testing.T) {
	outDir := "./tmp3"
	outDir, _ = filepath.Abs(outDir)
	SyncVnd(config.Cve11ModifiedMetaUrl, config.Cve11ModifiedJsonUrl, outDir)
}

func TestGzUnpack(t *testing.T) {
	err := gzUnpack("./nvdcve-1.1-modified.json.gz", "./nvdcve-1.1-modified.json")
	if err != nil {
		t.Error(err)
	}
}

func TestSha(t *testing.T) {
	st := sha("./tmp/nvdcve-1.1-modified.json")
	t.Log(st)
}
