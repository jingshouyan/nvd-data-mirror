package utils

import (
	"compress/gzip"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"time"

	"github.com/jingshouyan/nvd-data-mirror/meta"
)

const (
	gz string = ".gz"
)

func SyncRetireJs(retireJsRrl, outDir string) {
	log.Println("---------- Begin Sync retire.js-----------")
	err := syncRetireJs(retireJsRrl, outDir)
	if err != nil {
		log.Println(err)
	}
	log.Println("---------- End Sync retire.js-----------")
}

func syncRetireJs(retireJsRrl, outDir string) error {
	subfix := ".tmp." + time.Now().Format("2006-01-02_15-04-05")
	tmp, err := Download(retireJsRrl, outDir, subfix)
	if err != nil {
		return err
	}
	if validJson(tmp) {
		replace(tmp, tmp[:len(tmp)-len(subfix)])
	} else {
		log.Println("invalid json, remove", tmp)
		os.Remove(tmp)
	}
	return nil
}

func validJson(filepath string) bool {
	b, err := ioutil.ReadFile(filepath)
	if err != nil {
		return false
	}
	return json.Valid(b)
}

func SyncVnd(metaUrl, dataUrl, outDir string) {
	idx := strings.LastIndex(metaUrl, "/")
	name := metaUrl[idx+1 : len(metaUrl)-5]
	log.Println("---------- Begin Sync", name, "-----------")
	err := syncVnd(metaUrl, dataUrl, outDir)
	if err != nil {
		log.Println(err)
	}
	log.Println("---------- End Sync", name, "-----------")
}

func syncVnd(metaUrl, dataUrl, outDir string) error {
	os.MkdirAll(outDir, os.ModePerm)
	subfix := ".tmp." + time.Now().Format("2006-01-02_15-04-05")
	// 下载 meta
	metaTmp, err := Download(metaUrl, outDir, subfix)
	if err != nil {
		return err
	}
	metaFile := metaTmp[:len(metaTmp)-len(subfix)]
	m0 := meta.LoadFromFile(metaFile)
	meta := meta.LoadFromFile(metaTmp)
	if m0.Equals(meta) {
		log.Println("meta is equal, no need to download.")
		os.Remove(metaTmp)
		return nil
	}
	// 下载数据
	dataTmp, err := Download(dataUrl, outDir, subfix)
	if err != nil {
		return err
	}
	dataFile := dataTmp[:len(dataTmp)-len(subfix)]
	jsonFile := dataFile[:len(dataFile)-len(gz)]
	err = gzUnpack(dataTmp, jsonFile)
	if err != nil {
		return err
	}
	if !checkData(jsonFile, meta) {
		os.Remove(dataTmp)
		return fmt.Errorf("metadata check failed.")
	}

	// 替换文件
	replace(metaTmp, metaFile)
	replace(dataTmp, dataFile)
	return nil
}

func checkData(jsonfile string, meta *meta.Meta) bool {
	fs, err := os.Stat(jsonfile)
	if err != nil {
		return false
	}
	if fs.Size() != meta.Size {
		return false
	}
	if sha(jsonfile) != meta.Sha256 {
		return false
	}
	return true
}

func replace(source, dist string) {
	if _, err := os.Stat(dist); err == nil {
		log.Println("removing", dist)
		os.Remove(dist)
	}
	log.Println("replacing", source, "to", dist)
	os.Rename(source, dist)
}

func gzUnpack(source, dist string) error {
	f, err := os.Open(source)
	if err != nil {
		return err
	}
	defer f.Close()
	gr, err := gzip.NewReader(f)
	if err != nil {
		return err
	}
	defer gr.Close()
	df, err := os.Create(dist)
	if err != nil {
		return err
	}
	defer df.Close()
	log.Println("unpacking", source, "to", dist)
	if _, err := io.Copy(df, gr); err != nil {
		return err
	}
	return nil
}

func sha(jsonFile string) string {
	f, err := os.Open(jsonFile)
	if err != nil {
		log.Println(err)
	}
	defer f.Close()

	h := sha256.New()
	if _, err := io.Copy(h, f); err != nil {
		log.Fatal(err)
	}
	return strings.ToUpper(fmt.Sprintf("%x", h.Sum(nil)))
}
