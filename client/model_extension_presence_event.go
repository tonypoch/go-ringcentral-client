/*
 * RingCentral Connect Platform API Explorer
 *
 * <p>This is a beta interactive API explorer for the RingCentral Connect Platform. To use this service, you will need to have an account with the proper credentials to generate an OAuth2 access token.</p><p><h2>Quick Start</h2></p><ol><li>1) Go to <b>Authentication > /oauth/token</b></li><li>2) Enter <b>app_key, app_secret, username, password</b> fields and then click \"Try it out!\"</li><li>3) Upon success, your access_token is loaded and you can access any form requiring authorization.</li></ol><h2>Links</h2><ul><li><a href=\"https://github.com/ringcentral\" target=\"_blank\">RingCentral SDKs on Github</a></li><li><a href=\"mailto:devsupport@ringcentral.com\">RingCentral Developer Support Email</a></li></ul>
 *
 * API version: 1.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech) DO NOT EDIT.

package ringcentral

type ExtensionPresenceEvent struct {
	// Internal identifier of an extension. Optional parameter
	ExtensionId string `json:"extensionId,omitempty"`
	// Telephony presence status. Returned if telephony status is changed.
	TelephonyStatus string `json:"telephonyStatus,omitempty"`
	// Order number of a notification to state the chronology
	Sequence int32 `json:"sequence,omitempty"`
	// Aggregated presence status, calculated from a number of sources
	PresenceStatus string `json:"presenceStatus,omitempty"`
	// User-defined presence status (as previously published by the user)
	UserStatus string `json:"userStatus,omitempty"`
	// Extended DnD (Do not Disturb) status
	DndStatus string `json:"dndStatus,omitempty"`
	// If 'True' enables other extensions to see the extension presence status
	AllowSeeMyPresence bool `json:"allowSeeMyPresence,omitempty"`
	// If 'True' enables to ring extension phone, if any user monitored by this extension is ringing
	RingOnMonitoredCall bool `json:"ringOnMonitoredCall,omitempty"`
	// If 'True' enables the extension user to pick up a monitored line on hold
	PickUpCallsOnHold bool `json:"pickUpCallsOnHold,omitempty"`
}
