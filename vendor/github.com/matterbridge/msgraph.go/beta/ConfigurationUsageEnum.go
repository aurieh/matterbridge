// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// ConfigurationUsage undocumented
type ConfigurationUsage int

const (
	// ConfigurationUsageVBlocked undocumented
	ConfigurationUsageVBlocked ConfigurationUsage = 0
	// ConfigurationUsageVRequired undocumented
	ConfigurationUsageVRequired ConfigurationUsage = 1
	// ConfigurationUsageVAllowed undocumented
	ConfigurationUsageVAllowed ConfigurationUsage = 2
)

// ConfigurationUsagePBlocked returns a pointer to ConfigurationUsageVBlocked
func ConfigurationUsagePBlocked() *ConfigurationUsage {
	v := ConfigurationUsageVBlocked
	return &v
}

// ConfigurationUsagePRequired returns a pointer to ConfigurationUsageVRequired
func ConfigurationUsagePRequired() *ConfigurationUsage {
	v := ConfigurationUsageVRequired
	return &v
}

// ConfigurationUsagePAllowed returns a pointer to ConfigurationUsageVAllowed
func ConfigurationUsagePAllowed() *ConfigurationUsage {
	v := ConfigurationUsageVAllowed
	return &v
}