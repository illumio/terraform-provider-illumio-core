// Copyright 2021 Illumio, Inc. All Rights Reserved.

package models

// Model Interface
type Model interface {
	ToMap() (map[string]interface{}, error)
}
