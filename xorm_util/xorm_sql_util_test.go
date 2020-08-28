package mysql_util

import (
	"fmt"
	"testing"
	"time"

	"github.com/emirpasic/gods/sets/hashset"
)

type TQywxExternalContactExtInfo struct {
	FId             uint64    `json:"f_id" xorm:"not null pk autoincr comment('主键') BIGINT(20) 'f_id'"`
	FUserid         string    `json:"f_userid" xorm:"not null pk default '' comment('企业服务人员的userid') VARCHAR(64) 'f_userid'"`
	FExternalUserid string    `json:"f_external_userid" xorm:"not null default '' comment('外部联系人的userid') VARCHAR(64) 'f_external_userid'"`
	FType           uint32    `json:"f_type" xorm:"not null default 0 comment('信息类型 0-未知，1-备注，2-备注手机号，3-三级项目id标签, 4-fr值') INT(10) 'f_type'"`
	FValue          string    `json:"f_value" xorm:"not null default '' comment('信息内容') VARCHAR(256) 'f_value'"`
	FCreateTime     time.Time `json:"f_create_time" xorm:"not null default 'current_timestamp()' comment('创建时间') TIMESTAMP 'f_create_time'"`
	FUpdateTime     time.Time `json:"f_update_time" xorm:"not null default 'current_timestamp()' comment('更新时间') TIMESTAMP 'f_update_time'"`
	FDelFlag        uint32    `json:"f_del_flag" xorm:"not null default 0 comment('删除标记，0正常，1删除') TINYINT(4) 'f_del_flag'"`
}

func TestGenUpsertSQLAndParams(t *testing.T) {
	info := TQywxExternalContactExtInfo{
		FUserid:         "",
		FExternalUserid: "",
		FType:           0,
		FValue:          "",
		FDelFlag:        0,
	}

	sql, params := GenUpsertSQLAndParams(info, hashset.New("f_id", "f_create_time", "f_update_time"), hashset.New("f_id", "f_create_time", "f_update_time"))
	fmt.Printf("sql=%s\n", sql)
	fmt.Printf("params=%v\n", params)
}
