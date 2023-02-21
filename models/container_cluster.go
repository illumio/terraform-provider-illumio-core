// Copyright 2021 Illumio, Inc. All Rights Reserved.

package models

// Sample
/*
{
  "name": "string",
  "description": "string"
}
*/

type ContainerCluster struct {
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
}

func (cc *ContainerCluster) ToMap() (map[string]interface{}, error) {
	return toMap(cc)
}
