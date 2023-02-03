package services

import (
	"context"
	"fmt"
	"reflect"
	"strings"
	"testing"

	"github.com/bf2fc6cc711aee1a0c2a/kas-fleet-manager/internal/kafka/constants"
	"github.com/bf2fc6cc711aee1a0c2a/kas-fleet-manager/internal/kafka/internal/api/dbapi"
	"github.com/bf2fc6cc711aee1a0c2a/kas-fleet-manager/internal/kafka/internal/config"
	"github.com/bf2fc6cc711aee1a0c2a/kas-fleet-manager/pkg/api"
	"github.com/bf2fc6cc711aee1a0c2a/kas-fleet-manager/pkg/errors"
	"github.com/bf2fc6cc711aee1a0c2a/kas-fleet-manager/pkg/logger"
	"github.com/bf2fc6cc711aee1a0c2a/kas-fleet-manager/pkg/shared/utils/arrays"
	"github.com/onsi/gomega"
)

func Test_dataPlaneKafkaService_UpdateDataPlaneKafkaService(t *testing.T) {
	nonSecretKafkaStatus := "test failed message"
	secretError := "'secret': leaked secret"
	testErrorCondMessage := fmt.Sprintf("test failed message including '%s", secretError)
	bootstrapServer := "test.kafka.example.com"
	ingress := fmt.Sprintf("elb.%s", bootstrapServer)
	type fields struct {
		clusterService ClusterService
		kafkaService   func(c map[string]int) KafkaService
	}
	type args struct {
		clusterId string
		status    []*dbapi.DataPlaneKafkaStatus
	}
	tests := []struct {
		name           string
		fields         fields
		args           args
		expectCounters map[string]int
		want           *errors.ServiceError
	}{
		{
			name: "should return error when cluster id is not valid",
			fields: fields{
				clusterService: &ClusterServiceMock{
					FindClusterByIDFunc: func(clusterID string) (*api.Cluster, *errors.ServiceError) {
						return nil, nil
					},
				},
				kafkaService: func(c map[string]int) KafkaService {
					return &KafkaServiceMock{}
				},
			},
			args: args{
				clusterId: "test-cluster-id",
				status:    []*dbapi.DataPlaneKafkaStatus{},
			},
			want: errors.BadRequest("clusterID \"test-cluster-id\" not found"),
			expectCounters: map[string]int{
				"ready":     0,
				"failed":    0,
				"deleting":  0,
				"rejected":  0,
				"suspended": 0,
			},
		},
		{
			name: "should return no error and update dataplane kafkas with various conditions",
			fields: fields{
				clusterService: &ClusterServiceMock{
					FindClusterByIDFunc: func(clusterID string) (*api.Cluster, *errors.ServiceError) {
						return &api.Cluster{ClusterID: "test-cluster-id"}, nil
					},
				},
				kafkaService: func(c map[string]int) KafkaService {
					return &KafkaServiceMock{
						GetByIDFunc: func(id string) (*dbapi.KafkaRequest, *errors.ServiceError) {
							return &dbapi.KafkaRequest{
								ClusterID:     "test-cluster-id",
								Status:        constants.KafkaRequestStatusProvisioning.String(),
								Routes:        []byte("[{'domain':'test.example.com', 'router':'test.example.com'}]"),
								RoutesCreated: true,
							}, nil
						},
						UpdateFunc: func(kafkaRequest *dbapi.KafkaRequest) *errors.ServiceError {
							if kafkaRequest.Status == string(constants.KafkaRequestStatusFailed) {
								if strings.Contains(kafkaRequest.FailedReason, secretError) {
									return errors.GeneralError("test failure error. Expected FailedReason is empty")
								}
								c["failed"]++
							} else if kafkaRequest.Status == string(constants.KafkaRequestStatusReady) {
								c["ready"]++
							} else if kafkaRequest.Status == string(constants.KafkaRequestStatusDeleting) {
								c["deleting"]++
							} else {
								c["rejected"]++
							}
							return nil
						},
						UpdatesFunc: func(kafkaRequest *dbapi.KafkaRequest, values map[string]interface{}) *errors.ServiceError {
							v, ok := values["status"]
							if ok {
								statusValue := v.(string)
								c[statusValue]++
							}
							return nil
						},
						UpdateStatusFunc: func(id string, status constants.KafkaStatus) (bool, *errors.ServiceError) {
							if status == constants.KafkaRequestStatusReady {
								c["ready"]++
							} else if status == constants.KafkaRequestStatusDeleting {
								c["deleting"]++
							} else if status == constants.KafkaRequestStatusFailed {
								c["failed"]++
							}
							return true, nil
						},
						DeleteFunc: func(in1 *dbapi.KafkaRequest) *errors.ServiceError {
							return nil
						},
					}
				},
			},
			args: args{
				clusterId: "test-cluster-id",
				status: []*dbapi.DataPlaneKafkaStatus{
					{
						Conditions: []dbapi.DataPlaneKafkaStatusCondition{
							{
								Type:   "Ready",
								Status: "True",
							},
						},
					},
					{
						Conditions: []dbapi.DataPlaneKafkaStatusCondition{
							{
								Type:   "Ready",
								Status: "False",
								Reason: "Installing",
							},
						},
					},
					{
						Conditions: []dbapi.DataPlaneKafkaStatusCondition{
							{
								Type:    "Ready",
								Status:  "False",
								Reason:  "Error",
								Message: testErrorCondMessage,
							},
						},
					},
					{
						Conditions: []dbapi.DataPlaneKafkaStatusCondition{
							{
								Type:   "Ready",
								Status: "False",
								Reason: "Deleted",
							},
						},
					},
					{
						Conditions: []dbapi.DataPlaneKafkaStatusCondition{
							{
								Type:   "Ready",
								Status: "False",
								Reason: "Rejected",
							},
						},
					},
				},
			},
			want: nil,
			expectCounters: map[string]int{
				"ready":     1,
				"failed":    1,
				"deleting":  1,
				"rejected":  1,
				"suspended": 0,
			},
		},
		{
			name: "should use routes in the requests if they are present",
			fields: fields{
				clusterService: &ClusterServiceMock{
					FindClusterByIDFunc: func(clusterID string) (*api.Cluster, *errors.ServiceError) {
						return &api.Cluster{ClusterID: "test-cluster-id"}, nil
					},
					GetClusterDNSFunc: func(clusterID string) (string, *errors.ServiceError) {
						return bootstrapServer, nil
					},
				},
				kafkaService: func(c map[string]int) KafkaService {
					routesCreated := false
					expectedRoutes := []dbapi.DataPlaneKafkaRoute{
						{
							Domain: bootstrapServer,
							Router: ingress,
						},
						{
							Domain: fmt.Sprintf("admin-api-%s", bootstrapServer),
							Router: ingress,
						},
						{
							Domain: fmt.Sprintf("broker-0-%s", bootstrapServer),
							Router: ingress,
						},
					}
					return &KafkaServiceMock{
						GetByIDFunc: func(id string) (*dbapi.KafkaRequest, *errors.ServiceError) {
							return &dbapi.KafkaRequest{
								ClusterID:           "test-cluster-id",
								Status:              constants.KafkaRequestStatusProvisioning.String(),
								BootstrapServerHost: bootstrapServer,
								RoutesCreated:       routesCreated,
							}, nil
						},
						UpdateFunc: func(kafkaRequest *dbapi.KafkaRequest) *errors.ServiceError {
							routes, err := kafkaRequest.GetRoutes()
							if err != nil || !reflect.DeepEqual(routes, expectedRoutes) {
								c["rejected"]++
							} else {
								routesCreated = true
							}
							if kafkaRequest.Status == string(constants.KafkaRequestStatusReady) {
								c["ready"]++
							} else if kafkaRequest.Status == string(constants.KafkaRequestStatusDeleting) {
								c["deleting"]++
							} else if kafkaRequest.Status == string(constants.KafkaRequestStatusFailed) {
								c["failed"]++
							}
							return nil
						},
						UpdatesFunc: func(kafkaRequest *dbapi.KafkaRequest, values map[string]interface{}) *errors.ServiceError {
							v, ok := values["status"]
							if ok {
								statusValue := v.(string)
								c[statusValue]++
							}
							return nil
						},
						UpdateStatusFunc: func(id string, status constants.KafkaStatus) (bool, *errors.ServiceError) {
							if status == constants.KafkaRequestStatusReady {
								c["ready"]++
							} else if status == constants.KafkaRequestStatusDeleting {
								c["deleting"]++
							} else if status == constants.KafkaRequestStatusFailed {
								c["failed"]++
							}
							return true, nil
						},
						DeleteFunc: func(in1 *dbapi.KafkaRequest) *errors.ServiceError {
							return nil
						},
					}
				},
			},
			args: args{
				clusterId: "test-cluster-id",
				status: []*dbapi.DataPlaneKafkaStatus{
					// route not available yet, so Kafka will not update (rejected count +1)
					{
						Conditions: []dbapi.DataPlaneKafkaStatusCondition{
							{
								Type:   "Ready",
								Status: "False",
								Reason: "Installing",
							},
						},
					},
					// routes available, this will set "RoutesCreated" to true but should not set status to Ready
					{
						Conditions: []dbapi.DataPlaneKafkaStatusCondition{
							{
								Type:   "Ready",
								Status: "False",
								Reason: "Installing",
							},
						},
						Routes: []dbapi.DataPlaneKafkaRouteRequest{
							{
								Name:   "bootstrap",
								Prefix: "",
								Router: ingress,
							},
							{
								Name:   "admin-api",
								Prefix: "admin-api",
								Router: ingress,
							},
							{
								Name:   "broker-0",
								Prefix: "broker-0",
								Router: ingress,
							},
						},
					},
					// This will then set the kafka instance to be ready
					{
						Conditions: []dbapi.DataPlaneKafkaStatusCondition{
							{
								Type:   "Ready",
								Status: "True",
							},
						},
						Routes: []dbapi.DataPlaneKafkaRouteRequest{
							{
								Name:   "bootstrap",
								Prefix: "",
								Router: ingress,
							},
							{
								Name:   "admin-api",
								Prefix: "admin-api",
								Router: ingress,
							},
							{
								Name:   "broker-0",
								Prefix: "broker-0",
								Router: ingress,
							},
						},
					},
				},
			},
			want: nil,
			expectCounters: map[string]int{
				"ready":     1,
				"failed":    0,
				"deleting":  0,
				"rejected":  0,
				"suspended": 0,
			},
		},
		{
			name: "success when updates kafka status to ready and removes failed reason",
			fields: fields{
				clusterService: &ClusterServiceMock{
					FindClusterByIDFunc: func(clusterID string) (*api.Cluster, *errors.ServiceError) {
						return &api.Cluster{ClusterID: "test-cluster-id"}, nil
					},
				},
				kafkaService: func(c map[string]int) KafkaService {
					return &KafkaServiceMock{
						GetByIDFunc: func(id string) (*dbapi.KafkaRequest, *errors.ServiceError) {
							return &dbapi.KafkaRequest{
								ClusterID:     "test-cluster-id",
								Status:        constants.KafkaRequestStatusProvisioning.String(),
								Routes:        []byte("[{'domain':'test.example.com', 'router':'test.example.com'}]"),
								RoutesCreated: true,
								FailedReason:  nonSecretKafkaStatus,
							}, nil
						},
						UpdateFunc: func(kafkaRequest *dbapi.KafkaRequest) *errors.ServiceError {
							if kafkaRequest.Status == string(constants.KafkaRequestStatusFailed) {
								if !strings.Contains(kafkaRequest.FailedReason, nonSecretKafkaStatus) {
									return errors.GeneralError("test failure error. Expected FailedReason is empty")
								}
								c["failed"]++
							} else if kafkaRequest.Status == string(constants.KafkaRequestStatusReady) {
								c["ready"]++
							} else if kafkaRequest.Status == string(constants.KafkaRequestStatusDeleting) {
								c["deleting"]++
							} else {
								c["rejected"]++
							}
							return nil
						},
						UpdatesFunc: func(kafkaRequest *dbapi.KafkaRequest, values map[string]interface{}) *errors.ServiceError {
							v, ok := values["status"]
							if ok {
								statusValue := v.(string)
								c[statusValue]++
							}
							return nil
						},
						UpdateStatusFunc: func(id string, status constants.KafkaStatus) (bool, *errors.ServiceError) {
							if status == constants.KafkaRequestStatusReady {
								c["ready"]++
							} else if status == constants.KafkaRequestStatusDeleting {
								c["deleting"]++
							} else if status == constants.KafkaRequestStatusFailed {
								c["failed"]++
							}
							return true, nil
						},
						DeleteFunc: func(in1 *dbapi.KafkaRequest) *errors.ServiceError {
							return nil
						},
					}
				},
			},
			args: args{
				clusterId: "test-cluster-id",
				status: []*dbapi.DataPlaneKafkaStatus{
					{
						Conditions: []dbapi.DataPlaneKafkaStatusCondition{
							{
								Type:   "Ready",
								Status: "True",
							},
						},
					},
				},
			},
			want: nil,
			expectCounters: map[string]int{
				"ready":     1,
				"failed":    0,
				"deleting":  0,
				"rejected":  0,
				"suspended": 0,
			},
		},
		// Kafka suspension test cases
		{
			name: "should only update a suspending Kafka instance to suspended",
			fields: fields{
				clusterService: &ClusterServiceMock{
					FindClusterByIDFunc: func(clusterID string) (*api.Cluster, *errors.ServiceError) {
						return &api.Cluster{ClusterID: "test-cluster-id"}, nil
					},
				},
				kafkaService: func(c map[string]int) KafkaService {
					return &KafkaServiceMock{
						GetByIDFunc: func(id string) (*dbapi.KafkaRequest, *errors.ServiceError) {
							return &dbapi.KafkaRequest{
								ClusterID:     "test-cluster-id",
								Status:        constants.KafkaRequestStatusSuspending.String(),
								Routes:        []byte("[{'domain':'test.example.com', 'router':'test.example.com'}]"),
								RoutesCreated: true,
							}, nil
						},
						UpdateStatusFunc: func(id string, status constants.KafkaStatus) (bool, *errors.ServiceError) {
							if status == constants.KafkaRequestStatusSuspended {
								c["suspended"]++
							}
							return true, nil
						},
					}
				},
			},
			args: args{
				clusterId: "test-cluster-id",
				status: []*dbapi.DataPlaneKafkaStatus{
					{
						Conditions: []dbapi.DataPlaneKafkaStatusCondition{
							{
								Type:   "Ready",
								Status: "True",
							},
						},
					},
					{
						Conditions: []dbapi.DataPlaneKafkaStatusCondition{
							{
								Type:   "Ready",
								Reason: "Installing",
								Status: "False",
							},
						},
					},
					{
						Conditions: []dbapi.DataPlaneKafkaStatusCondition{
							{
								Type:    "Ready",
								Reason:  "Error",
								Status:  "False",
								Message: "kafka reported as failed by data plane",
							},
						},
					},
					{
						Conditions: []dbapi.DataPlaneKafkaStatusCondition{
							{
								Type:    "Ready",
								Reason:  "Rejected",
								Status:  "False",
								Message: "kafka reported as rejected by data plane",
							},
						},
					},
					{
						Conditions: []dbapi.DataPlaneKafkaStatusCondition{
							{
								Type:    "Ready",
								Reason:  "Rejected",
								Status:  "False",
								Message: "Cluster has insufficient resources",
							},
						},
					},
					{
						Conditions: []dbapi.DataPlaneKafkaStatusCondition{
							{
								Type:   "Ready",
								Reason: "Suspended",
								Status: "False",
							},
						},
					},
				},
			},
			want: nil,
			expectCounters: map[string]int{
				"ready":     0,
				"deleting":  0,
				"failed":    0,
				"rejected":  0,
				"suspended": 1,
			},
		},
		{
			name: "should only update a resuming Kafka instance to ready or failed",
			fields: fields{
				clusterService: &ClusterServiceMock{
					FindClusterByIDFunc: func(clusterID string) (*api.Cluster, *errors.ServiceError) {
						return &api.Cluster{ClusterID: "test-cluster-id"}, nil
					},
				},
				kafkaService: func(c map[string]int) KafkaService {
					return &KafkaServiceMock{
						GetByIDFunc: func(id string) (*dbapi.KafkaRequest, *errors.ServiceError) {
							return &dbapi.KafkaRequest{
								ClusterID:     "test-cluster-id",
								Status:        constants.KafkaRequestStatusResuming.String(),
								Routes:        []byte("[{'domain':'test.example.com', 'router':'test.example.com'}]"),
								RoutesCreated: true,
							}, nil
						},
						UpdatesFunc: func(kafkaRequest *dbapi.KafkaRequest, values map[string]interface{}) *errors.ServiceError {
							v, ok := values["status"]
							if ok {
								statusValue := v.(string)
								c[statusValue]++
							}
							return nil
						},
						UpdateFunc: func(kafkaRequest *dbapi.KafkaRequest) *errors.ServiceError {
							if kafkaRequest.Status == string(constants.KafkaRequestStatusFailed) {
								if arrays.StringEmptyPredicate(kafkaRequest.FailedReason) {
									return errors.GeneralError("Test failure error. FailedReason should not be empty")
								}
								c["failed"]++
							}
							return nil
						},
					}
				},
			},
			args: args{
				clusterId: "test-cluster-id",
				status: []*dbapi.DataPlaneKafkaStatus{
					{
						Conditions: []dbapi.DataPlaneKafkaStatusCondition{
							{
								Type:   "Ready",
								Status: "True",
							},
						},
					},
					{
						Conditions: []dbapi.DataPlaneKafkaStatusCondition{
							{
								Type:   "Ready",
								Reason: "Installing",
								Status: "False",
							},
						},
					},
					{
						Conditions: []dbapi.DataPlaneKafkaStatusCondition{
							{
								Type:    "Ready",
								Reason:  "Error",
								Status:  "False",
								Message: "kafka reported as failed by data plane",
							},
						},
					},
					{
						Conditions: []dbapi.DataPlaneKafkaStatusCondition{
							{
								Type:    "Ready",
								Reason:  "Rejected",
								Status:  "False",
								Message: "kafka reported as rejected by data plane",
							},
						},
					},
					{
						Conditions: []dbapi.DataPlaneKafkaStatusCondition{
							{
								Type:    "Ready",
								Reason:  "Rejected",
								Status:  "False",
								Message: "Cluster has insufficient resources",
							},
						},
					},
					{
						Conditions: []dbapi.DataPlaneKafkaStatusCondition{
							{
								Type:   "Ready",
								Reason: "Suspended",
								Status: "False",
							},
						},
					},
				},
			},
			want: nil,
			expectCounters: map[string]int{
				"ready":     1,
				"deleting":  0,
				"failed":    1,
				"rejected":  0,
				"suspended": 0,
			},
		},
		{
			name: "should never update a suspended Kafka instance",
			fields: fields{
				clusterService: &ClusterServiceMock{
					FindClusterByIDFunc: func(clusterID string) (*api.Cluster, *errors.ServiceError) {
						return &api.Cluster{ClusterID: "test-cluster-id"}, nil
					},
				},
				kafkaService: func(c map[string]int) KafkaService {
					return &KafkaServiceMock{
						GetByIDFunc: func(id string) (*dbapi.KafkaRequest, *errors.ServiceError) {
							return &dbapi.KafkaRequest{
								ClusterID:     "test-cluster-id",
								Status:        constants.KafkaRequestStatusSuspended.String(),
								Routes:        []byte("[{'domain':'test.example.com', 'router':'test.example.com'}]"),
								RoutesCreated: true,
							}, nil
						},
					}
				},
			},
			args: args{
				clusterId: "test-cluster-id",
				status: []*dbapi.DataPlaneKafkaStatus{
					{
						Conditions: []dbapi.DataPlaneKafkaStatusCondition{
							{
								Type:   "Ready",
								Status: "True",
							},
						},
					},
					{
						Conditions: []dbapi.DataPlaneKafkaStatusCondition{
							{
								Type:   "Ready",
								Reason: "Installing",
								Status: "False",
							},
						},
					},
					{
						Conditions: []dbapi.DataPlaneKafkaStatusCondition{
							{
								Type:    "Ready",
								Reason:  "Error",
								Status:  "False",
								Message: "kafka reported as failed by data plane",
							},
						},
					},
					{
						Conditions: []dbapi.DataPlaneKafkaStatusCondition{
							{
								Type:    "Ready",
								Reason:  "Rejected",
								Status:  "False",
								Message: "kafka reported as rejected by data plane",
							},
						},
					},
					{
						Conditions: []dbapi.DataPlaneKafkaStatusCondition{
							{
								Type:    "Ready",
								Reason:  "Rejected",
								Status:  "False",
								Message: "Cluster has insufficient resources",
							},
						},
					},
					{
						Conditions: []dbapi.DataPlaneKafkaStatusCondition{
							{
								Type:   "Ready",
								Reason: "Suspended",
								Status: "False",
							},
						},
					},
				},
			},
			want: nil,
			expectCounters: map[string]int{
				"ready":     0,
				"deleting":  0,
				"failed":    0,
				"rejected":  0,
				"suspended": 0,
			},
		},
	}

	for _, testcase := range tests {
		tt := testcase

		t.Run(tt.name, func(t *testing.T) {
			g := gomega.NewWithT(t)
			counter := map[string]int{
				"ready":     0,
				"failed":    0,
				"deleting":  0,
				"rejected":  0,
				"suspended": 0,
			}
			s := NewDataPlaneKafkaService(tt.fields.kafkaService(counter), tt.fields.clusterService, &config.KafkaConfig{})
			err := s.UpdateDataPlaneKafkaService(context.TODO(), tt.args.clusterId, tt.args.status)
			g.Expect(err).To(gomega.Equal(tt.want))
			g.Expect(counter).To(gomega.Equal(tt.expectCounters))
		})
	}
}

