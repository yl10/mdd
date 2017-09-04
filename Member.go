package mdd

const (
	DTYPEUnspecified DATAType = iota
	DTYPECurrency
	DTYPENoncurrency
	DTYPEPercentage
	DTYPEEnum
	DTYPEDate
	DTYPEText
)

//DATAType 成员数据类型
type DATAType int

const (
	StoreData StoreType = iota
	NeverShare
	LabelOnly
	SharedMember
	DynamicCalcStore
	Dynamic
)

//StoreType 数据存储类型
type StoreType int

/*Member 成员接口
 *只要实现下面的接口即可作为成员信息使用
 *GetStandarMember方法，要求返回一个StandardMember
 *如果成员有层级，请按照要求存储在公共的member表中，可在struct中放一个字段关联StandardMember
 *如果成员没有层级，或临时需要的，可在实现方法的时候虚拟够着一个返回
 */
type Member interface {
	GetStandarMember() StandardMember
}

//MemberAttribute 成员属性扩展接口，用于扩展成员
type MemberAttribute interface {
	GetMemberID() int64
}

/*StandardMember 标准成员
带多层级，所有数据都存在一个标准的数据表中
*/
type StandardMember struct {
	ID          int64      `xorm:"id pk"`
	Dim         *Dimension `xorm:"notnull"`
	DataStorage StoreType
	TwopassCalc bool //双向计算
	ParentID    int64
	DataType    DATAType       //数据类型
	Enum        *[]interface{} //当数据类型是枚举的时候有值
	HasFormula  bool           //是否有计算公式

	Consol ConsolationOption //合并计算方式
}

//TableName 实现XORM的接口
func (s StandardMember) TableName() string {
	return "member"
}

// //Key 成员的主键，用接口是为了扩展数据类型
// type Key interface {
// }
