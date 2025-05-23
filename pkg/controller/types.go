package controller

import (
	"github.com/go-logr/logr"
	"github.com/muraduiurie/kubegpt/pkg/ai"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/tools/record"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type AiReconciler struct {
	client.Client
	Scheme        *runtime.Scheme
	eventRecorder record.EventRecorder
	Log           logr.Logger
	AiClient      ai.AiClient
}
