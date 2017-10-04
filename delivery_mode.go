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

type DeliveryMode struct {

	// Notifications transportation provider name. 'APNS' (Apple Push Notifications Service)
	TransportType string `json:"transportType,omitempty"`

	// Optional parameter. Specifies if the message will be encrypted or not. For APNS transport type the value is always \"false\"
	Encryption bool `json:"encryption,omitempty"`

	// PubNub channel name. For APNS transport type - internal identifier of a device \"device_token\"
	Address string `json:"address,omitempty"`

	// PubNub subscriber credentials required to subscribe to the channel
	SubscriberKey string `json:"subscriberKey,omitempty"`

	// PubNub subscriber credentials required to subscribe to the channel. Optional (for PubNub transport type only)
	SecretKey string `json:"secretKey,omitempty"`

	// Encryption algorithm 'AES' (for PubNub transport type only)
	EncryptionAlgorithm string `json:"encryptionAlgorithm,omitempty"`

	// Key for notification message decryption (for PubNub transport type only)
	EncryptionKey string `json:"encryptionKey,omitempty"`
}
