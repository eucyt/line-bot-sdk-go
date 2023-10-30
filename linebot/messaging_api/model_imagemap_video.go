/**
 * LINE Messaging API
 * This document describes LINE Messaging API.
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
package messaging_api

// ImagemapVideo
// ImagemapVideo

type ImagemapVideo struct {

	/**
	 * Get OriginalContentUrl
	 */
	OriginalContentUrl string `json:"originalContentUrl,omitempty"`

	/**
	 * Get PreviewImageUrl
	 */
	PreviewImageUrl string `json:"previewImageUrl,omitempty"`

	/**
	 * Get Area
	 */
	Area *ImagemapArea `json:"area,omitempty"`

	/**
	 * Get ExternalLink
	 */
	ExternalLink *ImagemapExternalLink `json:"externalLink,omitempty"`
}