func TestDataPlaneKafkaService_UpdateVersions(t *testing.T) {
	type versions struct {
		actualKafkaVersion    string
		actualStrimziVersion  string
		actualKafkaIBPVersion string
		strimziUpgrading      bool
		kafkaUpgrading        bool
		kafkaIBPUpgrading     bool
	}

	tests := []struct {
		name             string
		clusterService   ClusterService
		kafkaService     func(v *versions) KafkaService
		clusterId        string
		status           []*dbapi.DataPlaneKafkaStatus
		wantErr          bool
		expectedVersions versions
	}{
		{
			name: "should update versions",
			clusterService: &ClusterServiceMock{
				FindClusterByIDFunc: func(clusterID string) (*api.Cluster, *errors.ServiceError) {
					return &api.Cluster{ClusterID: "test-cluster-id"}, nil
				},
			},
			kafkaService: func(v *versions) KafkaService {
				return &KafkaServiceMock{
					GetByIDFunc: func(id string) (*dbapi.KafkaRequest, *errors.ServiceError) {
						return &dbapi.KafkaRequest{
							ClusterID:             "test-cluster-id",
							Status:                constants.KafkaRequestStatusProvisioning.String(),
							Routes:                []byte("[{'domain':'test.example.com', 'router':'test.example.com'}]"),
							RoutesCreated:         true,
							ActualKafkaVersion:    "kafka-original-ver-0",
							ActualKafkaIBPVersion: "kafka-ibp-original-ver-0",
							ActualStrimziVersion:  "strimzi-original-ver-0",
						}, nil
					},
					UpdatesFunc: func(kafkaRequest *dbapi.KafkaRequest, fields map[string]interface{}) *errors.ServiceError {
						v.actualKafkaVersion = kafkaRequest.ActualKafkaVersion
						v.actualKafkaIBPVersion = kafkaRequest.ActualKafkaIBPVersion
						v.actualStrimziVersion = kafkaRequest.ActualStrimziVersion
						v.strimziUpgrading = kafkaRequest.StrimziUpgrading
						v.kafkaUpgrading = kafkaRequest.KafkaUpgrading
						v.kafkaIBPUpgrading = kafkaRequest.KafkaIBPUpgrading
						return nil
					},
					UpdateStatusFunc: func(id string, status constants.KafkaStatus) (bool, *errors.ServiceError) {
						return true, nil
					},
					DeleteFunc: func(in1 *dbapi.KafkaRequest) *errors.ServiceError {
						return nil
					},
				}
			},
			clusterId: "test-cluster-id",
			status: []*dbapi.DataPlaneKafkaStatus{
				{
					Conditions: []dbapi.DataPlaneKafkaStatusCondition{
						{
							Type:   "Ready",
							Status: "True",
							Reason: "StrimziUpdating",
						},
					},
					KafkaVersion:    "kafka-1",
					StrimziVersion:  "strimzi-1",
					KafkaIBPVersion: "kafka-ibp-3",
				},
			},
			wantErr: false,
			expectedVersions: versions{
				actualKafkaVersion:    "kafka-1",
				actualStrimziVersion:  "strimzi-1",
				actualKafkaIBPVersion: "kafka-ibp-3",
				strimziUpgrading:      true,
				kafkaUpgrading:        false,
				kafkaIBPUpgrading:     false,
			},
		},
		{
			name: "when the condition does not contain a reason then all upgrading fields should be set to false",
			clusterService: &ClusterServiceMock{
				FindClusterByIDFunc: func(clusterID string) (*api.Cluster, *errors.ServiceError) {
					return &api.Cluster{ClusterID: "test-cluster-id"}, nil
				},
			},
			kafkaService: func(v *versions) KafkaService {
				return &KafkaServiceMock{
					GetByIDFunc: func(id string) (*dbapi.KafkaRequest, *errors.ServiceError) {
						return &dbapi.KafkaRequest{
							ClusterID:             "test-cluster-id",
							Status:                constants.KafkaRequestStatusProvisioning.String(),
							Routes:                []byte("[{'domain':'test.example.com', 'router':'test.example.com'}]"),
							RoutesCreated:         true,
							ActualKafkaVersion:    "kafka-1",
							ActualStrimziVersion:  "strimzi-1",
							ActualKafkaIBPVersion: "kafka-ibp-1",
							StrimziUpgrading:      true,
							KafkaUpgrading:        true,
							KafkaIBPUpgrading:     true,
						}, nil
					},
					UpdatesFunc: func(kafkaRequest *dbapi.KafkaRequest, fields map[string]interface{}) *errors.ServiceError {
						v.actualKafkaVersion = kafkaRequest.ActualKafkaVersion
						v.actualKafkaIBPVersion = kafkaRequest.ActualKafkaIBPVersion
						v.actualStrimziVersion = kafkaRequest.ActualStrimziVersion
						v.strimziUpgrading = kafkaRequest.StrimziUpgrading
						v.kafkaUpgrading = kafkaRequest.KafkaUpgrading
						v.kafkaIBPUpgrading = kafkaRequest.KafkaIBPUpgrading
						return nil
					},
					UpdateStatusFunc: func(id string, status constants.KafkaStatus) (bool, *errors.ServiceError) {
						return true, nil
					},
					DeleteFunc: func(in1 *dbapi.KafkaRequest) *errors.ServiceError {
						return nil
					},
				}
			},
			clusterId: "test-cluster-id",
			status: []*dbapi.DataPlaneKafkaStatus{
				{
					Conditions: []dbapi.DataPlaneKafkaStatusCondition{
						{
							Type:   "Ready",
							Status: "True",
						},
					},
					KafkaVersion:    "kafka-1",
					KafkaIBPVersion: "kafka-ibp-1",
					StrimziVersion:  "strimzi-1",
				},
			},
			wantErr: true,
			expectedVersions: versions{
				actualKafkaVersion:    "kafka-1",
				actualKafkaIBPVersion: "kafka-ibp-1",
				actualStrimziVersion:  "strimzi-1",
				strimziUpgrading:      false,
				kafkaUpgrading:        false,
				kafkaIBPUpgrading:     false,
			},
		},
		{
			name: "when received condition is upgrading kafka then it is set to true if it wasn't",
			clusterService: &ClusterServiceMock{
				FindClusterByIDFunc: func(clusterID string) (*api.Cluster, *errors.ServiceError) {
					return &api.Cluster{ClusterID: "test-cluster-id"}, nil
				},
			},
			kafkaService: func(v *versions) KafkaService {
				return &KafkaServiceMock{
					GetByIDFunc: func(id string) (*dbapi.KafkaRequest, *errors.ServiceError) {
						return &dbapi.KafkaRequest{
							ClusterID:     "test-cluster-id",
							Status:        constants.KafkaRequestStatusProvisioning.String(),
							Routes:        []byte("[{'domain':'test.example.com', 'router':'test.example.com'}]"),
							RoutesCreated: true,
						}, nil
					},
					UpdatesFunc: func(kafkaRequest *dbapi.KafkaRequest, fields map[string]interface{}) *errors.ServiceError {
						v.actualKafkaVersion = kafkaRequest.ActualKafkaVersion
						v.actualKafkaIBPVersion = kafkaRequest.ActualKafkaIBPVersion
						v.actualStrimziVersion = kafkaRequest.ActualStrimziVersion
						v.strimziUpgrading = kafkaRequest.StrimziUpgrading
						v.kafkaUpgrading = kafkaRequest.KafkaUpgrading
						v.kafkaIBPUpgrading = kafkaRequest.KafkaIBPUpgrading
						return nil
					},
					UpdateStatusFunc: func(id string, status constants.KafkaStatus) (bool, *errors.ServiceError) {
						return true, nil
					},
					DeleteFunc: func(in1 *dbapi.KafkaRequest) *errors.ServiceError {
						return nil
					},
				}
			},
			clusterId: "test-cluster-id",
			status: []*dbapi.DataPlaneKafkaStatus{
				{
					Conditions: []dbapi.DataPlaneKafkaStatusCondition{
						{
							Type:   "Ready",
							Status: "True",
							Reason: "KafkaUpdating",
						},
					},
					KafkaVersion:    "kafka-1",
					StrimziVersion:  "strimzi-1",
					KafkaIBPVersion: "kafka-ibp-3",
				},
			},
			wantErr: false,
			expectedVersions: versions{
				actualKafkaVersion:    "kafka-1",
				actualStrimziVersion:  "strimzi-1",
				actualKafkaIBPVersion: "kafka-ibp-3",
				strimziUpgrading:      false,
				kafkaUpgrading:        true,
				kafkaIBPUpgrading:     false,
			},
		},
		{
			name: "when received condition is upgrading kafka ibp then it is set to true if it wasn't",
			clusterService: &ClusterServiceMock{
				FindClusterByIDFunc: func(clusterID string) (*api.Cluster, *errors.ServiceError) {
					return &api.Cluster{ClusterID: "test-cluster-id"}, nil
				},
			},
			kafkaService: func(v *versions) KafkaService {
				return &KafkaServiceMock{
					GetByIDFunc: func(id string) (*dbapi.KafkaRequest, *errors.ServiceError) {
						return &dbapi.KafkaRequest{
							ClusterID:     "test-cluster-id",
							Status:        constants.KafkaRequestStatusProvisioning.String(),
							Routes:        []byte("[{'domain':'test.example.com', 'router':'test.example.com'}]"),
							RoutesCreated: true,
						}, nil
					},
					UpdatesFunc: func(kafkaRequest *dbapi.KafkaRequest, fields map[string]interface{}) *errors.ServiceError {
						v.actualKafkaVersion = kafkaRequest.ActualKafkaVersion
						v.actualKafkaIBPVersion = kafkaRequest.ActualKafkaIBPVersion
						v.actualStrimziVersion = kafkaRequest.ActualStrimziVersion
						v.strimziUpgrading = kafkaRequest.StrimziUpgrading
						v.kafkaUpgrading = kafkaRequest.KafkaUpgrading
						v.kafkaIBPUpgrading = kafkaRequest.KafkaIBPUpgrading
						return nil
					},
					UpdateStatusFunc: func(id string, status constants.KafkaStatus) (bool, *errors.ServiceError) {
						return true, nil
					},
					DeleteFunc: func(in1 *dbapi.KafkaRequest) *errors.ServiceError {
						return nil
					},
				}
			},
			clusterId: "test-cluster-id",
			status: []*dbapi.DataPlaneKafkaStatus{
				{
					Conditions: []dbapi.DataPlaneKafkaStatusCondition{
						{
							Type:   "Ready",
							Status: "True",
							Reason: "KafkaIbpUpdating",
						},
					},
					KafkaVersion:    "kafka-1",
					StrimziVersion:  "strimzi-1",
					KafkaIBPVersion: "kafka-ibp-3",
				},
			},
			wantErr: false,
			expectedVersions: versions{
				actualKafkaVersion:    "kafka-1",
				actualStrimziVersion:  "strimzi-1",
				actualKafkaIBPVersion: "kafka-ibp-3",
				strimziUpgrading:      false,
				kafkaUpgrading:        false,
				kafkaIBPUpgrading:     true,
			},
		},
	}

	for _, testcase := range tests {
		tt := testcase

		t.Run(tt.name, func(t *testing.T) {
			g := gomega.NewWithT(t)
			v := versions{}
			s := NewDataPlaneKafkaService(tt.kafkaService(&v), tt.clusterService, &config.KafkaConfig{})
			err := s.UpdateDataPlaneKafkaService(context.TODO(), tt.clusterId, tt.status)
			if err != nil && !tt.wantErr {
				t.Errorf("unexpected error %v", err)
			}
			g.Expect(v).To(gomega.Equal(tt.expectedVersions))
		})
	}
}

