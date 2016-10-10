package kylin

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
			StartTime: time.Now().AddDate(0, 0, -2),
			EndTime:   time.Now(),
		},
	})
	fmt.Println(re, err)
}
