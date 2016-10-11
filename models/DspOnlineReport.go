package models

import (
	"time"
)

//DspOnlineReport dsp数据模型
//model 实例，支持类型包括：string,int,KylinTime,float64
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
	PartTs         KylinTime `json:"part_ts" kylin:"necessary"`
	BidFloor       float64   `json:"bid_floor omitempty" kylin:"nowhere"`
	BidPrice       float64   `json:"bid_price omitempty" kylin:"nowhere"`
	Ereq           int       `json:"ereq omitempty" kylin:"nowhere"`
	Request        int       `json:"request omitempty" kylin:"nowhere"`
	OemPrice       float64   `json:"oem_price omitempty" kylin:"nowhere"`
	AdxPrice       float64   `json:"adx_price omitempty" kylin:"nowhere"`
	PV             int       `json:"pv omitempty" kylin:"nowhere"`
	Epv            int       `json:"epv omitempty" kylin:"nowhere"`
	UV             int       `json:"uv omitempty" kylin:"nowhere"`
	IP             int       `json:"ip omitempty" kylin:"nowhere"`
	Price          float64   `json:"price omitempty" kylin:"nowhere"`
	Money          float64   `json:"money omitempty" kylin:"nowhere"`
	Click          int       `json:"click omitempty" kylin:"nowhere"`
	Eclick         int       `json:"eclick omitempty" kylin:"nowhere"`
	Arrive         int       `json:"arrive omitempty" kylin:"nowhere"`
	Stay           int       `json:"stay omitempty" kylin:"nowhere"`
	Conversion     int       `json:"conversion omitempty" kylin:"nowhere"`
	Register       int       `json:"register omitempty" kylin:"nowhere"`
}

//KylinTime kylin-orm对于时间的处理类型
type KylinTime struct {
	StartTime time.Time `json:"start_time"`
	EndTime   time.Time `json:"end_time"`
}

//IsZero 判断是否为0值
func (kt KylinTime) IsZero() bool {
	return kt.StartTime.IsZero() && kt.EndTime.IsZero()
}
