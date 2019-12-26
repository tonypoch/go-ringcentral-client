/*
 * Engage Digital API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package engagedigital

// attributes in extra_infos are optional and unspecified/not guaranteed, don’t rely on it for critical tasks, this is provided as information only
type EventExtraInfo struct {
	AuthenticationStrategy string                 `json:"authentication_strategy,omitempty"`
	BusinessTimeSheetId    string                 `json:"business/time_sheet_id,omitempty"`
	CategoryIds            []string               `json:"category_ids,omitempty"`
	ClosedAutomatically    string                 `json:"closed_automatically,omitempty"`
	ContentId              string                 `json:"content_id,omitempty"`
	ContentSourceId        string                 `json:"content_source_id,omitempty"`
	ContentThreadId        string                 `json:"content_thread_id,omitempty"`
	DeferredDuration       string                 `json:"deferred_duration,omitempty"`
	Deletions              map[string]interface{} `json:"deletions,omitempty"`
	// can be null
	ForeignId      string `json:"foreign_id,omitempty"`
	InterventionId string `json:"intervention_id,omitempty"`
	RequestIp      string `json:"request_ip,omitempty"`
	TaskId         string `json:"task_id,omitempty"`
	UserId         string `json:"user_id,omitempty"`
}
