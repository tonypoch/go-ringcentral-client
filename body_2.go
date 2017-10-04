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

type Body2 struct {

	// Must hold password value for Resource Owner Credentials flow. If client application is not authorized by the specified grant_type, response does not contain refresh_token and refresh_token_ttl attributes
	GrantType string `json:"grant_type,omitempty"`

	// Optional. Access token lifetime in seconds; the possible values are from 600 sec (10 min) to 3600 sec (1 hour). The default value is 3600 sec. If the value specified exceeds the default one, the default value is set. If the value specified is less than 600 seconds, the minimum value (600 sec) is set
	AccessTokenTtl int32 `json:"access_token_ttl,omitempty"`

	// Optional. Refresh token lifetime in seconds. The default value depends on the client application, but as usual it equals to 7 days. If the value specified exceeds the default one, the default value is applied. If client specifies refresh_token_ttl<=0, refresh token is not returned even if the corresponding grant type is supported
	RefreshTokenTtl int32 `json:"refresh_token_ttl,omitempty"`

	// Phone number linked to account or extension in account in E.164 format with or without leading \"+\" sign
	Username string `json:"username,omitempty"`

	// Optional. Extension short number. If company number is specified as a username, and extension is not specified, the server will attempt to authenticate client as main company administrator
	Extension string `json:"extension,omitempty"`

	// User's password
	Password string `json:"password,omitempty"`

	// Optional. List of API permissions to be used with access token (see Application Permissions). Can be omitted when requesting all permissions defined during the application registration phase
	Scope string `json:"scope,omitempty"`

	// Optional. Unique identifier of a client application. You can pass it in request according to pattern [a-zA-Z0-9_\\-]{1,64}. Otherwise it is auto-generated by server. The value will be returned in response in both cases
	EndpointId string `json:"endpoint_id,omitempty"`
}
