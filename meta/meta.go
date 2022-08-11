package meta

import (
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

type Meta struct {
	LastModifiedDate string
	Size             int64
	ZipSize          int64
	GzSize           int64
	Sha256           string
}

func LoadFromFile(filename string) *Meta {
	c, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Println("load meta error:", err)
		return &Meta{}
	}
	context := string(c)
	lines := strings.Split(context, "\r\n")
	meta := &Meta{}
	for _, line := range lines {
		if strings.HasPrefix(line, "lastModifiedDate:") {
			meta.LastModifiedDate = strings.TrimPrefix(line, "lastModifiedDate:")

		} else if strings.HasPrefix(line, "size:") {
			meta.Size, _ = strconv.ParseInt(strings.TrimPrefix(line, "size:"), 10, 32)
		} else if strings.HasPrefix(line, "zipSize:") {
			meta.ZipSize, _ = strconv.ParseInt(strings.TrimPrefix(line, "zipSize:"), 10, 32)
		} else if strings.HasPrefix(line, "gzSize:") {
			meta.GzSize, _ = strconv.ParseInt(strings.TrimPrefix(line, "gzSize:"), 10, 32)
		} else if strings.HasPrefix(line, "sha256:") {
			meta.Sha256 = strings.TrimPrefix(line, "sha256:")
		}
	}
	return meta
}

func (m *Meta) Equals(o *Meta) bool {
	return m.Sha256 == o.Sha256
}
