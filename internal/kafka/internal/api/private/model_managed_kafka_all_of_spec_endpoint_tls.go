/*
 * Kafka Service Fleet Manager
 *
 * Kafka Service Fleet Manager APIs that are used by internal services e.g kas-fleetshard operators.
 *
 * API version: 1.8.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package private

// ManagedKafkaAllOfSpecEndpointTls struct for ManagedKafkaAllOfSpecEndpointTls
type ManagedKafkaAllOfSpecEndpointTls struct {
	Cert string `json:"cert,omitempty"`
	Key  string `json:"key,omitempty"`
}
