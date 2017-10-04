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

type InlineResponseDefault35 struct {

	// Canonical URI of the phone numbers resource
	Uri string `json:"uri,omitempty"`

	// List of phone numbers filtered by the specified criteria
	Records []LookUpPhoneNumberPhoneNumberInfo `json:"records,omitempty"`
}
