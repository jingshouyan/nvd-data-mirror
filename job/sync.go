package job

import (
	"fmt"
	"strings"

	"github.com/jingshouyan/nvd-data-mirror/log"

	"github.com/jingshouyan/nvd-data-mirror/config"
	"github.com/jingshouyan/nvd-data-mirror/utils"
	cp "github.com/otiai10/copy"
	"github.com/robfig/cron/v3"
)

var c = cron.New()

func Start() {
	syncAll()
	c.AddFunc(config.Cron, syncAll)
	c.Start()
}

func Stop() {
	c.Stop()
}

func Sync() {
	syncAll()
}

func syncAll() {
	utils.SyncVnd(config.Cve11ModifiedMetaUrl, config.Cve11ModifiedJsonUrl, config.TmpDir)
	utils.SyncVnd(config.Cve11RecentMetaUrl, config.Cve11RecentJsonUrl, config.TmpDir)
	for i := config.StartYear; i <= config.EndYear; i++ {
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
