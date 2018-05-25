package gumi

type Screen struct {
	Pipeline *Pipeline
	Meta     *Meta
	Resource *Resource
}

func NewScreen() *Screen {
	res := &Screen{}
	res.Pipeline = NewPipeline(res)
	return res
}

//func (s *Screen) Width() int {
//	//return s.RenderingPipeline.Width()
//	return 0
//}
//func (s *Screen) Height() int {
//	//return s.RenderingPipeline.Height()
//	return 0
//}
//func (s *Screen) size() (width, height int) {
//	//return s.RenderingPipeline.size()
//	return 0, 0
//}
//func (s *Screen) Update(info Information) {
//	//s.root.GUMIInfomation(info)
//}
//func (s *Screen) Event(event Event) {
//	//for _, v := range s._hook {
//	//	if v != nil {
//	//		event = v(event)
//	//	}
//	//}
//	if event == nil {
//		return
//	}
//	//s.root.GUMIHappen(event)
//}
//func (s *Screen) Draw() {
//
//}

//
//func (s *Screen) Frame() image.Image {
//	//return s.RenderingPipeline.Image()
//	return nil
//}
//func (s *Screen) RGBA() *image.RGBA {
//	//return s.RenderingPipeline.Image()
//	return nil
//}

//
//func (s *Screen) hookReserve() (id uint64) {
//	defer func() {
//		s._hook[id] = nil
//	}()
//	for {
//		id = rand.Uint64()
//		if id == 0 {
//			continue
//		}
//		if _, ok := s._hook[id]; !ok {
//			return id
//		}
//	}
//}
//func (s *Screen) hookRequest(id uint64, hooking func(event Event) Event) {
//	s._hook[id] = hooking
//}
