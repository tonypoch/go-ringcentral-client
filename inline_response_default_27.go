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

type InlineResponseDefault27 struct {

	// Internal identifier of the call recording
	Id string `json:"id,omitempty"`

	// Link to the call recording binary content
	ContentUri string `json:"contentUri,omitempty"`

	// Call recording file format. Supported format is audio/x-wav
	ContentType string `json:"contentType,omitempty"`

	// Recorded call duration
	Duration int32 `json:"duration,omitempty"`
}
