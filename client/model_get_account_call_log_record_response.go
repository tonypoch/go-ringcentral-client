/*
 * RingCentral Connect Platform API Explorer
 *
 * <p>This is a beta interactive API explorer for the RingCentral Connect Platform. To use this service, you will need to have an account with the proper credentials to generate an OAuth2 access token.</p><p><h2>Quick Start</h2></p><ol><li>1) Go to <b>Authentication > /oauth/token</b></li><li>2) Enter <b>app_key, app_secret, username, password</b> fields and then click \"Try it out!\"</li><li>3) Upon success, your access_token is loaded and you can access any form requiring authorization.</li></ol><h2>Links</h2><ul><li><a href=\"https://github.com/ringcentral\" target=\"_blank\">RingCentral SDKs on Github</a></li><li><a href=\"mailto:devsupport@ringcentral.com\">RingCentral Developer Support Email</a></li></ul>
 *
 * API version: 1.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech) DO NOT EDIT.

package ringcentral

import (
	"time"
)

type GetAccountCallLogRecordResponse struct {
	// Internal identifier of a cal log record
	Id string `json:"id,omitempty"`
	// Canonical URI of a call log record
	Uri string `json:"uri,omitempty"`
	// Internal identifier of a call session
	SessionId string            `json:"sessionId,omitempty"`
	From      CallLogCallerInfo `json:"from,omitempty"`
	To        CallLogCallerInfo `json:"to,omitempty"`
	// Call type
	Type string `json:"type,omitempty"`
	// Call direction
	Direction string `json:"direction,omitempty"`
	// Action description of the call operation
	Action string `json:"action,omitempty"`
	// Status description of the call operation
	Result string `json:"result,omitempty"`
	// The call start datetime in ISO 8601 format including timezone, for example 2016-03-10T18:07:52.534Z
	StartTime time.Time `json:"startTime,omitempty"`
	// Call duration in seconds
	Duration  int32         `json:"duration,omitempty"`
	Recording RecordingInfo `json:"recording,omitempty"`
}
