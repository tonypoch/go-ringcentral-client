/*
 * RingCentral Connect Platform API Explorer
 *
 * <p>This is a beta interactive API explorer for the RingCentral Connect Platform. To use this service, you will need to have an account with the proper credentials to generate an OAuth2 access token.</p><p><h2>Quick Start</h2></p><ol><li>1) Go to <b>Authentication > /oauth/token</b></li><li>2) Enter <b>app_key, app_secret, username, password</b> fields and then click \"Try it out!\"</li><li>3) Upon success, your access_token is loaded and you can access any form requiring authorization.</li></ol><h2>Links</h2><ul><li><a href=\"https://github.com/ringcentral\" target=\"_blank\">RingCentral SDKs on Github</a></li><li><a href=\"mailto:devsupport@ringcentral.com\">RingCentral Developer Support Email</a></li></ul>
 *
 * API version: 1.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech) DO NOT EDIT.

package ringcentral

type PagingInfo struct {
	// The current page number. 1-indexed, so the first page is 1 by default. May be omitted if result is empty (because non-existent page was specified or perPage=0 was requested)
	Page int32 `json:"page,omitempty"`
	// Current page size, describes how many items are in each page. Default value is 100. Maximum value is 1000. If perPage value in the request is greater than 1000, the maximum value (1000) is applied
	PerPage int32 `json:"perPage,omitempty"`
	// The zero-based number of the first element on the current page. Omitted if the page is omitted or result is empty
	PageStart int32 `json:"pageStart,omitempty"`
	// The zero-based index of the last element on the current page. Omitted if the page is omitted or result is empty
	PageEnd int32 `json:"pageEnd,omitempty"`
	// The total number of pages in a dataset. May be omitted for some resources due to performance reasons
	TotalPages int32 `json:"totalPages,omitempty"`
	// The total number of elements in a dataset. May be omitted for some resource due to performance reasons
	TotalElements int32 `json:"totalElements,omitempty"`
}
