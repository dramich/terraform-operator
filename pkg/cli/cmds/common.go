package cmds

import (
	"context"
	"github.com/rancher/terraform-controller/pkg/generated/controllers/terraformcontroller.cattle.io"
	"github.com/rancher/terraform-controller/pkg/generated/controllers/terraformcontroller.cattle.io/v1"
	"github.com/rancher/wrangler/pkg/signals"
	"github.com/rancher/wrangler/pkg/start"
	"github.com/sirupsen/logrus"
	"k8s.io/client-go/tools/clientcmd"
)

type controllers struct {
	modules       v1.ModuleController
	executions    v1.ExecutionController
	executionRuns v1.ExecutionRunController
}

func getControllers(kubeconfig, ns string) (*controllers, error) {
	//todo add masterurl flag?

	cfg, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		logrus.Fatalf("Error building kubeconfig: %s", err.Error())
	}

	tfFactory, err := terraformcontroller.NewFactoryFromConfigWithNamespace(cfg, ns)
	if err != nil {
		logrus.Fatalf("Error building terraform controllers: %s", err.Error())
	}

	controllers := &controllers{
		modules:       tfFactory.Terraformcontroller().V1().Module(),
		executions:    tfFactory.Terraformcontroller().V1().Execution(),
		executionRuns: tfFactory.Terraformcontroller().V1().ExecutionRun(),
	}

	ctx := signals.SetupSignalHandler(context.Background())
	if err := start.All(ctx, 1, tfFactory); err != nil {
		logrus.Fatalf("Error starting: %s", err.Error())
	}

	return controllers, nil
}
