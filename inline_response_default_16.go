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

type InlineResponseDefault16 struct {

	// List of call log records with synchronization information. For ISync the total number of returned records is limited to 250; this includes both new records and the old ones, specified by the recordCount parameter
	Records []CallLogRecord `json:"records,omitempty"`

	// Sync type, token and time
	SyncInfo SyncInfo `json:"syncInfo,omitempty"`
}
