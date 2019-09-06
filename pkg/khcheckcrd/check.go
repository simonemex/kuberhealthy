package khcheckcrd

import (
	"time"

	apiv1 "k8s.io/api/core/v1"
	// metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	// typedv1 "k8s.io/client-go/kubernetes/typed/core/v1"
)

// CheckConfig represents a configuration for a kuberhealthy external
// checker.  This includes the pod spec to run, the interval, and
// the whitelisted UUID that is currently allowed to report-in to
// the status reporting endpoint.
type CheckConfig struct {
	RunInterval string        // the interval at which the check runs
	PodSpec     apiv1.PodSpec // a spec for the external checker
	CurrentUUID string        // the UUID that is authorized to report statuses into the kuberhealthy endpoint
}

// NewCheckConfig creates a new check configuration
func NewCheckConfig() CheckConfig {
	c := CheckConfig{
		RunInterval: (time.Minute * 10).String(),
		PodSpec:     apiv1.PodSpec{},
	}

	return c
}