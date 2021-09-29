/*
 * Kafka Service Fleet Manager
 *
 * Kafka Service Fleet Manager is a Rest API to manage Kafka instances.
 *
 * API version: 1.1.3
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package public

// MetricsInstantQueryList struct for MetricsInstantQueryList
type MetricsInstantQueryList struct {
	Kind  string         `json:"kind,omitempty"`
	Id    string         `json:"id,omitempty"`
	Items []InstantQuery `json:"items,omitempty"`
}
