package litmus

import (
	"github.com/litmuschaos/chaos-operator/pkg/apis/litmuschaos/v1alpha1"
	"github.com/uditgaurav/go-sdk-poc/pkg/clients"

	corev1 "k8s.io/api/core/v1"
)

func CreateChaosEngine(chaosEngine *v1alpha1.ChaosEngine, clients clients.ClientSets) (*v1alpha1.ChaosEngine, error) {

	if chaosEngine.APIVersion == "" {
		chaosEngine.APIVersion = "litmuschaos.io/v1alpha1"
	}
	if chaosEngine.Kind == "" {
		chaosEngine.Kind = "ChaosEngine"
	}
	if chaosEngine.Spec.EngineState == "" {
		chaosEngine.Spec.EngineState = v1alpha1.EngineStateActive
	}
	if chaosEngine.Spec.JobCleanUpPolicy == "" {
		chaosEngine.Spec.JobCleanUpPolicy = v1alpha1.CleanUpPolicy("retain")
	}
	if chaosEngine.Spec.Components.Runner.ImagePullPolicy == "" {
		chaosEngine.Spec.Components.Runner.ImagePullPolicy = corev1.PullPolicy("Always")
	}

	if chaosEngine.ObjectMeta.Name == "" {
		chaosEngine.ObjectMeta.Name = "nginx-chaos"
	}
	if chaosEngine.ObjectMeta.Namespace == "" {
		chaosEngine.ObjectMeta.Namespace = "default"
	}
	if chaosEngine.Spec.Appinfo.Appns == "" {
		chaosEngine.Spec.Appinfo.Appns = "default"
	}
	if chaosEngine.Spec.Appinfo.Applabel == "" {
		chaosEngine.Spec.Appinfo.Applabel = "app=nginx"
	}
	if chaosEngine.Spec.Appinfo.AppKind == "" {
		chaosEngine.Spec.Appinfo.AppKind = "deployment"
	}
	if chaosEngine.Spec.ChaosServiceAccount == "" {
		chaosEngine.Spec.ChaosServiceAccount = "litmus-admin"
	}
	expList := v1alpha1.ExperimentList{}
	expList.Name = "pod-delete"
	chaosEngine.Spec.Experiments = append(chaosEngine.Spec.Experiments, expList)
	if chaosEngine.Spec.AnnotationCheck == "" {
		chaosEngine.Spec.AnnotationCheck = "false"
	}

	resp, err := clients.LitmusClient.ChaosEngines("default").Create(chaosEngine)
	if err != nil {
		return resp, err
	}

	return resp, nil
}
