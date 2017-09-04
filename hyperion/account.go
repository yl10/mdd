package hyperion

import (
	"github.com/yl10/mdd"
)

var _ mdd.Member = Account{}

func init() {

}

const (
	//Expense 费用
	Expense AccountType = iota + 1
	//Revenue 收入
	Revenue
	//Asset 资产
	Asset
	//Liability 责任
	Liability
	//Equity 权益
	Equity
	//Statistical 统计
	Statistical
	//SavedAssumption 假设
	SavedAssumption
)

// AccountType 科目类型
type AccountType int

const (
	//Currency 货币
	Currency AccountDataType = iota + 1
	//NonCurrency 非货币
	NonCurrency
	//Percentage 百分比
	Percentage
)

//AccountDataType 科目数据类型
type AccountDataType int

//DimAccount 科目维度
var DimAccount mdd.Dimension

//Account 科目
type Account struct {
	ID       int64 `xorm:"pk"`
	Type     AccountType
	DataType AccountDataType
	Std      *mdd.StandardMember
}

//GetStandarMember 实现接口
func (a Account) GetStandarMember() mdd.StandardMember {
	return *a.Std
}
