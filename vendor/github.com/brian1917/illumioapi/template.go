package illumioapi

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

// IllumioSecurityTemplateFile is a file with a slice of templates
type IllumioSecurityTemplateFile struct {
	IllumioSecurityTemplates []*IllumioSecurityTemplate `json:"illumio_security_templates"`
}

// IllumioSecurityTemplate contains Labels, IP Lists, Services
type IllumioSecurityTemplate struct {
	Name                  string     `json:"name"`
	Version               int        `json:"version"`
	OsFamily              string     `json:"os_family"`
	Icon                  string     `json:"icon"`
	CompatiblePceVersions []int      `json:"compatible_pce_versions"`
	Labels                []*Label   `json:"labels,omitempty"`
	IPLists               []*IPList  `json:"ip_lists,omitempty"`
	Services              []*Service `json:"services,omitempty"`
}

// ParseTemplateFile imports a JSON template file into the PCE
func ParseTemplateFile(filename string) (IllumioSecurityTemplateFile, error) {
	// Open the file
	jsonFile, err := os.Open(filename)
	if err != nil {
		return IllumioSecurityTemplateFile{}, err
	}
	defer jsonFile.Close()

	// Unmarshal the JSON
	byteValue, _ := ioutil.ReadAll(jsonFile)
	var template []IllumioSecurityTemplateFile
	json.Unmarshal(byteValue, &template)

	return template[0], nil
}
