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

type MessageChange struct {

	// Message type
	Type_ string `json:"type,omitempty"`

	// The number of new messages. Can be omitted if the value is zero
	NewCount int32 `json:"newCount,omitempty"`

	// The number of updated messages. Can be omitted if the value is zero
	UpdatedCount int32 `json:"updatedCount,omitempty"`
}
