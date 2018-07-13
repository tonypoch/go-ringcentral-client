/*
 * RingCentral Connect Platform API Explorer
 *
 * <p>This is a beta interactive API explorer for the RingCentral Connect Platform. To use this service, you will need to have an account with the proper credentials to generate an OAuth2 access token.</p><p><h2>Quick Start</h2></p><ol><li>1) Go to <b>Authentication > /oauth/token</b></li><li>2) Enter <b>app_key, app_secret, username, password</b> fields and then click \"Try it out!\"</li><li>3) Upon success, your access_token is loaded and you can access any form requiring authorization.</li></ol><h2>Links</h2><ul><li><a href=\"https://github.com/ringcentral\" target=\"_blank\">RingCentral SDKs on Github</a></li><li><a href=\"mailto:devsupport@ringcentral.com\">RingCentral Developer Support Email</a></li></ul>
 *
 * API version: 1.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech) DO NOT EDIT.

package ringcentral

type BulkProvisionExtensionResource struct {
	Uri              string                       `json:"uri,omitempty"`
	Id               string                       `json:"id,omitempty"`
	PartnerId        string                       `json:"partnerId,omitempty"`
	ExtensionNumber  string                       `json:"extensionNumber,omitempty"`
	LoginName        string                       `json:"loginName,omitempty"`
	Contact          ContactInfo                  `json:"contact,omitempty"`
	References       []Reference                  `json:"references,omitempty"`
	Name             string                       `json:"name,omitempty"`
	Type             string                       `json:"type,omitempty"`
	Status           string                       `json:"status,omitempty"`
	StatusInfo       StatusInfo                   `json:"statusInfo,omitempty"`
	Departments      []DepartmentResource         `json:"departments,omitempty"`
	ServiceFeatures  []ServiceFeatureValue        `json:"serviceFeatures,omitempty"`
	RegionalSettings RegionalSettingsResource     `json:"regionalSettings,omitempty"`
	SetupWizardState string                       `json:"setupWizardState,omitempty"`
	Permissions      ExtensionPermissionsResource `json:"permissions,omitempty"`
	Password         string                       `json:"password,omitempty"`
	IvrPin           string                       `json:"ivrPin,omitempty"`
	ProfileImage     ProfileImageResource         `json:"profileImage,omitempty"`
}
