// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// SecurityProviderStatus undocumented
type SecurityProviderStatus struct {
	// Object is the base model of SecurityProviderStatus
	Object
	// Enabled undocumented
	Enabled *bool `json:"enabled,omitempty"`
	// Endpoint undocumented
	Endpoint *string `json:"endpoint,omitempty"`
	// Provider undocumented
	Provider *string `json:"provider,omitempty"`
	// Region undocumented
	Region *string `json:"region,omitempty"`
	// Vendor undocumented
	Vendor *string `json:"vendor,omitempty"`
}