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

type ModelInfo struct {

	// Device model identifier. Mandatory when ordering a HardPhone if boxBillingId is not used for ordering
	Id string `json:"id,omitempty"`

	// Device name
	Name string `json:"name,omitempty"`

	// Addons description
	Addons []AddonInfo `json:"addons,omitempty"`
}
