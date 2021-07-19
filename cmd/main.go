package main

import (
	"fmt"

	"github.com/litmuschaos/chaos-operator/pkg/apis/litmuschaos/v1alpha1"
	"github.com/uditgaurav/go-sdk-poc/pkg/clients"
	"github.com/uditgaurav/go-sdk-poc/pkg/litmus"
	"github.com/uditgaurav/go-sdk-poc/pkg/log"
	v1 "k8s.io/api/core/v1"
)

func main() {

	clients := clients.ClientSets{}
	if err := clients.GenerateClientSetFromKubeConfig(); err != nil {
		log.Errorf("fail to set litmus client error: %v", err)
		return
	}
	chaosEngine := v1alpha1.ChaosEngine{}

	chaosEngine.Name = "my-chaos-test"
	chaosEngine.Spec.JobCleanUpPolicy = v1alpha1.CleanUpPolicyRetain

	expList := v1alpha1.ExperimentList{}
	expList.Name = "pod-delete"
	env := []v1.EnvVar{
		{
			Name:  "TOTAL_CHAOS_DURATION",
			Value: "120",
		},
		{
			Name:  "CHAOS_INTERVAL",
			Value: "30",
		},
	}

	for i := range env {
		expList.Spec.Components.ENV = append(expList.Spec.Components.ENV, env[i])
	}
	chaosEngine.Spec.Experiments = append(chaosEngine.Spec.Experiments, expList)

	resp, err := litmus.CreateChaosEngine(&chaosEngine, clients)
	if err != nil {
		log.Errorf("Error in creating chaos engine: %v", err)
	}
	fmt.Println(resp)
}
