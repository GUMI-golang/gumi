package gcore

import "sort"

func init() {
	sort.Sort(DefinedResolutions)
}

var DefinedResolutions = STDResolutions([]STDResolution{
	// CIF
	{[]string{"Sub-QCIF"}, 128, 96},
	{[]string{"QCIF"}, 220, 176},
	{[]string{"CIF"}, 352, 288},
	// CGA
	{[]string{"CGA"}, 320, 200},
	// VGA
	{[]string{"qVGA"}, 320, 240},
	{[]string{"HVGA"}, 480, 320},
	{[]string{"VGA"}, 640, 480},
	{[]string{"SVGA"}, 800, 600},
	{[]string{"WSVGA"}, 1024, 576},
	// XGA
	{[]string{"XGA"}, 1024, 768},
	{[]string{"FWXGA"}, 1366, 768},
	{[]string{"SXGA"}, 1280, 1024},
	{[]string{"UXGA"}, 1600, 1200},
	// HD
	{[]string{"nHD"}, 640, 360},
	{[]string{"qHD"}, 960, 540},
	{[]string{"HD", "720p"}, 1280, 720},
	{[]string{"HD+"}, 1600, 900},
	{[]string{"FHD", "1080p", "1080i"}, 1920, 1080},
	{[]string{"QHD", "2K"}, 2560, 1440},
	{[]string{"UHD", "4K"}, 3840, 2160},
	{[]string{"UHD+", "5K"}, 5120, 2880},
	{[]string{"QUHD", "8K"}, 7680, 4320},
})

// sorting util
type STDResolution struct {
	Name   []string
	Width  int
	Height int
}
type STDResolutions []STDResolution

func (s STDResolutions) Len() int {
	return len(([]STDResolution)(s))
}
func (s STDResolutions) Swap(i, j int) {
	temp := ([]STDResolution)(s)
	temp[i], temp[j] = temp[j], temp[i]
}
func (s STDResolutions) Less(i, j int) bool {
	temp := ([]STDResolution)(s)
	return temp[i].Width*temp[i].Height < temp[j].Width*temp[j].Height
}
func (s STDResolutions) Smaller(w, h int) (res []STDResolution) {
	for _, v := range s {
		if v.Width <= w && v.Height <= h {
			res = append(res, v)
		}
	}
	return res
}
func (s STDResolutions) Get(name string) (w, h int) {
	for _, v := range ([]STDResolution)(s) {
		if exist(v.Name, name) {
			return v.Width, v.Height
		}
	}
	return 0, 0
}
func (s STDResolutions) Kinds() []string {
	tmp := ([]STDResolution)(s)
	temp := make([]string, len(tmp))
	for i, v := range tmp {
		temp[i] = v.Name[0]
	}
	return temp
}
func (s STDResolutions) Under(w, h int) []string {
	var temp []string
	for _, v := range ([]STDResolution)(s) {
		if v.Width <= w && v.Height < h {
			temp = append(temp, v.Name[0])
		}
	}
	return temp
}

func exist(testset []string, val string) bool {
	for _, v := range testset {
		if v == val {
			return true
		}
	}
	return false
}
