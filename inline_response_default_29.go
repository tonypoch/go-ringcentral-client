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

type InlineResponseDefault29 struct {

	// List of countries with the country data
	Records []FullCountryInfo `json:"records,omitempty"`

	// Information on navigation
	Navigation NavigationInfo `json:"navigation,omitempty"`

	// Information on paging
	Paging PagingInfo `json:"paging,omitempty"`
}
