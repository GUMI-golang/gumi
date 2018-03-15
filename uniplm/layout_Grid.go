package uniplm

import "image"

type nGrid struct {
	GUMILINK_MULTIPLE
	Node gridNode
}

func (s *nGrid) size(drawing *Drawing, style *Style) Size {
	return Size{
		Horizontal: AUTOLENGTH,
		Vertical:   AUTOLENGTH,
	}
}
func (s *nGrid) draw(drawing *Drawing, style *Style, frame Frame) {
	var scn = make(chan struct{}, 2)
	var cnt = len(s.child) - 1
	temp := frame.Bounds()
	var ok, wls, hls = s.Node.LoadCache(temp.Dx(), temp.Dy())
	if !ok {
		s.Node.StoreCache(temp.Dx(), temp.Dy())
		_, wls, hls = s.Node.LoadCache(temp.Dx(), temp.Dy())
	}
	for y := 0; y < s.Node.hcount; y++ {
		for x := 0; x < s.Node.wcount; x++ {
			i := s.Node.Indexize(x, y)
			if i <= cnt {
				go func(ai, ax, ay int) {
					var stx, edx, sty, edy int
					if ax-1 < 0 {
						stx = 0
					} else {
						stx = wls[ax-1]
					}
					edx = wls[ax]
					if ay-1 < 0 {
						sty = 0
					} else {
						sty = hls[ay-1]
					}
					edy = hls[ay]
					//
					s.child[ai].(GUMIElem).draw(drawing, style, frame.SubFrame(image.Rect(stx, sty, edx, edy)))
					if ai == cnt {
						close(scn)
					}
				}(i, x, y)
			}
		}
	}

	for range scn {
	}
}

func NGrid(wc, hc int, childrun ...GUMILinker) *nGrid {
	temp := &nGrid{
		Node: GridNode(wc, hc),
	}
	temp.Link(childrun...)
	return temp
}
