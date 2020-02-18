package behavior

type Nums struct {
	n1 int
	n2 int
}

//接口是包含一组方法定义，表示行为
type Calculate interface {
	add() int
	minute() int
	multiply() int
	divide() int
	mod() int
}

func (num Nums) add() int {
	return num.n1 + num.n2
}

func (num Nums) minute() int  {
	return num.n1 - num.n2
}

func (num Nums) multiply() int  {
	return num.n1 * num.n2
}

func (num Nums) divide() int  {
	return num.n1 / num.n2
}

func (num Nums) mod() int  {
	return num.n1 % num.n2
}