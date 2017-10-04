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

type RuleInfoForwardingNumberInfo struct {

	// Link to a forwarding number resource
	Uri string `json:"uri,omitempty"`

	// Internal identifier of a forwarding number
	Id string `json:"id,omitempty"`

	// Phone number to which the call is forwarded
	PhoneNumber string `json:"phoneNumber,omitempty"`

	// Title of a forwarding number
	Label string `json:"label,omitempty"`
}
