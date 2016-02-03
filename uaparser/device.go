package uaparser

import (
	"regexp"
	"strings"
)

type Device struct {
	Family string
	Brand  string
	Model  string
}

type DevicePattern struct {
	Regexp            *regexp.Regexp
	Regex             string
	RegexFlag         string
	BrandReplacement  string
	DeviceReplacement string
	ModelReplacement  string
}

func (dvcPattern *DevicePattern) Match(line string, dvc *Device) {
	matches := dvcPattern.Regexp.FindStringSubmatch(line)
	if len(matches) == 0 {
		return
	}
	groupCount := dvcPattern.Regexp.NumSubexp()

	if len(dvcPattern.DeviceReplacement) > 0 {
		dvc.Family = allMatchesReplacement(dvcPattern.DeviceReplacement, matches)
	} else if groupCount >= 1 {
		dvc.Family = matches[1]
	}

	if len(dvcPattern.BrandReplacement) > 0 {
		dvc.Brand = allMatchesReplacement(dvcPattern.BrandReplacement, matches)
	} else if groupCount >= 2 {
		dvc.Brand = matches[2]
	}

	if len(dvcPattern.ModelReplacement) > 0 {
		dvc.Model = allMatchesReplacement(dvcPattern.ModelReplacement, matches)
	} else if groupCount >= 3 {
		dvc.Model = matches[3]
	}

	dvc.Family = strings.TrimSpace(dvc.Family)
	dvc.Brand = strings.TrimSpace(dvc.Brand)
	dvc.Model = strings.TrimSpace(dvc.Model)
}

func (dvc *Device) ToString() string {
	var str string
	if dvc.Family != "" {
		str += dvc.Family
	}
	if dvc.Brand != "" {
		str += " " + dvc.Brand
	}
	if dvc.Model != "" {
		str += " " + dvc.Model
	}
	return str
}
