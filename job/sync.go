package job

import (
	"fmt"
	"strings"
	"sync/atomic"
	"time"

	"github.com/jingshouyan/nvd-data-mirror/log"

	"github.com/jingshouyan/nvd-data-mirror/config"
	"github.com/jingshouyan/nvd-data-mirror/utils"
	cp "github.com/otiai10/copy"
	"github.com/robfig/cron/v3"
)

var c = cron.New()
var b = &atomic.Bool{}

func Start() {
	sync()
	c.AddFunc(config.Cron, sync)
	c.Start()
}

func Stop() {
	c.Stop()
}

func Sync() {
	sync()
}

func sync() {
	if b.CompareAndSwap(false, true) {
		syncData()
		defer b.Store(false)
	} else {
		log.Println("sync is running, skip.")
	}
}

func syncData() {
	endYear := time.Now().Year()
	utils.SyncVnd(config.Cve11ModifiedMetaUrl, config.Cve11ModifiedJsonUrl, config.TmpDir)
	utils.SyncVnd(config.Cve11RecentMetaUrl, config.Cve11RecentJsonUrl, config.TmpDir)
	for i := config.StartYear; i <= endYear; i++ {
		metaUrl := fmt.Sprintf(config.Cve11BaseMetaUrl, i)
		dataUrl := fmt.Sprintf(config.Cve11BaseJsonUrl, i)
		utils.SyncVnd(metaUrl, dataUrl, config.TmpDir)
	}
	utils.SyncRetireJs(config.RetireJsUrl, config.TmpDir)
	log.Println("Copy files [", config.TmpDir, "] to [", config.OutputDir, "]")
	ops := cp.Options{
		Skip: func(path string) (bool, error) {
			return strings.Contains(path, ".tmp"), nil
		},
	}
	cp.Copy(config.TmpDir, config.OutputDir, ops)
	log.Println("Done.")
}