func Test_DataPlaneKafkaStatus_getManagedKafkaStatus(t *testing.T) {
	type args struct {
		status *dbapi.DataPlaneKafkaStatus
	}
	tests := []struct {
		name string
		args args
		want managedKafkaStatus
	}{
		{
			name: "should return statusInstalling if status condition Type is not ready.",
			args: args{
				status: &dbapi.DataPlaneKafkaStatus{
					Conditions: []dbapi.DataPlaneKafkaStatusCondition{
						{
							Type: "Test",
						},
					},
				},
			},
			want: statusInstalling,
		},
		{
			name: "should return statusReady if kafka status is true.",
			args: args{
				status: &dbapi.DataPlaneKafkaStatus{
					Conditions: []dbapi.DataPlaneKafkaStatusCondition{
						{
							Type:   "Ready",
							Status: "True",
						},
					},
				},
			},
			want: statusReady,
		},
		{
			name: "should return statusUnknown if if kafka status is unknown.",
			args: args{
				status: &dbapi.DataPlaneKafkaStatus{
					Conditions: []dbapi.DataPlaneKafkaStatusCondition{
						{
							Type:   "Ready",
							Status: "Unknown",
						},
					},
				},
			},
			want: statusUnknown,
		},
		{
			name: "should return statusInstalling if kafka is Installing.",
			args: args{
				status: &dbapi.DataPlaneKafkaStatus{
					Conditions: []dbapi.DataPlaneKafkaStatusCondition{
						{
							Type:   "Ready",
							Status: "False",
							Reason: "Installing",
						},
					},
				},
			},
			want: statusInstalling,
		},
		{
			name: "should return statusDeleted if kafka is Deleted.",
			args: args{
				status: &dbapi.DataPlaneKafkaStatus{
					Conditions: []dbapi.DataPlaneKafkaStatusCondition{
						{
							Type:   "Ready",
							Status: "False",
							Reason: "Deleted",
						},
					},
				},
			},
			want: statusDeleted,
		},
		{
			name: "should return statusError if Error.",
			args: args{
				status: &dbapi.DataPlaneKafkaStatus{
					Conditions: []dbapi.DataPlaneKafkaStatusCondition{
						{
							Type:   "Ready",
							Status: "False",
							Reason: "Error",
						},
					},
				},
			},
			want: statusError,
		},
		{
			name: "should return statusRejectedClusterFull if cluster is full.",
			args: args{
				status: &dbapi.DataPlaneKafkaStatus{
					Conditions: []dbapi.DataPlaneKafkaStatusCondition{
						{
							Type:    "Ready",
							Status:  "False",
							Reason:  "Rejected",
							Message: "Cluster has insufficient resources",
						},
					},
				},
			},
			want: statusRejectedClusterFull,
		},
		{
			name: "should return statusRejected if Rejected.",
			args: args{
				status: &dbapi.DataPlaneKafkaStatus{
					Conditions: []dbapi.DataPlaneKafkaStatusCondition{
						{
							Type:   "Ready",
							Status: "False",
							Reason: "Rejected",
						},
					},
				},
			},
			want: statusRejected,
		},
	}

	for _, testcase := range tests {
		tt := testcase
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			g := gomega.NewWithT(t)
			d := &dataPlaneKafkaService{}
			got := d.getManagedKafkaStatus(tt.args.status)
			g.Expect(got).To(gomega.Equal(tt.want))
		})
	}
}

