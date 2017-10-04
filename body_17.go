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

type Body17 struct {

	// Custom data access key. Optional. If specified, must match the custom key in the URL
	Id string `json:"id,omitempty"`

	// Description of custom data. Mandatory for create, if there is no attachment specified. Maximum length is limited to 256 symbols
	Value string `json:"value,omitempty"`
}
