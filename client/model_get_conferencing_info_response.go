/*
 * RingCentral Connect Platform API Explorer
 *
 * <p>This is a beta interactive API explorer for the RingCentral Connect Platform. To use this service, you will need to have an account with the proper credentials to generate an OAuth2 access token.</p><p><h2>Quick Start</h2></p><ol><li>1) Go to <b>Authentication > /oauth/token</b></li><li>2) Enter <b>app_key, app_secret, username, password</b> fields and then click \"Try it out!\"</li><li>3) Upon success, your access_token is loaded and you can access any form requiring authorization.</li></ol><h2>Links</h2><ul><li><a href=\"https://github.com/ringcentral\" target=\"_blank\">RingCentral SDKs on Github</a></li><li><a href=\"mailto:devsupport@ringcentral.com\">RingCentral Developer Support Email</a></li></ul>
 *
 * API version: 1.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech) DO NOT EDIT.

package ringcentral

type GetConferencingInfoResponse struct {
	// Canonical URI of a conferencing
	Uri string `json:"uri,omitempty"`
	// Determines if host user allows conference participants to join before the host
	AllowJoinBeforeHost bool `json:"allowJoinBeforeHost,omitempty"`
	// Access code for a host user
	HostCode string `json:"hostCode,omitempty"`
	// Internal parameter specifying conferencing engine
	Mode string `json:"mode,omitempty"`
	// Access code for any participant
	ParticipantCode string `json:"participantCode,omitempty"`
	// Primary conference phone number for user's home country returned in E.164 (11-digits) format
	PhoneNumber string `json:"phoneNumber,omitempty"`
	// Short URL leading to the service web page Tap to Join for audio conference bridge
	TapToJoinUri string `json:"tapToJoinUri,omitempty"`
	// List of multiple dial-in phone numbers to connect to audio conference service, relevant for user's brand. Each number is given with the country and location information, in order to let the user choose the less expensive way to connect to a conference. The first number in the list is the primary conference number, that is default and domestic
	PhoneNumbers []PhoneNumberInfoConferencing `json:"phoneNumbers,omitempty"`
}
