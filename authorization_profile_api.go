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

import (
	"net/url"
	"strings"
	"encoding/json"
	"fmt"
)

type AuthorizationProfileApi struct {
	Configuration *Configuration
}

func NewAuthorizationProfileApi() *AuthorizationProfileApi {
	configuration := NewConfiguration()
	return &AuthorizationProfileApi{
		Configuration: configuration,
	}
}

func NewAuthorizationProfileApiWithBasePath(basePath string) *AuthorizationProfileApi {
	configuration := NewConfiguration()
	configuration.BasePath = basePath

	return &AuthorizationProfileApi{
		Configuration: configuration,
	}
}

/**
 * 
 * Check User Permissions
 *
 * @param accountId Internal identifier of a RingCentral account or tilde (~) to indicate the account logged-in within the current session
 * @param extensionId Internal identifier of an extension or tilde (~) to indicate the extension assigned to the account logged-in within the current session
 * @param permissionId Permission to check
 * @param targetExtensionId Optional. Internal identifier of an extension for which user permissions are to be checked. The default value is the currently logged-in extension
 * @return *InlineResponseDefault13
 */
func (a AuthorizationProfileApi) RestapiV10AccountAccountIdExtensionExtensionIdAuthzProfileCheckGet(accountId string, extensionId string, permissionId string, targetExtensionId string) (*InlineResponseDefault13, *APIResponse, error) {

	var localVarHttpMethod = strings.ToUpper("Get")
	// create path and map variables
	localVarPath := a.Configuration.BasePath + "/restapi/v1.0/account/{accountId}/extension/{extensionId}/authz-profile/check"
	localVarPath = strings.Replace(localVarPath, "{"+"accountId"+"}", fmt.Sprintf("%v", accountId), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"extensionId"+"}", fmt.Sprintf("%v", extensionId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make(map[string]string)
	var localVarPostBody interface{}
	var localVarFileName string
	var localVarFileBytes []byte
	// authentication '(oauth)' required
	// oauth required
	if a.Configuration.AccessToken != ""{
		localVarHeaderParams["Authorization"] =  "Bearer " + a.Configuration.AccessToken
	}
	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		localVarHeaderParams[key] = a.Configuration.DefaultHeader[key]
	}
	localVarQueryParams.Add("permissionId", a.Configuration.APIClient.ParameterToString(permissionId, ""))
	localVarQueryParams.Add("targetExtensionId", a.Configuration.APIClient.ParameterToString(targetExtensionId, ""))

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	var successPayload = new(InlineResponseDefault13)
	localVarHttpResponse, err := a.Configuration.APIClient.CallAPI(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)

	var localVarURL, _ = url.Parse(localVarPath)
	localVarURL.RawQuery = localVarQueryParams.Encode()
	var localVarAPIResponse = &APIResponse{Operation: "RestapiV10AccountAccountIdExtensionExtensionIdAuthzProfileCheckGet", Method: localVarHttpMethod, RequestURL: localVarURL.String()}
	if localVarHttpResponse != nil {
		localVarAPIResponse.Response = localVarHttpResponse.RawResponse
		localVarAPIResponse.Payload = localVarHttpResponse.Body()
	}

	if err != nil {
		return successPayload, localVarAPIResponse, err
	}
	err = json.Unmarshal(localVarHttpResponse.Body(), &successPayload)
	return successPayload, localVarAPIResponse, err
}

/**
 * 
 * Get User Permissions
 *
 * @param accountId Internal identifier of a RingCentral account or tilde (~) to indicate the account logged-in within the current session
 * @param extensionId Internal identifier of an extension or tilde (~) to indicate the extension assigned to the account logged-in within the current session
 * @return *InlineResponseDefault12
 */
func (a AuthorizationProfileApi) RestapiV10AccountAccountIdExtensionExtensionIdAuthzProfileGet(accountId string, extensionId string) (*InlineResponseDefault12, *APIResponse, error) {

	var localVarHttpMethod = strings.ToUpper("Get")
	// create path and map variables
	localVarPath := a.Configuration.BasePath + "/restapi/v1.0/account/{accountId}/extension/{extensionId}/authz-profile"
	localVarPath = strings.Replace(localVarPath, "{"+"accountId"+"}", fmt.Sprintf("%v", accountId), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"extensionId"+"}", fmt.Sprintf("%v", extensionId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make(map[string]string)
	var localVarPostBody interface{}
	var localVarFileName string
	var localVarFileBytes []byte
	// authentication '(oauth)' required
	// oauth required
	if a.Configuration.AccessToken != ""{
		localVarHeaderParams["Authorization"] =  "Bearer " + a.Configuration.AccessToken
	}
	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		localVarHeaderParams[key] = a.Configuration.DefaultHeader[key]
	}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	var successPayload = new(InlineResponseDefault12)
	localVarHttpResponse, err := a.Configuration.APIClient.CallAPI(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)

	var localVarURL, _ = url.Parse(localVarPath)
	localVarURL.RawQuery = localVarQueryParams.Encode()
	var localVarAPIResponse = &APIResponse{Operation: "RestapiV10AccountAccountIdExtensionExtensionIdAuthzProfileGet", Method: localVarHttpMethod, RequestURL: localVarURL.String()}
	if localVarHttpResponse != nil {
		localVarAPIResponse.Response = localVarHttpResponse.RawResponse
		localVarAPIResponse.Payload = localVarHttpResponse.Body()
	}

	if err != nil {
		return successPayload, localVarAPIResponse, err
	}
	err = json.Unmarshal(localVarHttpResponse.Body(), &successPayload)
	return successPayload, localVarAPIResponse, err
}

