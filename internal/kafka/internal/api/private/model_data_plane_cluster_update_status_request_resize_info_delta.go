/*
 * Kafka Service Fleet Manager
 *
 * Kafka Service Fleet Manager APIs that are used by internal services e.g kas-fleetshard operators.
 *
 * API version: 1.2.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package private

// DataPlaneClusterUpdateStatusRequestResizeInfoDelta struct for DataPlaneClusterUpdateStatusRequestResizeInfoDelta
type DataPlaneClusterUpdateStatusRequestResizeInfoDelta struct {
	IngressEgressThroughputPerSec *string `json:"ingress_egress_throughput_per_sec,omitempty"`
	Connections                   *int32  `json:"connections,omitempty"`
	DataRetentionSize             *string `json:"data_retention_size,omitempty"`
	Partitions                    *int32  `json:"partitions,omitempty"`
}