func Test_dataPlaneKafkaService_unassignKafkaFromDataplaneCluster(t *testing.T) {
	type fields struct {
		kafkaService *KafkaServiceMock
	}
	type args struct {
		kafka *dbapi.KafkaRequest
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *errors.ServiceError
	}{
		{
			name: "should remove the kafka from the current assigned cluster",
			fields: fields{
				kafkaService: &KafkaServiceMock{
					UpdatesFunc: func(kafkaRequest *dbapi.KafkaRequest, values map[string]interface{}) *errors.ServiceError {
						return nil
					},
				},
			},
			args: args{
				kafka: &dbapi.KafkaRequest{
					Status:    "provisioning",
					ClusterID: "test-cluster-id",
				},
			},
			want: nil,
		},
		{
			name: "should return nil if kafka status is not provisioning",
			args: args{
				kafka: &dbapi.KafkaRequest{
					Status:    "ready",
					ClusterID: "test-cluster-id",
				},
			},
			want: nil,
		},
		{
			name: "should return nil if kafka status is kafka is enterprise",
			args: args{
				kafka: &dbapi.KafkaRequest{
					Status:                   "provisioning",
					ClusterID:                "test-cluster-id",
					DesiredKafkaBillingModel: constants.BillingModelEnterprise.String(),
				},
			},
			want: nil,
		},
		{
			name: "should return error if updateFunc returns error",
			fields: fields{
				kafkaService: &KafkaServiceMock{
					UpdatesFunc: func(kafkaRequest *dbapi.KafkaRequest, values map[string]interface{}) *errors.ServiceError {
						return errors.GeneralError("test")
					},
				},
			},
			args: args{
				kafka: &dbapi.KafkaRequest{
					Status:    "provisioning",
					ClusterID: "test-cluster-id",
				},
			},
			want: errors.NewWithCause(errors.ErrorGeneral, errors.GeneralError("test"), "failed to reset fields for kafka \"\""),
		},
	}

	for _, testcase := range tests {
		tt := testcase
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			g := gomega.NewWithT(t)
			d := &dataPlaneKafkaService{
				kafkaService: tt.fields.kafkaService,
			}
			got := d.unassignKafkaFromDataplaneCluster(tt.args.kafka)
			g.Expect(got).To(gomega.Equal(tt.want))
		})
	}
}

