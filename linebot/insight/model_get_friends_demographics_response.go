/**
 * LINE Messaging API(Insight)
 * This document describes LINE Messaging API(Insight).
 *
 * The version of the OpenAPI document: 0.0.1
 *
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

/**
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

//go:generate python3 ../../generate-code.py
package insight

// GetFriendsDemographicsResponse
// Get friend demographics
// https://developers.line.biz/en/reference/messaging-api/#get-demographic
type GetFriendsDemographicsResponse struct {

	/**
	 * true if friend demographic information is available.
	 */
	Available bool `json:"available"`

	/**
	 * Percentage per gender.
	 */
	Genders []GenderTile `json:"genders,omitempty"`

	/**
	 * Percentage per age group.
	 */
	Ages []AgeTile `json:"ages,omitempty"`

	/**
	 * Percentage per area.
	 */
	Areas []AreaTile `json:"areas,omitempty"`

	/**
	 * Percentage by OS.
	 */
	AppTypes []AppTypeTile `json:"appTypes,omitempty"`

	/**
	 * Percentage per friendship duration.
	 */
	SubscriptionPeriods []SubscriptionPeriodTile `json:"subscriptionPeriods,omitempty"`
}