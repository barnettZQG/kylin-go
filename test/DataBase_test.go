package test

import (
	"fmt"
	"kylin-orm/kylin"
	"kylin-orm/models"
	"testing"
	"time"
)

//TestQueryAll 测试查询全部用例
func TestQueryAll(t *testing.T) {
	db := &kylin.DataBase{
		ProjectName: "dsp_online_test",
	}
	// day, err := time.ParseDuration("24h")
	// if err != nil {
	// 	fmt.Println(err)
	// }
	re, err := db.QueryAll("dsp_OFFLINE_report", &models.DspOnlineReport{
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
func TestQueryOne(t *testing.T) {
	db := &kylin.DataBase{
		ProjectName: "dsp_online_test",
	}

	re, err := db.QueryOne("dsp_OFFLINE_report", &models.DspOnlineReport{
		AdCreative: "006a3a9e4156c000291a",
		PartTs: models.KylinTime{
			StartTime: time.Now().AddDate(0, -1, 0),
			EndTime:   time.Now(),
		},
	})
	fmt.Println(re, err)
}

//TestQueryPart 测试查询部分
func TestQueryPart(t *testing.T) {
	db := &kylin.DataBase{
		ProjectName: "dsp_online_test",
	}

	re, err := db.QueryPart("dsp_OFFLINE_report", []string{"pv", "epv"}, &models.DspOnlineReport{
		AdCreative: "006a3a9e4156c000291a",
		PartTs: models.KylinTime{
			StartTime: time.Now().AddDate(0, -1, 0),
			EndTime:   time.Now(),
		},
	}, 0, 10)
	fmt.Println(re, err)
}

func TestQueryBySQL(t *testing.T) {
	db := &kylin.DataBase{
		ProjectName: "dsp_online_test",
	}
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
	fmt.Println(db.QueryBySQL(sql))
}
