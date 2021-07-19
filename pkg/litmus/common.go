package litmus

import (
	"github.com/litmuschaos/chaos-operator/pkg/apis/litmuschaos/v1alpha1"
	corev1 "k8s.io/api/core/v1"
)

// CreateChaosEngineManifest will prepare the manifest or chaosengine
type CreateChaosEngineManifest struct {
	Name                  string
	Namespace             string
	EngineState           v1alpha1.EngineState
	JobCleanUpPolicy      v1alpha1.CleanUpPolicy
	RunnerImagePullPolicy corev1.PullPolicy
	AppNS                 string
	Applabel              string
	AppKing               string
	ChaosServiceAccount   string
	AnnotationCheck       string
	ExperimentName        string
	ENVs                  map[string]string
}
