// Copyright 2021 Illumio, Inc. All Rights Reserved.

package models

type VENsUpgrade struct {
	Hrefs   []Href `json:"vens"`
	Release string `json:"release"`
}

func (o *VENsUpgrade) ToMap() (map[string]interface{}, error) {
	return map[string]interface{}{
		"vens":    GetHrefMaps(o.Hrefs),
		"release": o.Release,
	}, nil
}
