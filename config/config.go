package config

import (
	"strings"

	"github.com/spf13/pflag"
)

var (
	CveUrlPriffix        string
	Cve11ModifiedJsonUrl string
	Cve11RecentJsonUrl   string
	Cve11BaseJsonUrl     string
	Cve11ModifiedMetaUrl string
	Cve11RecentMetaUrl   string
	Cve11BaseMetaUrl     string
	RetireJsUrl          string
	StartYear            int
	TmpDir               string
	OutputDir            string
	Cron                 string
	Addr                 string
)

func init() {
	pflag.StringVar(&CveUrlPriffix, "nvd", "https://nvd.nist.gov/feeds/json/cve/1.1/", "NVD CVE 1.1 JSON base url")
	pflag.StringVar(&Cve11ModifiedJsonUrl, "cve-modified-json-url", "nvdcve-1.1-modified.json.gz", "NVD CVE 1.1 JSON modified url")
	pflag.StringVar(&Cve11RecentJsonUrl, "cve-recent-json-url", "nvdcve-1.1-recent.json.gz", "NVD CVE 1.1 JSON recent url")
	pflag.StringVar(&Cve11BaseJsonUrl, "cve-base-json-url", "nvdcve-1.1-%d.json.gz", "NVD CVE 1.1 JSON base url")
	pflag.StringVar(&Cve11ModifiedMetaUrl, "cve-modified-meta-url", "nvdcve-1.1-modified.meta", "NVD CVE 1.1 JSON modified meta url")
	pflag.StringVar(&Cve11RecentMetaUrl, "cve-recent-meta-url", "nvdcve-1.1-recent.meta", "NVD CVE 1.1 JSON recent meta url")
	pflag.StringVar(&Cve11BaseMetaUrl, "cve-base-meta-url", "nvdcve-1.1-%d.meta", "NVD CVE 1.1 JSON base meta url")
	pflag.StringVar(&RetireJsUrl, "retire-json-url", "https://cdn.jsdelivr.net/gh/Retirejs/retire.js@master/repository/jsrepository.json", "retrieJs repository url")
	pflag.IntVar(&StartYear, "start-year", 2002, "Start year")
	pflag.StringVar(&TmpDir, "tmp-dir", "./tmp/nvd", "Output directory")
	pflag.StringVar(&OutputDir, "output-dir", "./data", "Output directory")
	pflag.StringVar(&Cron, "cron", "0 */4 * * *", "Cron expression")
	pflag.StringVar(&Addr, "addr", ":80", "Listen address")
	pflag.Parse()
	if !strings.HasPrefix(Cve11ModifiedMetaUrl, "http") {
		Cve11ModifiedMetaUrl = CveUrlPriffix + Cve11ModifiedMetaUrl
	}
	if !strings.HasPrefix(Cve11RecentMetaUrl, "http") {
		Cve11RecentMetaUrl = CveUrlPriffix + Cve11RecentMetaUrl
	}
	if !strings.HasPrefix(Cve11BaseMetaUrl, "http") {
		Cve11BaseMetaUrl = CveUrlPriffix + Cve11BaseMetaUrl
	}
	if !strings.HasPrefix(Cve11ModifiedJsonUrl, "http") {
		Cve11ModifiedJsonUrl = CveUrlPriffix + Cve11ModifiedJsonUrl
	}
	if !strings.HasPrefix(Cve11RecentJsonUrl, "http") {
		Cve11RecentJsonUrl = CveUrlPriffix + Cve11RecentJsonUrl
	}
	if !strings.HasPrefix(Cve11BaseJsonUrl, "http") {
		Cve11BaseJsonUrl = CveUrlPriffix + Cve11BaseJsonUrl
	}

}
