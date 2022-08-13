package main

import (
	"fmt"
	"net/http"
	"runtime"

	"github.com/gin-contrib/pprof"

	"github.com/gin-gonic/gin"
	"github.com/jingshouyan/nvd-data-mirror/config"
	"github.com/jingshouyan/nvd-data-mirror/job"
	"github.com/jingshouyan/nvd-data-mirror/log"
)

var (
	version = "dev"
	commit  = "none"
	date    = "unknown"
	builtBy = "unknown"
)

func main() {
	log.Printf("nvd-data-mirror version: %s, commit: %s, built by: %s, built at: %s\n", version, commit, builtBy, date)
	go job.Start()
	defer job.Stop()
	r := gin.Default()
	pprof.Register(r)
	r.GET("/", func(c *gin.Context) {
		var m runtime.MemStats
		runtime.ReadMemStats(&m)
		c.JSON(http.StatusOK, gin.H{
			"app":      "nvd-data-mirror",
			"version":  version,
			"commit":   commit,
			"built_by": builtBy,
			"built_at": date,
			"system": gin.H{
				"heapAlloc":  B2MB(m.HeapAlloc),
				"heapInuse":  B2MB(m.HeapInuse),
				"heapSys":    B2MB(m.HeapSys),
				"stackInuse": B2MB(m.StackInuse),
				"stackSys":   B2MB(m.StackSys),
				"mspanInuse": B2MB(m.MSpanInuse),
				"mspanSys":   B2MB(m.MSpanSys),
			},
		})
	})
	r.GET("/sync", func(ctx *gin.Context) {
		go job.Sync()
		ctx.JSON(http.StatusOK, gin.H{
			"status": "ok",
		})
	})
	r.StaticFS("/data", http.Dir(config.OutputDir))
	r.Run(config.Addr)
}

func B2MB(b uint64) string {
	return fmt.Sprintf("%.2fMB", float64(b)/1024/1024)
}
