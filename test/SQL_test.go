package test

import (
	"fmt"
	"kylin-orm/kylin"
	"kylin-orm/models"
	"testing"
	"time"
)

//TestSQL 测试构建sql
func TestSQL(t *testing.T) {
	qu := kylin.SQL("DSP_OFFLINE_REPORT")
	w := &models.DspOnlineReport{
		AdCreative: "006a3a9e4156c000291a",
		PartTs: models.KylinTime{
			StartTime: time.Now().AddDate(0, -1, 0),
			EndTime:   time.Now(),
		},
		BidFloor: 0.2,
	}
	sql, err := qu.WhereAll(w).Group("Ad_Creative").Order("Bid_Floor", "desc", "sum").Build()
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("SQL:", sql.String())

}
