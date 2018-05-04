package verGo

import (
	"fmt"
	"github.com/GUMI-golang/gumi/gcore"
	"math"
	"regexp"
	"strconv"
)

type Version struct {
	Major, Minor, Patch int
}

func (s Version) String() string {
	if s == Invalid {
		return "Invalid"
	}
	mj := fmt.Sprint(s.Major)
	if s.Major == Any {
		mj = "*"
	}
	mn := fmt.Sprint(s.Minor)
	if s.Minor == Any {
		mn = "*"
	}
	pt := fmt.Sprint(s.Patch)
	if s.Patch == Any {
		pt = "*"
	}
	return mj + "." + mn + "." + pt
}
func (s Version) Compatible(o Version) bool {
	if s == Invalid || o == Invalid {
		return false
	}
	if s.Major == o.Major || s.Major == Any || o.Major == Any {
		if s.Minor >= o.Minor || s.Minor == Any || o.Minor == Any {
			return true
		}
	}
	return false
}

var (
	Invalid = Version{math.MinInt32, math.MinInt32, math.MinInt32}
	//Any = Version{-1, -1,-1}
)

const Any = -1

func MakeVersion(major, minor, patch int) Version {
	if major <= Any {
		major = Any
	}
	if minor <= Any {
		minor = Any
	}
	if patch <= Any {
		patch = Any
	}
	return Version{
		Major: major,
		Minor: minor,
		Patch: patch,
	}
}

var re_version_MMP = regexp.MustCompile(`^(?P<major>[+]?[1-9]\d*|0|\*)\.(?P<minor>[+]?[1-9]\d*|0|\*)\.(?P<patch>[+]?[1-9]\d*|0|\*)$`)
var re_version_MM = regexp.MustCompile(`^(?P<major>[+]?[1-9]\d*|0|\*)\.(?P<minor>[+]?[1-9]\d*|0|\*)$`)
var re_version_M = regexp.MustCompile(`^(?P<major>[+]?[1-9]\d*|0|\*)$`)
var re_version_vMMP = regexp.MustCompile(`^v(?P<major>[+]?[1-9]\d*|0|\*)\.(?P<minor>[+]?[1-9]\d*|0|\*)\.(?P<patch>[+]?[1-9]\d*|0|\*)$`)
var re_version_vMM = regexp.MustCompile(`^v(?P<major>[+]?[1-9]\d*|0|\*)\.(?P<minor>[+]?[1-9]\d*|0|\*)$`)
var re_version_vM = regexp.MustCompile(`^v(?P<major>[+]?[1-9]\d*|0|\*)$`)

func ParseVersion(text string) Version {
	if v := re_version_MMP.FindStringSubmatch(text); len(v) > 0 {
		return Version{
			parseVersionInt(v[1]),
			parseVersionInt(v[2]),
			parseVersionInt(v[3]),
		}
	}
	if v := re_version_vMMP.FindStringSubmatch(text); len(v) > 0 {
		return Version{
			parseVersionInt(v[1]),
			parseVersionInt(v[2]),
			parseVersionInt(v[3]),
		}
	}
	if v := re_version_MM.FindStringSubmatch(text); len(v) > 0 {
		return Version{
			parseVersionInt(v[1]),
			parseVersionInt(v[2]),
			Any,
		}
	}
	if v := re_version_vMM.FindStringSubmatch(text); len(v) > 0 {
		return Version{
			parseVersionInt(v[1]),
			parseVersionInt(v[2]),
			Any,
		}
	}
	if v := re_version_M.FindStringSubmatch(text); len(v) > 0 {
		return Version{
			parseVersionInt(v[1]),
			Any,
			Any,
		}
	}
	if v := re_version_vM.FindStringSubmatch(text); len(v) > 0 {
		return Version{
			parseVersionInt(v[1]),
			Any,
			Any,
		}
	}
	return Invalid

}
func parseVersionInt(s string) int {
	if s == "*" {
		return Any
	} else {
		return int(gcore.MustValue(strconv.ParseInt(s, 10, 32)).(int64))
	}
}
