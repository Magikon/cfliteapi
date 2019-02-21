package cfstructs

import "time"

//CFConf struct for write/read config file
type CFConf struct {
	Name string `json:"name"`
	Data struct {
		Email    string `json:"email"`
		XAuthKey string `json:"xAuthKey"`
	} `json:"Data"`
}

//AllInfo struct is for return values of request
type AllInfo struct {
	Success  bool          `json:"success"`
	Errors   []interface{} `json:"errors"`
	Messages []interface{} `json:"messages"`
	Result   []struct {
		ID         string    `json:"id"`
		Type       string    `json:"type"`
		Name       string    `json:"name"`
		Content    string    `json:"content"`
		Proxiable  bool      `json:"proxiable"`
		Proxied    bool      `json:"proxied"`
		Priority   int       `json:"priority"`
		TTL        int       `json:"ttl"`
		Locked     bool      `json:"locked"`
		ZoneID     string    `json:"zone_id"`
		ZoneName   string    `json:"zone_name"`
		CreatedOn  time.Time `json:"created_on"`
		ModifiedOn time.Time `json:"modified_on"`
		Meta       struct {
			AutoAdded           bool `json:"auto_added"`
			ManagedByApps       bool `json:"managed_by_apps"`
			ManagedByArgoTunnel bool `json:"managed_by_argo_tunnel"`
		} `json:"meta"`
		Data struct {
		} `json:"data"`
	} `json:"result"`
	ResultInfo struct {
		Page       int `json:"page"`
		PerPage    int `json:"per_page"`
		Count      int `json:"count"`
		TotalCount int `json:"total_count"`
	} `json:"result_info"`
}

