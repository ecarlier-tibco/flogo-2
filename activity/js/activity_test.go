package js

import (
	"fmt"
	"io/ioutil"
	"testing"

	"github.com/TIBCOSoftware/flogo-contrib/action/flow/test"
	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	"github.com/TIBCOSoftware/flogo-lib/core/data"
	"github.com/stretchr/testify/assert"
)

var activityMetadata *activity.Metadata

func getActivityMetadata() *activity.Metadata {

	if activityMetadata == nil {
		jsonMetadataBytes, err := ioutil.ReadFile("activity.json")
		if err != nil {
			panic("No Json Metadata found for activity.json path")
		}

		activityMetadata = activity.NewMetadata(string(jsonMetadataBytes))
	}

	return activityMetadata
}

func TestCreate(t *testing.T) {

	act := NewActivity(getActivityMetadata())

	if act == nil {
		t.Error("Activity Not Created")
		t.Fail()
		return
	}
}

func TestEval_InputVars(t *testing.T) {

	act := NewActivity(getActivityMetadata())
	tc := test.NewTestActivityContext(getActivityMetadata())

	//setup attrs
	inputVar := make(map[string]interface{}, 2)
	inputVar["n1"] = 2
	inputVar["n2"] = 3
	fmt.Printf("JavaScript Input Var Value: %v\n", inputVar)
	tc.SetInput(ivActivityInput, inputVar)
	tc.SetInput(ivJsInputVarName, "msg")
	tc.SetInput(ivJs, `var total = msg.n1 + msg.n2;`)
	tc.SetInput(ivJsOutputVarName, "total")
	ok, err := act.Eval(tc)
	assert.NoError(t, err)

	if ok {
		sum, _ := data.CoerceToInteger(tc.GetOutput(ovActitvityOutput))
		assert.Equal(t, 5, sum)
	}

	// test node-red function
	inputVar = map[string]interface{}{
		"req": map[string]interface{}{
			"params": map[string]interface{}{
				"firstname": "Eric",
			},
		},
		"payload": map[string]interface{}{},
	}

	fmt.Printf("JavaScript Input Var Value: %v\n", inputVar)
	tc.SetInput(ivActivityInput, inputVar)
	tc.SetInput(ivJs, "if (msg.req.params.firstname == 'Eric') {\n    msg.payload.Name = 'Carlier';\n}\nelse {\n    msg.payload.Name = 'Unknown';\n}\nmsg.statusCode = 200;")
	tc.SetInput(ivJsInputVarName, "msg")
	tc.SetInput(ivJsOutputVarName, "msg")

	ok, err = act.Eval(tc)
	assert.NoError(t, err)

	if ok {
		fmt.Printf("Return :%v\n", tc.GetOutput(ovActitvityOutput))
	}

}
