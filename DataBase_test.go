package kylin

import (
	"fmt"
	"kylin-orm/kylin"
	"kylin-orm/models"
	"testing"
	"time"
)

func TestQueryAll(t *testing.T) {
	db := &kylin.DataBase{}
	body, err := db.QueryAll("dsp_online_test", "dsp_online_report", &models.DspOnlineReport{
		AdCreative: "001b2c21c1573bfc5481",
		BidFloor:   0.23,
		PartTs:     time.Now(),
	}, 0, 10)
	fmt.Println(string(body), err)
}
