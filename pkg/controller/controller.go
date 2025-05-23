package controller

import (
	"context"
	"fmt"
	"github.com/go-logr/logr"
	"github.com/muraduiurie/kubegpt/pkg/ai"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"os"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/controller"
	"sigs.k8s.io/controller-runtime/pkg/manager"
)

func RunController(ctx context.Context, log logr.Logger, mgr manager.Manager, scheme *runtime.Scheme) error {
	log.Info("setting up controller")
	err := setupAiController(mgr, scheme, log)
	if err != nil {
		return fmt.Errorf("unable to setup ai controller: %w", err)
	}

	// start manager
	log.Info("starting controller manager")
	if err = mgr.Start(ctx); err != nil {
		return fmt.Errorf("error running controller manager: %w", err)
	}

	return nil
}

func setupAiController(mgr manager.Manager, scheme *runtime.Scheme, log logr.Logger) error {
	eventRecorder := mgr.GetEventRecorderFor("pod")

	c, err := ai.InitAiClient(os.Getenv("KUBEGPT_AI_PROVIDER"), log.WithName("ai"))
	if err != nil {
		return fmt.Errorf("unable to initialize ai client: %w", err)
	}

	ar := AiReconciler{
		Client:        mgr.GetClient(),
		Scheme:        scheme,
		eventRecorder: eventRecorder,
		Log:           log.WithName("ai-reconciler"),
		AiClient:      c,
	}

	// create controller
	err = ar.SetupWithManager(mgr)
	if err != nil {
		return fmt.Errorf("unable to create ai controller: %w", err)
	}

	return nil
}

// SetupWithManager sets up the controller with the Manager.
func (ar *AiReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&corev1.Pod{}).
		WithOptions(controller.Options{
			MaxConcurrentReconciles: 1,
		}).
		Complete(ar)
}

func (ar *AiReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	ar.Log.Info("reconciling pod", "name", req.Name, "namespace", req.Namespace)

	return ctrl.Result{}, nil
}
