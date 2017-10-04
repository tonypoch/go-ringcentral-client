/* 
 * Ringcentral API
 *
 * RingCentral Connect Platform API
 *
 * OpenAPI spec version: v1.0
 * 
 * Generated by: https://github.com/swagger-api/swagger-codegen.git
 */

package ringcentral

type InlineResponseDefault13 struct {

	// Canonical URI of a permission resource
	Uri string `json:"uri,omitempty"`

	// Specifies if check result is successful or not
	Successful bool `json:"successful,omitempty"`

	// Information on a permission checked. Returned if successful is 'True'
	Details PermissionDetailsInfo `json:"details,omitempty"`

	// List of active scopes for permission. Returned if successful is 'True'
	Scopes []string `json:"scopes,omitempty"`
}
