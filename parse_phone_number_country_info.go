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

type ParsePhoneNumberCountryInfo struct {

	// Internal identifier of a country
	Id string `json:"id,omitempty"`

	// Canonical URI of a country
	Uri string `json:"uri,omitempty"`

	// Country calling code defined by ITU-T recommendations E.123 and E.164, see Calling Codes
	CallingCode string `json:"callingCode,omitempty"`

	// Emergency calling feature availability/emergency address requirement indicator
	EmergencyCalling bool `json:"emergencyCalling,omitempty"`

	// Country code according to the ISO standard, see ISO 3166
	IsoCode string `json:"isoCode,omitempty"`

	// Official name of a country
	Name string `json:"name,omitempty"`
}
