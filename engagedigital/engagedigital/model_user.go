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

type User struct {
	CategoryIds       []string  `json:"category_ids,omitempty"`
	CreatedAt         time.Time `json:"created_at,omitempty"`
	Email             string    `json:"email,omitempty"`
	Enabled           bool      `json:"enabled,omitempty"`
	ExternalId        string    `json:"external_id,omitempty"`
	Firstname         string    `json:"firstname,omitempty"`
	Gender            string    `json:"gender,omitempty"`
	Id                string    `json:"id"`
	IdentityIds       []string  `json:"identity_ids,omitempty"`
	InvitationPending bool      `json:"invitation_pending,omitempty"`
	Lastname          string    `json:"lastname,omitempty"`
	Locale            string    `json:"locale,omitempty"`
	Nickname          string    `json:"nickname,omitempty"`
	RcUserId          string    `json:"rc_user_id,omitempty"`
	RoleId            string    `json:"role_id,omitempty"`
	SpokenLanguages   []string  `json:"spoken_languages,omitempty"`
	TeamIds           []string  `json:"team_ids,omitempty"`
	Timezone          string    `json:"timezone,omitempty"`
	UpdatedAt         time.Time `json:"updated_at,omitempty"`
}
