package litmus

import (
	"github.com/litmuschaos/chaos-operator/pkg/apis/litmuschaos/v1alpha1"
	"github.com/uditgaurav/go-sdk-poc/pkg/clients"

	corev1 "k8s.io/api/core/v1"
	v1 "k8s.io/api/core/v1"
)

func CreateChaosEngine(chaosEngine *v1alpha1.ChaosEngine, engineManifest CreateChaosEngineManifest, clients clients.ClientSets) (*v1alpha1.ChaosEngine, error) {

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
	if engineManifest.Name != "" {
		chaosEngine.ObjectMeta.Name = engineManifest.Name
	} else {
		chaosEngine.ObjectMeta.Name = "nginx-chaos"
	}
	if engineManifest.Namespace != "" {
		chaosEngine.ObjectMeta.Namespace = engineManifest.Namespace
	} else {
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
	if engineManifest.ExperimentName != "" {
		expList.Name = engineManifest.ExperimentName
	} else {
		expList.Name = "pod-delete"
	}
	chaosEngine.Spec.Experiments = append(chaosEngine.Spec.Experiments, expList)
	for key, value := range engineManifest.ENVs {
		chaosEngine.Spec.Experiments[0].Spec.Components.ENV = append(chaosEngine.Spec.Experiments[0].Spec.Components.ENV, v1.EnvVar{Name: key, Value: value})
	}
	if chaosEngine.Spec.AnnotationCheck == "" {
		chaosEngine.Spec.AnnotationCheck = "false"
	}

	resp, err := clients.LitmusClient.ChaosEngines(chaosEngine.Namespace).Create(chaosEngine)
	if err != nil {
		return resp, err
	}

	return resp, nil
}
