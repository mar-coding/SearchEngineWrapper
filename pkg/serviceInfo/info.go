package serviceInfo

import (
	"fmt"

	"gopkg.in/yaml.v3"
)

// NewFromEmbed read service info from embedded file
func NewFromEmbed(serviceInfo []byte) (*ServiceInfo, error) {
	info := new(ServiceInfo)
	if err := yaml.Unmarshal(serviceInfo, info); err != nil {
		return nil, fmt.Errorf("service info: failed to unmarshal service info file got error %s", err.Error())
	}

	return info, nil
}
