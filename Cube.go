package mdd

//Slice 切片
type Slice struct {
	X        []Dimension
	Y        []Dimension
	XMembers []*Member
	YMembers []*Member
}

//Dice 切块 实际也是一个立方体
type Dice Cube

//Cube 立方体
type Cube struct {
	Name     string
	x        []Dimension
	y        []Dimension
	z        []Dimension
	XMembers []interface{}
	YMembers []interface{}
	ZMembers []interface{}
}

//NewCube 创建一个立方体
func NewCube(name string, ftable []FactTable, x, y, z []Dimension) *Cube {
	c := &Cube{
		Name: name,
		x:    x,
		y:    y,
		z:    z,
	}
	// keys := ftable[0].GetDimension()

	if len(x) > 1 {

	}
	xl := len(x)
	xrow = make([]*Member, xl)

	return c
}

//CubeDimension 立方体维度
type CubeDimension struct {
	Hierarchy int
	Level     int
}

//CubeMember 立方体成员
type CubeMember struct {
}

// //Slice 对立方体进行切片操作
// func (c Cube) Slice(mbr *Members) (Slice, error) {

// }
