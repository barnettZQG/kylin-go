package models

import (
	"time"
)

//DspOnlineReport dsp数据模型
//model 实例，支持类型包括：string,int,Time,float64
//kylin:"necessary"代表此字段必须指定，才能查询
type DspOnlineReport struct {
	AdCreative     string    `json:"ad_creative" kylin:"necessary"`
	AdCategroy     string    `json:"ad_categroy"`
	MediaID        string    `json:"media_id"`
	MediaSlot      string    `json:"media_slot"`
	MediaType      string    `json:"media_type"`
	MediaTrafficID string    `json:"media_traffic_id"`
	MediaOemID     string    `json:"media_oem_id"`
	MediaOS        string    `json:"media_os"`
	MediaBrowser   string    `json:"media_browser"`
	MediaProvince  string    `json:"media_province"`
	MediaCity      string    `json:"media_city"`
	MediaDomain    string    `json:"media_domain"`
	MediaApp       string    `json:"media_app"`
	PartTs         time.Time `json:"part_ts" kylin:"necessary"`
	BidFloor       float64   `json:"bid_floor"`
	BidPrice       float64   `json:"bid_price"`
	Ereq           int       `json:"ereq"`
	Request        int       `json:"request"`
	OemPrice       float64   `json:"oem_price"`
	AdxPrice       float64   `json:"adx_price"`
	PV             int       `json:"pv"`
	Epv            int       `json:"epv"`
	UV             int       `json:"uv"`
	IP             int       `json:"ip"`
	Price          float64   `json:"price"`
	Money          float64   `json:"money"`
	Click          int       `json:"click"`
	Eclick         int       `json:"eclick"`
	Arrive         int       `json:"arrive"`
	Stay           int       `json:"stay"`
	Conversion     int       `json:"conversion"`
	Register       int       `json:"register"`
}
