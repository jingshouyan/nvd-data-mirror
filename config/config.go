package config

import (
	"time"

	"github.com/spf13/pflag"
)

var (
	Cve11ModifiedJsonUrl string
	Cve11RecentJsonUrl   string
	Cve11BaseJsonUrl     string
	Cve11ModifiedMetaUrl string
	Cve11RecentMetaUrl   string
	Cve11BaseMetaUrl     string
	StartYear            int
	EndYear              int
	TmpDir               string
	OutputDir            string
)

func init() {
	pflag.StringVar(&Cve11ModifiedJsonUrl, "cve-modified-json-url", "https://nvd.nist.gov/feeds/json/cve/1.1/nvdcve-1.1-modified.json.gz", "NVD CVE 1.1 JSON modified url")
	pflag.StringVar(&Cve11RecentJsonUrl, "cve-recent-json-url", "https://nvd.nist.gov/feeds/json/cve/1.1/nvdcve-1.1-recent.json.gz", "NVD CVE 1.1 JSON recent url")
	pflag.StringVar(&Cve11BaseJsonUrl, "cve-base-json-url", "https://nvd.nist.gov/feeds/json/cve/1.1/nvdcve-1.1-%d.json.gz", "NVD CVE 1.1 JSON base url")
	pflag.StringVar(&Cve11ModifiedMetaUrl, "cve-modified-meta-url", "https://nvd.nist.gov/feeds/json/cve/1.1/nvdcve-1.1-modified.meta", "NVD CVE 1.1 JSON modified meta url")
	pflag.StringVar(&Cve11RecentMetaUrl, "cve-recent-meta-url", "https://nvd.nist.gov/feeds/json/cve/1.1/nvdcve-1.1-recent.meta", "NVD CVE 1.1 JSON recent meta url")
	pflag.StringVar(&Cve11BaseMetaUrl, "cve-base-meta-url", "https://nvd.nist.gov/feeds/json/cve/1.1/nvdcve-1.1-%d.meta", "NVD CVE 1.1 JSON base meta url")
	pflag.IntVar(&StartYear, "start-year", 2002, "Start year")
	pflag.IntVar(&EndYear, "end-year", time.Now().Year(), "End year")
	pflag.StringVar(&TmpDir, "tmp-dir", "./tmp/", "Output directory")
	pflag.StringVar(&OutputDir, "output-dir", "./out/", "Output directory")
	pflag.Parse()
}
