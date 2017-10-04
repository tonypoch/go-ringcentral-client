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

type ExtensionInfoRequestProvision struct {

	// Mandatory. Resulting extension status
	Status string `json:"status,omitempty"`

	// Mandatory. Extension user contact information
	Contact ExtensionInfoRequestProvisionContactInfo `json:"contact,omitempty"`
}
