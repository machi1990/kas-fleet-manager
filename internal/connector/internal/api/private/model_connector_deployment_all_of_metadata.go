/*
 * Connector Service Fleet Manager Private APIs
 *
 * Connector Service Fleet Manager apis that are used by internal services.
 *
 * API version: 0.0.3
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package private

import (
	"time"
)

// ConnectorDeploymentAllOfMetadata struct for ConnectorDeploymentAllOfMetadata
type ConnectorDeploymentAllOfMetadata struct {
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
	ResourceVersion int64     `json:"resource_version"`
	ResolvedSecrets bool      `json:"resolved_secrets"`
}
