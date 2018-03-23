package material

type mtColorSingle struct {
	mcl1 *MaterialColor
}

func (s *mtColorSingle) GetMaterialColor() *MaterialColor {
	return s.mcl1
}
func (s *mtColorSingle) SetMaterialColor(mc *MaterialColor) {
	s.mcl1 = mc
}

type mtColorFromTo struct {
	mcl1 *MaterialColor
	mcl2 *MaterialColor
}

func (s *mtColorFromTo) GetFromMaterialColor() *MaterialColor {
	return s.mcl1
}
func (s *mtColorFromTo) SetFromMaterialColor(mc *MaterialColor) {
	s.mcl1 = mc
}
func (s *mtColorFromTo) GetToMaterialColor() *MaterialColor {
	return s.mcl2
}
func (s *mtColorFromTo) SetToMaterialColor(mc *MaterialColor) {
	s.mcl2 = mc
}
