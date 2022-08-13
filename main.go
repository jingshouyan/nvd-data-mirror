package main

import (
	"net/http"

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
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"app":      "nvd-data-mirror",
			"version":  version,
			"commit":   commit,
			"built_by": builtBy,
			"built_at": date,
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
