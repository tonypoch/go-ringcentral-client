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

type TimezoneInfo struct {

	// Internal identifier of a timezone
	Id string `json:"id,omitempty"`

	// Canonical URI of a timezone
	Uri string `json:"uri,omitempty"`

	// Short name of a timezone
	Name string `json:"name,omitempty"`

	// Meaningful description of the timezone
	Description string `json:"description,omitempty"`
}
