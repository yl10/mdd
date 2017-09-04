package mdd

//FactTable 事实表
type FactTable interface {
	GetDimension() map[string]Dimension //实现事实表和维度的映射关系
}

// func FactTO(details []FactTable, x, y, z []Dimension) Cube {

// 	cb:=Cube{

// 	}
// }
