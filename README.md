# Go-SDK-POC

## How to use?

- The ChaosEngine is the main user-facing chaos custom resource with a namespace scope and is designed to hold information around how the chaos experiments are executed.


```golang
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
```

Refer [Litmus Docs](https://docs.litmuschaos.io/docs/chaosengine/) to Prepare the engine structure.
