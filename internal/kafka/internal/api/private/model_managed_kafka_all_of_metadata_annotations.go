/*
 * Kafka Service Fleet Manager
 *
 * Kafka Service Fleet Manager APIs that are used by internal services e.g kas-fleetshard operators.
 *
 * API version: 1.2.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package private

// ManagedKafkaAllOfMetadataAnnotations struct for ManagedKafkaAllOfMetadataAnnotations
type ManagedKafkaAllOfMetadataAnnotations struct {
	// Deprecated
	DeprecatedBf2OrgId string `json:"bf2.org/id"`
	// Deprecated
	DeprecatedBf2OrgPlacementId string `json:"bf2.org/placementId"`
	Id                          string `json:"id"`
	PlacementId                 string `json:"placement_id"`
}