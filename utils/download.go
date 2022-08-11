package utils

import (
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

func Download(url, outDir, subfix string) (string, error) {
	idx := strings.LastIndex(url, "/")
	filename := url[idx+1:] + subfix
	absOutputFile := filepath.Join(outDir, filename)
	log.Println("downloading", url, "to", absOutputFile)
	if _, err := get(url, absOutputFile); err != nil {
		return "", err
	}
	return absOutputFile, nil
}

func get(url, absOutputFile string) (int64, error) {
	out, err := os.Create(absOutputFile)
	if err != nil {
		return 0, err
	}
	defer out.Close()
	resp, err := http.Get(url)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()
	n, err := io.Copy(out, resp.Body)
	return n, err
}
