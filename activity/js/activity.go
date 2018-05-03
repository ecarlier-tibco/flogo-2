package js

import (
	"fmt"

	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	"github.com/TIBCOSoftware/flogo-lib/logger"
	"github.com/robertkrimen/otto"
)

var activityLog = logger.GetLogger("activity-vijay-js")

const (
	ivActivityInput   = "jsIn"
	ivJsInputVarName  = "jsInputVarName"
	ivJs              = "jsCode"
	ivJsOutputVarName = "jsOutputVarName"
	ovActitvityOutput = "jsOut"
)

// JSActivity : Javascript Acitivity
type JSActivity struct {
	metadata *activity.Metadata
}

// NewActivity creates a new AppActivity
func NewActivity(metadata *activity.Metadata) activity.Activity {
	return &JSActivity{metadata: metadata}
}

// Metadata returns the activity's metadata
func (a *JSActivity) Metadata() *activity.Metadata {
	return a.metadata
}

// Eval implements api.Activity.Eval - Execute JS code
func (a *JSActivity) Eval(context activity.Context) (done bool, err error) {

	inJsVarName, _ := context.GetInput(ivJsInputVarName).(string)
	outJsVarName, _ := context.GetInput(ivJsOutputVarName).(string)
	inputVars, _ := context.GetInput(ivActivityInput).(interface{})
	jsCode, _ := context.GetInput(ivJs).(string)

	activityLog.Debugf("JavaScript Input Var Value: %v", inputVars)
	fmt.Printf("JavaScript Input Var Value: %v\n", inputVars)
	fmt.Printf("JavaScript Input Var Name: %s\n", inJsVarName)
	fmt.Printf("JavaScript Code: %s\n", jsCode)
	activityLog.Debugf("JavaScript Code: %s", jsCode)

	vm := otto.New()

	//Set Input Variable
	vm.Set(inJsVarName, inputVars)

	v, err := vm.Run(jsCode)
	if err != nil {
		return false, activity.NewError(fmt.Sprintf("Failed to execute JavaScript code due to error: %s", err.Error()), "", nil)
	}

	var jsOutput interface{}
	// Look for jsOutput variable value
	value, err := vm.Get(outJsVarName)
	if value.IsNull() || value.IsUndefined() {
		// Set returned value
		jsOutput, _ = v.Export()
	} else {
		// Set jsOutput variable value
		jsOutput, _ = value.Export()
	}
	context.SetOutput(ovActitvityOutput, jsOutput)

	activityLog.Debugf("JavaScript Output: %v", jsOutput)
	return true, nil
}
