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

type VersionInfo struct {

	// Canonical URI of API versions
	Uri string `json:"uri,omitempty"`

	// Version of the RingCentral REST API
	VersionString string `json:"versionString,omitempty"`

	// Release date of this version
	ReleaseDate string `json:"releaseDate,omitempty"`

	// URI part determining the current version
	UriString string `json:"uriString,omitempty"`
}
