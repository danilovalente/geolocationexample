/*
 * Geolocation API
 *
 * This is a Geolocation API
 *
 * API version: 1.0.0
 * Contact: valente.danilo@gmail.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package domain

//Position contains the combined Latitude and Longitude
type Position struct {
	Lat float32 `bson:"lat" json:"lat"`

	Lng float32 `bson:"lng" json:"lng"`
}
