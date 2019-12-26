/*
 * Engage Digital API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package engagedigital

import (
	"time"
)

type Intervention struct {
	CategoryIds             []string               `json:"category_ids,omitempty"`
	Closed                  bool                   `json:"closed,omitempty"`
	ClosedAt                time.Time              `json:"closed_at,omitempty"`
	CommentsCount           int32                  `json:"comments_count,omitempty"`
	ContentId               string                 `json:"content_id,omitempty"`
	CreatedAt               time.Time              `json:"created_at,omitempty"`
	CustomFields            map[string]interface{} `json:"custom_fields,omitempty"`
	DeferredAt              time.Time              `json:"deferred_at,omitempty"`
	FirstUserReplyId        string                 `json:"first_user_reply_id,omitempty"`
	FirstUserReplyIn        int32                  `json:"first_user_reply_in,omitempty"`
	FirstUserReplyInBh      int32                  `json:"first_user_reply_in_bh,omitempty"`
	Id                      string                 `json:"id"`
	IdentityId              string                 `json:"identity_id,omitempty"`
	LastUserReplyIn         int32                  `json:"last_user_reply_in,omitempty"`
	LastUserReplyInBh       int32                  `json:"last_user_reply_in_bh,omitempty"`
	SourceId                string                 `json:"source_id,omitempty"`
	Status                  string                 `json:"status,omitempty"`
	ThreadId                string                 `json:"thread_id,omitempty"`
	UpdatedAt               time.Time              `json:"updated_at,omitempty"`
	UserId                  string                 `json:"user_id,omitempty"`
	UserRepliesCount        int32                  `json:"user_replies_count,omitempty"`
	UserReplyInAverage      int32                  `json:"user_reply_in_average,omitempty"`
	UserReplyInAverageBh    int32                  `json:"user_reply_in_average_bh,omitempty"`
	UserReplyInAverageCount int32                  `json:"user_reply_in_average_count,omitempty"`
}
