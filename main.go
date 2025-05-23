package main

import (
	"github.com/muraduiurie/kubegpt/pkg/controller"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/log/zap"
)

var scheme = runtime.NewScheme()

func main() {
	// create main logger
	logger := zap.New()
	ctrl.SetLogger(logger)
	log := ctrl.Log.WithName("main")

	// create manager
	log.Info("set up manager")

	// load schemes
	utilruntime.Must(corev1.AddToScheme(scheme))

	mgr, err := ctrl.NewManager(ctrl.GetConfigOrDie(), ctrl.Options{
		Scheme: scheme,
	})
	if err != nil {
		log.Error(err, "unable to create controller manager")
	}

	// run controller
	ctx := ctrl.SetupSignalHandler()
	err = controller.RunController(ctx, log.WithName("controller"), mgr, scheme)
	if err != nil {
		log.Error(err, "unable to run controller")
	}
}
