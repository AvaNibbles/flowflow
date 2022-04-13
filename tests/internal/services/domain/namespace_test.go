package domain

import (
	"testing"

	"github.com/avanibbles/flowflow/tests"
	"github.com/stretchr/testify/assert"
	"github.com/teris-io/shortid"
)

func TestNamespaceService_GetNamespace_OK(t *testing.T) {
	df := tests.GetTestDependencyFactory()
	ns := df.GetDomain().NewNamespaceService()

	res, err := ns.Get("default")

	assert.NotNil(t, res)
	assert.Nil(t, err)

	assert.Equal(t, "default", res.Name)
}

func TestNamespaceService_CreateNamespace_OK(t *testing.T) {
	df := tests.GetTestDependencyFactory()
	ns := df.GetDomain().NewNamespaceService()

	testNsName := shortid.MustGenerate()
	res, err := ns.Create(testNsName)

	assert.NotNil(t, res)
	assert.Nil(t, err)

	assert.Equal(t, testNsName, res.Name)
}
