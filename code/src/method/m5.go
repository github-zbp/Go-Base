package method

type SpecialPoint struct{
	Point
	size int
}

func (sp *SpecialPoint) GetSize() int
func (sp *SpecialPoint) SetSize(size int)