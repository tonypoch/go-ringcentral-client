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

type ContactAddressInfo struct {

	// Country name of extension user company. Not returned for Address Book
	Country string `json:"country,omitempty"`

	// State/province name of extension user company
	State string `json:"state,omitempty"`

	// City name of extension user company
	City string `json:"city,omitempty"`

	// Street address of extension user company
	Street string `json:"street,omitempty"`

	// Zip code of extension user company
	Zip string `json:"zip,omitempty"`
}