func Test_dataPlaneKafkaService_getManagedKafkaDeploymentType(t *testing.T) {
	t.Parallel()
	type args struct {
		ks *dbapi.DataPlaneKafkaStatus
	}
	tests := []struct {
		name string
		args args
		want managedKafkaDeploymentType
	}{
		{
			name: "should return 'reserved' if the status id begins with 'reserved-kafka-'",
			args: args{
				&dbapi.DataPlaneKafkaStatus{
					KafkaClusterId: "reserved-kafka-standard",
				},
			},
			want: reservedDeploymentType,
		},
		{
			name: "should return 'real' if the status id does not begin with 'reserved-kafka-'",
			args: args{
				&dbapi.DataPlaneKafkaStatus{
					KafkaClusterId: "kafka-id",
				},
			},
			want: realDeploymentType,
		},
	}
	for _, tt := range tests {
		testcase := tt
		t.Run(testcase.name, func(t *testing.T) {
			t.Parallel()
			g := gomega.NewWithT(t)
			d := &dataPlaneKafkaService{}
			deploymentType := d.getManagedKafkaDeploymentType(testcase.args.ks)
			g.Expect(deploymentType).To(gomega.Equal(testcase.want))
		})
	}
}

func Test_dataPlaneKafkaService_processReservedKafkaDeployment(t *testing.T) {
	type args struct {
		ks                   *dbapi.DataPlaneKafkaStatus
		prewarmingStatusInfo reservedManagedKafkaStatusCountPerInstanceType
	}
	tests := []struct {
		name string
		args args
		want reservedManagedKafkaStatusCountPerInstanceType
	}{
		{
			name: "should increase the 'ready' count of standard by one",
			args: args{
				prewarmingStatusInfo: reservedManagedKafkaStatusCountPerInstanceType{
					api.StandardTypeSupport: {
						statusReady: 0,
					},
				},
				ks: &dbapi.DataPlaneKafkaStatus{
					KafkaClusterId: "reserved-kafka-standard-1",
					Conditions: []dbapi.DataPlaneKafkaStatusCondition{
						{
							Status: "True",
							Type:   "Ready",
						},
					},
				},
			},
			want: reservedManagedKafkaStatusCountPerInstanceType{
				api.StandardTypeSupport: {
					statusReady: 1,
				},
			},
		},
		{
			name: "should increase the 'ready' count of developer by one",
			args: args{
				prewarmingStatusInfo: reservedManagedKafkaStatusCountPerInstanceType{
					api.StandardTypeSupport: {
						statusReady: 0,
					},
					api.DeveloperTypeSupport: {
						statusReady: 1,
					},
				},
				ks: &dbapi.DataPlaneKafkaStatus{
					KafkaClusterId: "reserved-kafka-developer-1",
					Conditions: []dbapi.DataPlaneKafkaStatusCondition{
						{
							Status: "True",
							Type:   "Ready",
						},
					},
				},
			},
			want: reservedManagedKafkaStatusCountPerInstanceType{
				api.StandardTypeSupport: {
					statusReady: 0,
				},
				api.DeveloperTypeSupport: {
					"ready": 2,
				},
			},
		},
		{
			name: "should increase the 'error' count of standard by one",
			args: args{
				prewarmingStatusInfo: reservedManagedKafkaStatusCountPerInstanceType{
					api.StandardTypeSupport: {
						statusError: 0,
					},
				},
				ks: &dbapi.DataPlaneKafkaStatus{
					KafkaClusterId: "reserved-kafka-standard-1",
					Conditions: []dbapi.DataPlaneKafkaStatusCondition{
						{
							Status: "False",
							Reason: "Error",
							Type:   "Ready",
						},
					},
				},
			},
			want: reservedManagedKafkaStatusCountPerInstanceType{
				api.StandardTypeSupport: {
					statusError: 1,
				},
			},
		},
		{
			name: "should increase the 'rejected' count of standard by one",
			args: args{
				prewarmingStatusInfo: reservedManagedKafkaStatusCountPerInstanceType{
					api.StandardTypeSupport: {
						statusRejected: 0,
					},
				},
				ks: &dbapi.DataPlaneKafkaStatus{
					KafkaClusterId: "reserved-kafka-standard-1",
					Conditions: []dbapi.DataPlaneKafkaStatusCondition{
						{
							Status: "False",
							Reason: "Rejected",
							Type:   "Ready",
						},
					},
				},
			},
			want: reservedManagedKafkaStatusCountPerInstanceType{
				api.StandardTypeSupport: {
					statusRejected: 1,
				},
			},
		},
		{
			name: "should increase the 'rejected' count of developer by one",
			args: args{
				prewarmingStatusInfo: reservedManagedKafkaStatusCountPerInstanceType{
					api.DeveloperTypeSupport: {
						statusRejected: 0,
					},
				},
				ks: &dbapi.DataPlaneKafkaStatus{
					KafkaClusterId: "reserved-kafka-developer-1",
					Conditions: []dbapi.DataPlaneKafkaStatusCondition{
						{
							Status: "False",
							Reason: "Rejected",
							Type:   "Ready",
						},
					},
				},
			},
			want: reservedManagedKafkaStatusCountPerInstanceType{
				api.DeveloperTypeSupport: {
					statusRejected: 1,
				},
			},
		},
		{
			name: "should increase the 'rejectedClusterFull' count of standard by one",
			args: args{
				prewarmingStatusInfo: reservedManagedKafkaStatusCountPerInstanceType{
					api.StandardTypeSupport: {
						statusRejectedClusterFull: 1,
					},
				},
				ks: &dbapi.DataPlaneKafkaStatus{
					KafkaClusterId: "reserved-kafka-standard-1",
					Conditions: []dbapi.DataPlaneKafkaStatusCondition{
						{
							Status:  "False",
							Reason:  "Rejected",
							Message: "Cluster has insufficient resources",
							Type:    "Ready",
						},
					},
				},
			},
			want: reservedManagedKafkaStatusCountPerInstanceType{
				api.StandardTypeSupport: {
					statusRejectedClusterFull: 2,
				},
			},
		},
		{
			name: "should increase the 'rejectedClusterFull' count of developer by one",
			args: args{
				prewarmingStatusInfo: reservedManagedKafkaStatusCountPerInstanceType{
					api.DeveloperTypeSupport: {
						statusRejectedClusterFull: 1,
					},
				},
				ks: &dbapi.DataPlaneKafkaStatus{
					KafkaClusterId: "reserved-kafka-developer-1",
					Conditions: []dbapi.DataPlaneKafkaStatusCondition{
						{
							Status:  "False",
							Reason:  "Rejected",
							Message: "Cluster has insufficient resources",
							Type:    "Ready",
						},
					},
				},
			},
			want: reservedManagedKafkaStatusCountPerInstanceType{
				api.DeveloperTypeSupport: {
					statusRejectedClusterFull: 2,
				},
			},
		},
		{
			name: "should increase the 'deleted' count of standard by one",
			args: args{
				prewarmingStatusInfo: reservedManagedKafkaStatusCountPerInstanceType{
					api.StandardTypeSupport: {
						statusDeleted: 0,
					},
				},
				ks: &dbapi.DataPlaneKafkaStatus{
					KafkaClusterId: "reserved-kafka-standard-1",
					Conditions: []dbapi.DataPlaneKafkaStatusCondition{
						{
							Status: "False",
							Reason: "Deleted",
							Type:   "Ready",
						},
					},
				},
			},
			want: reservedManagedKafkaStatusCountPerInstanceType{
				api.StandardTypeSupport: {
					statusDeleted: 1,
				},
			},
		},
		{
			name: "should increase the 'deleted' count of developer by one",
			args: args{
				prewarmingStatusInfo: reservedManagedKafkaStatusCountPerInstanceType{
					api.DeveloperTypeSupport: {
						statusDeleted: 0,
					},
				},
				ks: &dbapi.DataPlaneKafkaStatus{
					KafkaClusterId: "reserved-kafka-developer-1",
					Conditions: []dbapi.DataPlaneKafkaStatusCondition{
						{
							Status: "False",
							Reason: "Deleted",
							Type:   "Ready",
						},
					},
				},
			},
			want: reservedManagedKafkaStatusCountPerInstanceType{
				api.DeveloperTypeSupport: {
					statusDeleted: 1,
				},
			},
		},
		{
			name: "should increase the 'error' count of developer by one",
			args: args{
				prewarmingStatusInfo: reservedManagedKafkaStatusCountPerInstanceType{
					api.DeveloperTypeSupport: {
						statusError: 0,
					},
				},
				ks: &dbapi.DataPlaneKafkaStatus{
					KafkaClusterId: "reserved-kafka-developer-1",
					Conditions: []dbapi.DataPlaneKafkaStatusCondition{
						{
							Status: "False",
							Reason: "Error",
							Type:   "Ready",
						},
					},
				},
			},
			want: reservedManagedKafkaStatusCountPerInstanceType{
				api.DeveloperTypeSupport: {
					statusError: 1,
				},
			},
		},
		{
			name: "should increase the 'installing' count of developer by one",
			args: args{
				prewarmingStatusInfo: reservedManagedKafkaStatusCountPerInstanceType{
					api.StandardTypeSupport: {
						statusReady: 0,
					},
					api.DeveloperTypeSupport: {
						statusInstalling: 1,
					},
				},
				ks: &dbapi.DataPlaneKafkaStatus{
					KafkaClusterId: "reserved-kafka-developer-1",
					Conditions: []dbapi.DataPlaneKafkaStatusCondition{
						{
							Status: "False",
							Type:   "Ready",
						},
					},
				},
			},
			want: reservedManagedKafkaStatusCountPerInstanceType{
				api.StandardTypeSupport: {
					statusReady: 0,
				},
				api.DeveloperTypeSupport: {
					statusInstalling: 2,
				},
			},
		},
		{
			name: "should increase the 'installing' count of standard by one",
			args: args{
				prewarmingStatusInfo: reservedManagedKafkaStatusCountPerInstanceType{
					api.StandardTypeSupport: {
						statusInstalling: 0,
					},
					api.DeveloperTypeSupport: {
						statusInstalling: 1,
					},
				},
				ks: &dbapi.DataPlaneKafkaStatus{
					KafkaClusterId: "reserved-kafka-standard-1",
					Conditions: []dbapi.DataPlaneKafkaStatusCondition{
						{
							Status: "False",
							Type:   "Ready",
						},
					},
				},
			},
			want: reservedManagedKafkaStatusCountPerInstanceType{
				api.StandardTypeSupport: {
					statusInstalling: 1,
				},
				api.DeveloperTypeSupport: {
					statusInstalling: 1,
				},
			},
		},
		{
			name: "should increase the 'unknown' count of developer by one",
			args: args{
				prewarmingStatusInfo: reservedManagedKafkaStatusCountPerInstanceType{
					api.StandardTypeSupport: {
						statusReady: 0,
					},
					api.DeveloperTypeSupport: {
						statusUnknown: 5,
					},
				},
				ks: &dbapi.DataPlaneKafkaStatus{
					KafkaClusterId: "reserved-kafka-developer-1",
					Conditions: []dbapi.DataPlaneKafkaStatusCondition{
						{
							Status: "Unknown",
							Reason: "",
							Type:   "Ready",
						},
					},
				},
			},
			want: reservedManagedKafkaStatusCountPerInstanceType{
				api.StandardTypeSupport: {
					statusReady: 0,
				},
				api.DeveloperTypeSupport: {
					statusUnknown: 6,
				},
			},
		},
		{
			name: "should increase the 'unknown' count of standard by one",
			args: args{
				prewarmingStatusInfo: reservedManagedKafkaStatusCountPerInstanceType{
					api.StandardTypeSupport: {
						statusUnknown: 0,
					},
					api.DeveloperTypeSupport: {
						statusInstalling: 1,
					},
				},
				ks: &dbapi.DataPlaneKafkaStatus{
					KafkaClusterId: "reserved-kafka-standard-1",
					Conditions: []dbapi.DataPlaneKafkaStatusCondition{
						{
							Status: "Unknown",
							Type:   "Ready",
						},
					},
				},
			},
			want: reservedManagedKafkaStatusCountPerInstanceType{
				api.StandardTypeSupport: {
					statusUnknown: 1,
				},
				api.DeveloperTypeSupport: {
					statusInstalling: 1,
				},
			},
		},
		{
			name: "shouldn't change the prewarming status info count when the kafka id does not follow the reserved kafka ID pattern",
			args: args{
				prewarmingStatusInfo: reservedManagedKafkaStatusCountPerInstanceType{
					api.StandardTypeSupport: {
						statusUnknown: 0,
					},
					api.DeveloperTypeSupport: {
						statusInstalling: 1,
					},
				},
				ks: &dbapi.DataPlaneKafkaStatus{
					KafkaClusterId: "reservedstandard1",
					Conditions: []dbapi.DataPlaneKafkaStatusCondition{
						{
							Status: "Unknown",
							Type:   "Ready",
						},
					},
				},
			},
			want: reservedManagedKafkaStatusCountPerInstanceType{
				api.StandardTypeSupport: {
					statusUnknown: 0,
				},
				api.DeveloperTypeSupport: {
					statusInstalling: 1,
				},
			},
		},
		{
			name: "shouldn't change the prewarming status info count when the reserved kafka is not supported",
			args: args{
				prewarmingStatusInfo: reservedManagedKafkaStatusCountPerInstanceType{
					api.StandardTypeSupport: {
						statusUnknown: 0,
					},
					api.DeveloperTypeSupport: {
						statusInstalling: 1,
					},
				},
				ks: &dbapi.DataPlaneKafkaStatus{
					KafkaClusterId: "reserved-kafka-standard1",
					Conditions: []dbapi.DataPlaneKafkaStatusCondition{
						{
							Status: "Unknown",
							Type:   "Ready",
						},
					},
				},
			},
			want: reservedManagedKafkaStatusCountPerInstanceType{
				api.StandardTypeSupport: {
					statusUnknown: 0,
				},
				api.DeveloperTypeSupport: {
					statusInstalling: 1,
				},
			},
		},
	}
	for _, tt := range tests {
		testcase := tt
		t.Run(testcase.name, func(t *testing.T) {
			t.Parallel()
			g := gomega.NewWithT(t)
			d := &dataPlaneKafkaService{}
			d.processReservedKafkaDeployment(testcase.args.ks, testcase.args.prewarmingStatusInfo, logger.Logger, "some-cluster-id")
			g.Expect(testcase.args.prewarmingStatusInfo).To(gomega.Equal(testcase.want))
		})
	}
}
