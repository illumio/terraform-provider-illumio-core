package models

// PairingKey represents PairingKey resource
type PairingKey struct {
}

// ToMap - Returns map for PairingKey model
func (pk *PairingKey) ToMap() (map[string]interface{}, error) {
	// pairing key accepts empty object
	return map[string]interface{}{}, nil
}
