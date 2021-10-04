package site24x7

var dataCenter = map[string]DataCenter{
	"US": {
		displayName:          "United States",
		code:                 "US",
		site24x7APIBaseURL:   "https://www.site24x7.com/api",
		zohoAccountsTokenURL: "https://accounts.zoho.com/oauth/v2/token",
	},
	"EU": {
		displayName:        "Europe",
		code:               "EU",
		site24x7APIBaseURL: "https://www.site24x7.eu/api",
		zohoAccountsTokenURL: "https://accounts.zoho.eu/oauth/v2/toke	n",
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
		site24x7APIBaseURL:   "https://www.site24x7.eu/api",
		zohoAccountsTokenURL: "https://accounts.zoho.com.cn/oauth/v2/token",
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
