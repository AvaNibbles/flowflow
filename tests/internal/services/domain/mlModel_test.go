package domain

import (
	"testing"

	"github.com/avanibbles/flowflow/internal/util"

	"github.com/teris-io/shortid"

	"github.com/avanibbles/flowflow/tests"
	"github.com/stretchr/testify/assert"
)

func TestMLModelService_Get_Exists_Ok(t *testing.T) {
	df := tests.GetTestDependencyFactory()
	svc := df.GetDomain().NewMLModelService()

	model, err := svc.Get("default", "test")

	assert.NotNil(t, model)
	assert.Nil(t, err)

	assert.Equal(t, model.Name, "test")
	assert.Equal(t, model.Namespace.Name, "default")
}

func TestMLModelService_Create_Ok(t *testing.T) {
	df := tests.GetTestDependencyFactory()
	svc := df.GetDomain().NewMLModelService()

	newModelName := shortid.MustGenerate()
	model, err := svc.Create("default", newModelName)

	assert.NotNil(t, model)
	assert.Nil(t, err)

	assert.Equal(t, model.Name, newModelName)
	assert.GreaterOrEqual(t, model.ID, uint(0))
}

func TestMLModelService_Get_NotFound_Err(t *testing.T) {
	df := tests.GetTestDependencyFactory()
	svc := df.GetDomain().NewMLModelService()

	model, err := svc.Get("default", "not-found")

	assert.Nil(t, model)
	assert.NotNil(t, err)

	if se, ok := err.(util.StatusError); ok { //nolint - we want to check for specific toplevel err here
		assert.Equal(t, 404, se.Code())
	} else {
		assert.Fail(t, "error was not of the correct type")
	}
}
