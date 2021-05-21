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
	Description string `json:"description"`
}

func (cc *ContainerCluster) ToMap() (map[string]interface{}, error) {
	ccMap := make(map[string]interface{})

	ccMap["name"] = cc.Name
	ccMap["description"] = cc.Description

	return ccMap, nil
}
