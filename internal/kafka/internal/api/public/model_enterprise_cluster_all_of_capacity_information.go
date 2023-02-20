/*
 * Kafka Management API
 *
 * Kafka Management API is a REST API to manage Kafka instances
 *
 * API version: 1.15.0
 * Contact: rhosak-support@redhat.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package public

// EnterpriseClusterAllOfCapacityInformation Returns the capacity related information
type EnterpriseClusterAllOfCapacityInformation struct {
	// The kafka machine pool node count provided during cluster registration
	KafkaMachinePoolNodeCount int32 `json:"kafka_machine_pool_node_count"`
	// The maximum number of Kafka streaming units that can be created on this cluster
	MaximumKafkaStreamingUnits int32 `json:"maximum_kafka_streaming_units"`
	// The remaining number of Kafka streaming units that can be still be created on this cluster
	RemainingKafkaStreamingUnits int32 `json:"remaining_kafka_streaming_units"`
	// The number of Kafka streaming units that have been consumed on this cluster
	ConsumedKafkaStreamingUnits int32 `json:"consumed_kafka_streaming_units"`
}