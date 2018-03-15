package uniplm

type nColumn12 struct {
	GUMILINK_MULTIPLE
}

type NColumn12Mode uint8

const (
	NC12M_APPEND = iota
	NC12M_AUTOSORT_LV1
	NC12M_AUTOSORT_LV2
	NC12M_AUTOSORT_LV3
)
