package gumi

import "github.com/GUMI-golang/gumi/gcore"

func saveGUMIChildrun(dst *[]GUMI, mode gcore.Mode, index gcore.Index, src ...GUMI) (writen int) {
	var ln = len(*dst)
	var idx int
	switch mode {
	default:
		fallthrough
	case gcore.REPLACE:
		for i, v := range src {
			idx = (index + gcore.Index(i)).Indexize(ln)
			if idx == gcore.IndexNotExist {
				break
			}
			(*dst)[idx] = v
			writen++
		}
	case gcore.PUSHBACKWARD:
		idx = (index + gcore.Index(1)).Indexize(ln)
		if idx == gcore.IndexNotExist {
			(*dst) = append(
				(*dst),
				src...,
			)
		} else {
			dstforw := (*dst)[:idx]
			dstback := (*dst)[idx:]
			(*dst) = append(
				dstforw,
				append(
					src,
					dstback...,
				)...,
			)
		}
	case gcore.PUSHONWARD:
		idx = index.Indexize(ln)
		dstforw := (*dst)[:idx]
		dstback := (*dst)[idx:]
		(*dst) = append(
			dstforw,
			append(
				src,
				dstback...,
			)...,
		)
		writen = len(src)
	}
	return writen
}
func loadGUMIChildrun(dst []GUMI, index gcore.Index, count int) (res []GUMI) {
	var ln = len(dst)
	var idx = index.Indexize(ln)
	var sz int
	if idx == gcore.IndexNotExist {
		return
	}
	if idx+count <= ln {
		sz = count
	} else {
		sz = ln - idx
	}
	res = make([]GUMI, sz)
	for i := 0; i < sz; i++ {
		res[i] = dst[idx+i]
	}
	return res
}
