package flogojsontransformer

import (
	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	"github.com/TIBCOSoftware/flogo-lib/logger"
	"github.com/kabukky/kazaam"
)

// activityLog is the default logger for the exec Activity
var log = logger.GetLogger("activity-tibco-flogojsontransformer")

type JsontransformerActivity struct {
	metadata *activity.Metadata
}

// NewActivity creates a new AppActivity
func NewActivity(metadata *activity.Metadata) activity.Activity {
	return &JsontransformerActivity{metadata: metadata}
}

// Metadata returns the activity's metadata
func (a *JsontransformerActivity) Metadata() *activity.Metadata {
	return a.metadata
}

func (a *JsontransformerActivity) Eval(context activity.Context) (done bool, err error) {

	log.Info("JSON Transformer starting...")

	jsonin, _ := context.GetInput("Input").(string)
	spec, _ := context.GetInput("Transformation").(string)

	kazaamTransform, _ := kazaam.NewKazaam(spec)
	kazaamOut, err := kazaamTransform.TransformJSONStringToString(jsonin)
	if err != nil {
		log.Info(err)
	}

	context.SetOutput("result", kazaamOut)

	return true, nil
}
