package test

import (
	"fmt"
	"kylin-orm/kylin"
	"kylin-orm/models"
	"testing"
	"time"
)

//TestQueryAll 测试查询全部用例
func TestKylinQueryAll(t *testing.T) {
	kylin := kylin.GetDefaultKylinBase()
	re, err := kylin.QueryAll("dsp_OFFLINE_report", &models.DspOnlineReport{
		AdCreative: "006a3a9e4156c000291a",
		PartTs: models.KylinTime{
			StartTime: time.Now().AddDate(0, -1, 0),
			EndTime:   time.Now(),
		},
		BidFloor: 0.2,
	}, 0, 10)
	fmt.Println(re, err)
}

//TestQueryOne 测试查询一个用例
func TestKylinQueryOne(t *testing.T) {
	kylin := kylin.GetDefaultKylinBase()
	re, err := kylin.QueryOne("dsp_OFFLINE_report", &models.DspOnlineReport{
		AdCreative: "006a3a9e4156c000291a",
		PartTs: models.KylinTime{
			StartTime: time.Now().AddDate(0, -1, 0),
			EndTime:   time.Now(),
		},
	})
	fmt.Println(re, err)
}

//TestQueryPart 测试查询部分
func TestKylinQueryPart(t *testing.T) {
	kylin := kylin.GetDefaultKylinBase()
	re, err := kylin.QueryPart("dsp_OFFLINE_report", []string{"pv", "epv"}, &models.DspOnlineReport{
		AdCreative: "006a3a9e4156c000291a",
		PartTs: models.KylinTime{
			StartTime: time.Now().AddDate(0, -1, 0),
			EndTime:   time.Now(),
		},
	}, 0, 10)
	fmt.Println(re, err)
}

//TestKylinQueryBySQL 单元测试 sql
func TestKylinQueryBySQL(t *testing.T) {
	base := kylin.GetDefaultKylinBase()
	qu := kylin.SQL("DSP_OFFLINE_REPORT")
	w := &models.DspOnlineReport{
		AdCreative: "006a3a9e4156c000291a",
		PartTs: models.KylinTime{
			StartTime: time.Now().AddDate(0, -1, 0),
			EndTime:   time.Now(),
		},
		BidFloor: 0.2,
	}
	sql, err := qu.Select("MEDIA_OS", "MEDIA_BROWSER", "sum(Bid_Floor)").WhereAll(w).Group("MEDIA_OS", "MEDIA_BROWSER").Order("Bid_Floor", "desc", "sum").Limit(10).Build()
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(base.QueryBySQL(sql))
}
