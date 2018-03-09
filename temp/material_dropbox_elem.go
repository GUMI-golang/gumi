package temp

type mtDropboxElemList []*mtDropboxElem
type mtDropboxElem struct {
	content string
	w, h    int
}
func newMTDropboxElem(content string) *mtDropboxElem {
	return &mtDropboxElem{
		content: content,
		w:       -1,
		h:       -1,
	}
}
func (s *mtDropboxElemList) Set(idx int, content string) {
	s.extend(idx)
	(*s)[idx] = newMTDropboxElem(content)
}
func (s *mtDropboxElemList) extend(idx int) {
	more := idx - s.Length() + 1
	if more > 0 {
		*s = append(*s, make([]*mtDropboxElem, more)...)
	}
}
func (s *mtDropboxElemList) Get(idx int) string {
	if s.Exist(idx) {
		return (*s)[idx].content
	}
	return ""
}
func (s *mtDropboxElemList) Exist(idx int) bool {
	return 0 <= idx && idx < len(*s)
}
func (s *mtDropboxElemList) Length() int {
	return len(*s)
}
//
func (s *mtDropboxElemList) update(sty *Style) {
	sty.Default.Font.Use()
	defer sty.Default.Font.Release()
	for _, v := range *s {
		v.w, v.h = sty.Default.Font.CalculateSize(v.content)
	}
}
func (s *mtDropboxElemList) needUpdate() bool {
	for _, v := range *s {
		if v.w < 0 || v.h < 0 {
			return true
		}
	}
	return false
}
func (s *mtDropboxElemList) heightSum() (res int) {
	for _, v := range *s {
		res += v.h
	}
	return res
}
func (s *mtDropboxElemList) refer() []*mtDropboxElem {
	return ([]*mtDropboxElem)(*s)
}
func (s *mtDropboxElemList) getForDraw(idx int) *mtDropboxElem {
	if s.Exist(idx) {
		return (*s)[idx]
	}
	return &mtDropboxElem{
		content: "",
		h:       0,
		w:       0,
	}
}
