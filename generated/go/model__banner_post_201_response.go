// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

/*
 * Сервис баннеров
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.0
 */

package openapi




type BannerPost201Response struct {

	// Идентификатор созданного баннера
	BannerId int32 `json:"banner_id,omitempty"`
}

// AssertBannerPost201ResponseRequired checks if the required fields are not zero-ed
func AssertBannerPost201ResponseRequired(obj BannerPost201Response) error {
	return nil
}

// AssertBannerPost201ResponseConstraints checks if the values respects the defined constraints
func AssertBannerPost201ResponseConstraints(obj BannerPost201Response) error {
	return nil
}