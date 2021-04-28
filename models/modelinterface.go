package models

// Model Interface
type Model interface {
	ToMap() (map[string]interface{}, error)
}
