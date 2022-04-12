package domain

import (
	"testing"

	"github.com/avanibbles/flowflow/tests"
	"github.com/stretchr/testify/assert"
)

func TestMLModelService_Get_Exists_Ok(t *testing.T) {
	df := tests.GetTestDependencyFactory()
	svc := df.GetDomain().NewMLModelService()

	model, err := svc.Get("default", "test")

	assert.NotNil(t, model)
	assert.Nil(t, err)
}
