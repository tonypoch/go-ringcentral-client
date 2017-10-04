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

type ExtensionPermissions struct {

	// Admin permission
	Admin PermissionInfo `json:"admin,omitempty"`

	// International Calling permission
	InternationalCalling PermissionInfo `json:"internationalCalling,omitempty"`
}
