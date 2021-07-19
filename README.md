# Go-SDK-POC

## Examples: SDK To Create ChasoEngine
#### How to use?

- The ChaosEngine is the main user-facing chaos custom resource with a namespace scope and is designed to hold information around how the chaos experiments are executed.


```golang
package main

import (
	"fmt"

	"github.com/litmuschaos/chaos-operator/pkg/apis/litmuschaos/v1alpha1"
	"github.com/uditgaurav/go-sdk-poc/pkg/clients"
	"github.com/uditgaurav/go-sdk-poc/pkg/litmus"
	"github.com/uditgaurav/go-sdk-poc/pkg/log"
)

func main() {

	clients := clients.ClientSets{}
	if err := clients.GenerateClientSetFromKubeConfig(); err != nil {
		log.Errorf("fail to set litmus client error: %v", err)
		return
	}
	chaosEngine := v1alpha1.ChaosEngine{}
	engineManifest := litmus.CreateChaosEngineManifest{}

	engineManifest.Name = "my-chaos-test-name"
	engineManifest.Namespace = "default"
	engineManifest.JobCleanUpPolicy = v1alpha1.CleanUpPolicyRetain
	engineManifest.ExperimentName = "pod-delete"
	engineManifest.ENVs = make(map[string]string)
	engineManifest.ENVs["TOTAL_CHAOS_DURATION"] = "120"
	engineManifest.ENVs["CHAOS_INTERVAL"] = "30"

	resp, err := litmus.CreateChaosEngine(&chaosEngine, engineManifest, clients)
	if err != nil {
		log.Errorf("Error in creating chaos engine: %v", err)
	}
	fmt.Println(resp)
}
```

Refer [Litmus Docs](https://docs.litmuschaos.io/docs) to Prepare the engine structure.
