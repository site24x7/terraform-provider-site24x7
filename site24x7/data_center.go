package site24x7

import "sort"

var dataCenter = map[string]DataCenter{
	"US": {
		displayName:          "United States",
		code:                 "US",
		site24x7APIBaseURL:   "https://www.site24x7.com/api",
		zohoAccountsTokenURL: "https://accounts.zoho.com/oauth/v2/token",
	},
	"EU": {
		displayName:          "Europe",
		code:                 "EU",
		site24x7APIBaseURL:   "https://www.site24x7.eu/api",
		zohoAccountsTokenURL: "https://accounts.zoho.eu/oauth/v2/token",
	},
	"IN": {
		displayName:          "India",
		code:                 "IN",
		site24x7APIBaseURL:   "https://www.site24x7.in/api",
		zohoAccountsTokenURL: "https://accounts.zoho.in/oauth/v2/token",
	},
	"AU": {
		displayName:          "Australia",
		code:                 "AU",
		site24x7APIBaseURL:   "https://www.site24x7.net.au/api",
		zohoAccountsTokenURL: "https://accounts.zoho.com.au/oauth/v2/token",
	},
	"CN": {
		displayName:          "China",
		code:                 "CN",
		site24x7APIBaseURL:   "https://www.site24x7.cn/api",
		zohoAccountsTokenURL: "https://accounts.zoho.com.cn/oauth/v2/token",
	},
	"JP": {
		displayName: "Japan",
		code:        "JP",
		// Japan's API is not served from a www host: www.site24x7.jp is the
		// Japanese marketing site, and the API root is app.site24x7.jp.
		// See https://www.site24x7.com/help/api/
		site24x7APIBaseURL:   "https://app.site24x7.jp/api",
		zohoAccountsTokenURL: "https://accounts.zoho.jp/oauth/v2/token",
	},
	"CA": {
		displayName:          "Canada",
		code:                 "CA",
		site24x7APIBaseURL:   "https://www.site24x7.ca/api",
		zohoAccountsTokenURL: "https://accounts.zohocloud.ca/oauth/v2/token",
	},
	"UK": {
		displayName: "United Kingdom",
		code:        "UK",
		// Like Japan, the documented UK API root is on the app host, not www.
		// See https://www.site24x7.com/help/api/
		site24x7APIBaseURL:   "https://app.site24x7.uk/api",
		zohoAccountsTokenURL: "https://accounts.zoho.uk/oauth/v2/token",
	},
	"AE": {
		displayName: "United Arab Emirates",
		code:        "AE",
		// Like Japan, the documented UAE API root is on the app host, not www.
		// See https://www.site24x7.com/help/api/
		site24x7APIBaseURL:   "https://app.site24x7.ae/api",
		zohoAccountsTokenURL: "https://accounts.zoho.ae/oauth/v2/token",
	},
	"SA": {
		displayName:          "Saudi Arabia",
		code:                 "SA",
		site24x7APIBaseURL:   "https://www.site24x7.sa/api",
		zohoAccountsTokenURL: "https://accounts.zoho.sa/oauth/v2/token",
	},
}

type DataCenter struct {
	displayName          string
	code                 string
	site24x7APIBaseURL   string
	zohoAccountsTokenURL string
}

func (dc *DataCenter) GetAPIBaseURL() string {
	return dc.site24x7APIBaseURL
}

func (dc *DataCenter) GetTokenURL() string {
	return dc.zohoAccountsTokenURL
}

func GetDataCenter(dataCenterCode string) DataCenter {
	return dataCenter[dataCenterCode]
}

// DataCenterCodes returns the supported data center codes in sorted order.
func DataCenterCodes() []string {
	codes := make([]string, 0, len(dataCenter))
	for code := range dataCenter {
		codes = append(codes, code)
	}
	sort.Strings(codes)
	return codes
}
