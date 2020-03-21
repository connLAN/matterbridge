// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import "context"

// AndroidDeviceOwnerEnterpriseWiFiConfigurationRequestBuilder is request builder for AndroidDeviceOwnerEnterpriseWiFiConfiguration
type AndroidDeviceOwnerEnterpriseWiFiConfigurationRequestBuilder struct{ BaseRequestBuilder }

// Request returns AndroidDeviceOwnerEnterpriseWiFiConfigurationRequest
func (b *AndroidDeviceOwnerEnterpriseWiFiConfigurationRequestBuilder) Request() *AndroidDeviceOwnerEnterpriseWiFiConfigurationRequest {
	return &AndroidDeviceOwnerEnterpriseWiFiConfigurationRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// AndroidDeviceOwnerEnterpriseWiFiConfigurationRequest is request for AndroidDeviceOwnerEnterpriseWiFiConfiguration
type AndroidDeviceOwnerEnterpriseWiFiConfigurationRequest struct{ BaseRequest }

// Get performs GET request for AndroidDeviceOwnerEnterpriseWiFiConfiguration
func (r *AndroidDeviceOwnerEnterpriseWiFiConfigurationRequest) Get(ctx context.Context) (resObj *AndroidDeviceOwnerEnterpriseWiFiConfiguration, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for AndroidDeviceOwnerEnterpriseWiFiConfiguration
func (r *AndroidDeviceOwnerEnterpriseWiFiConfigurationRequest) Update(ctx context.Context, reqObj *AndroidDeviceOwnerEnterpriseWiFiConfiguration) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for AndroidDeviceOwnerEnterpriseWiFiConfiguration
func (r *AndroidDeviceOwnerEnterpriseWiFiConfigurationRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// IdentityCertificateForClientAuthentication is navigation property
func (b *AndroidDeviceOwnerEnterpriseWiFiConfigurationRequestBuilder) IdentityCertificateForClientAuthentication() *AndroidDeviceOwnerCertificateProfileBaseRequestBuilder {
	bb := &AndroidDeviceOwnerCertificateProfileBaseRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/identityCertificateForClientAuthentication"
	return bb
}

// RootCertificateForServerValidation is navigation property
func (b *AndroidDeviceOwnerEnterpriseWiFiConfigurationRequestBuilder) RootCertificateForServerValidation() *AndroidDeviceOwnerTrustedRootCertificateRequestBuilder {
	bb := &AndroidDeviceOwnerTrustedRootCertificateRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/rootCertificateForServerValidation"
	return bb
}