package mdd

import (
	"strings"
)

// //Dimension 维度接口
// type Dimension interface {
// 	DimID() int64                //TODO:暂时先用int64
// 	DimName() string             //返回维度名称
// 	DimForeignKey() string       //返回维度在事实表中的外键名称
// 	Standard() StandardDimension //返回维度的标准数据

// }

const (
	//DimTypeNotSpeic none
	DimTypeNotSpeic DimType = iota
	//DimTypeAccount 科目
	DimTypeAccount
	//DimTypeTime 时间
	DimTypeTime
	// DimTypeEntity 实体
	DimTypeEntity
)

//DimType 维度类型
type DimType int

//Dimension 维度定义
type Dimension struct {
	DimID            int64  `xorm:"DimId pk"`
	DimForeignKey    string `xorm:"index"` //维度在事实表中的外键名称
	Name             string `xorm:"index"`
	UseStandarMember bool   //是否使用标准成员表
	MemberTable      string `xorm:"index"`
}

//Check 维度检查
func (d Dimension) Check() (bool, string) {
	msg := make([]string, 0)
	if d.DimForeignKey == "" {
		msg = append(msg, "维度没有设置事实表中的外键")
	}
	if d.Name == "" {
		msg = append(msg, "维度没有设置名称")
	}
	if !d.UseStandarMember && d.MemberTable == "" {
		msg = append(msg, "维度既没有使用标准成员表，也没有设置成员表")

	}
	if len(msg) == 0 {
		return true, ""
	}
	return false, strings.Join(msg, ";\r\n")
}

// func (d Dimension)Extend(st interface{}){

// }

// //SaveDimension 保存维度
// func SaveDimension(d Dimension) error {
// 	//TODO: 保存维度
// 	return nil
// }

// /*GetDimMembers 获取维度成员
//  *param d 维度
//  *param cond 成员的过滤条件
//  */
// func GetDimMembers(d Dimension, cond map[string]interface{}) {

// }
