package mdd

/*
 *处理合并方式
 */
const (
	//CONSOLAdd +
	CONSOLAdd ConsolationOption = iota
	//CONSOLSubtract -
	CONSOLSubtract
	//CONSOLMultiply *
	CONSOLMultiply
	//CONSOLDivide /
	CONSOLDivide
	// CONSOLPercentage %
	CONSOLPercentage
	//CONSOLIngore ~
	CONSOLIngore
)

//ConsolationOption 合并方式
type ConsolationOption int
