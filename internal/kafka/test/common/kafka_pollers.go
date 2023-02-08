package common

import (
	"context"
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/bf2fc6cc711aee1a0c2a/kas-fleet-manager/internal/kafka/constants"
	"github.com/bf2fc6cc711aee1a0c2a/kas-fleet-manager/internal/kafka/internal/api/dbapi"
	"github.com/bf2fc6cc711aee1a0c2a/kas-fleet-manager/internal/kafka/internal/api/public"
	"github.com/bf2fc6cc711aee1a0c2a/kas-fleet-manager/pkg/db"
	"github.com/bf2fc6cc711aee1a0c2a/kas-fleet-manager/test"
	"github.com/golang/glog"
	"github.com/pkg/errors"
)

const (
	defaultKafkaReadyTimeout             = 30 * time.Minute
	defaultKafkaClusterAssignmentTimeout = 2 * time.Minute
	metricPollInterval                   = 1 * time.Second
	metricPollTimeout                    = 1 * time.Minute // make it 1 minute to take into account the reconciler run interval which can be 30 seconds in some environments
)

// WaitForNumberOfKafkaToBeGivenCount - Awaits for the number of kafkas to be exactly X
func WaitForNumberOfKafkaToBeGivenCount(ctx context.Context, db *db.ConnectionFactory, client *public.APIClient, count int32, customize ...func(builder PollerBuilder) PollerBuilder) error {
	currentCount := int32(-1)

	pb := NewPollerBuilder(db).
		IntervalAndTimeout(defaultPollInterval, defaultKafkaPollTimeout).
		RetryLogFunction(func(retry int, maxRetry int) string {
			if currentCount == -1 {
				return fmt.Sprintf("Waiting for kafkas count to become %d", count)
			} else {
				return fmt.Sprintf("Waiting for kafkas count to become %d (current %d)", count, currentCount)
			}
		}).
		OnRetry(func(attempt int, maxRetries int) (done bool, err error) {
			if list, resp, err := client.DefaultApi.GetKafkas(ctx, nil); err != nil {
				if resp != nil {
					_ = resp.Body.Close()
				}
				return false, err
			} else {
				currentCount = list.Size
				return currentCount == count, err
			}
		})

	if len(customize) > 0 {
		pb = customize[0](pb)
	}

	return pb.Build().Poll()
}

// WaitForKafkaCreateToBeAccepted - Creates a kafka and awaits for the request to be accepted
func WaitForKafkaCreateToBeAccepted(ctx context.Context, db *db.ConnectionFactory, client *public.APIClient, k public.KafkaRequestPayload) (kafka public.KafkaRequest, resp *http.Response, err error) {
	currentStatus := ""

	err = NewPollerBuilder(db).
		IntervalAndTimeout(defaultPollInterval, defaultKafkaPollTimeout).
		RetryLogFunction(func(retry int, maxRetry int) string {
			if currentStatus == "" {
				return "Waiting for kafka creation to be accepted"
			} else {
				return fmt.Sprintf("Waiting for kafka creation to be accepted (current status %s)", currentStatus)
			}
		}).
		OnRetry(func(attempt int, maxRetries int) (done bool, err error) {
			//nolint: bodyclose
			kafka, resp, err = client.DefaultApi.CreateKafka(ctx, true, k)
			if err != nil {
				return true, err
			}
			return resp.StatusCode == http.StatusAccepted, err
		}).
		Build().Poll()
	return kafka, resp, err
}

// WaitForKafkaToReachStatus - Awaits for a kafka to reach a specified status
func WaitForKafkaToReachStatus(ctx context.Context, db *db.ConnectionFactory, client *public.APIClient, kafkaId string, status constants.KafkaStatus) (kafka public.KafkaRequest, err error) {
	currentStatus := ""

	glog.Infof("status: " + status.String())

	err = NewPollerBuilder(db).
		IntervalAndTimeout(1*time.Second, defaultKafkaReadyTimeout).
		RetryLogFunction(func(retry int, maxRetry int) string {
			if currentStatus == "" {
				return fmt.Sprintf("Waiting for kafka '%s' to reach status '%s'", kafkaId, status.String())
			} else {
				return fmt.Sprintf("Waiting for kafka '%s' to reach status '%s' (current status %s)", kafkaId, status.String(), currentStatus)
			}
		}).
		OnRetry(func(attempt int, maxRetries int) (done bool, err error) {
			k, resp, err := client.DefaultApi.GetKafkaById(ctx, kafkaId)

			if resp != nil {
				resp.Body.Close()
			}
			if err != nil {
				return true, err
			}

			kafka = k
			switch kafka.Status {
			case constants.KafkaRequestStatusFailed.String():
				fallthrough
			case constants.KafkaRequestStatusDeprovision.String():
				fallthrough
			case constants.KafkaRequestStatusDeleting.String():
				return false, errors.Errorf("Waiting for kafka '%s' to reach status '%s', but status '%s' has been reached instead", kafkaId, status.String(), kafka.Status)
			}

			currentStatus = kafka.Status
			return constants.KafkaStatus(kafka.Status).CompareTo(status) >= 0, nil
		}).
		Build().Poll()
	return kafka, err
}

// WaitForKafkaToBeDeleted - Awaits for a kafka to be deleted
func WaitForKafkaToBeDeleted(ctx context.Context, db *db.ConnectionFactory, client *public.APIClient, kafkaId string) error {
	return NewPollerBuilder(db).
		IntervalAndTimeout(defaultPollInterval, defaultKafkaReadyTimeout).
		RetryLogMessagef("Waiting for kafka '%s' to be deleted", kafkaId).
		OnRetry(func(attempt int, maxRetries int) (done bool, err error) {
			if _, resp, err := client.DefaultApi.GetKafkaById(ctx, kafkaId); err != nil {
				if resp != nil {
					resp.Body.Close()
				}
				if err.Error() == "404 Not Found" {
					return true, nil
				}

				return false, err
			}
			return false, nil
		}).
		Build().Poll()
}

func WaitForKafkaClusterIDToBeAssigned(dbFactory *db.ConnectionFactory, kafkaRequestName string) (*dbapi.KafkaRequest, error) {
	kafkaFound := &dbapi.KafkaRequest{}

	kafkaErr := NewPollerBuilder(dbFactory).
		IntervalAndTimeout(defaultPollInterval, defaultKafkaClusterAssignmentTimeout).
		RetryLogMessagef("Waiting for kafka named '%s' to have a ClusterID", kafkaRequestName).
		OnRetry(func(attempt int, maxRetries int) (done bool, err error) {
			if err := dbFactory.New().Where("name = ?", kafkaRequestName).First(kafkaFound).Error; err != nil {
				return false, err
			}
			glog.Infof("got kafka instance %v", kafkaFound)
			return kafkaFound.ClusterID != "", nil
		}).Build().Poll()

	return kafkaFound, kafkaErr
}

func WaitForMetricToBePresent(h *test.Helper, t *testing.T, metric string, values ...string) error {
	dbConn := h.DBFactory()
	return NewPollerBuilder(dbConn).
		IntervalAndTimeout(metricPollInterval, metricPollTimeout).
		RetryLogMessagef("Waiting for metric '%s' to contain values '%s", metric, values).
		OnRetry(func(attempt int, maxRetries int) (done bool, err error) {
			return IsMetricExposedWithValue(t, metric, values...), nil
		}).
		Build().Poll()
}
