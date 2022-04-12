package tests

import (
	"sync"

	"github.com/avanibbles/flowflow/internal"
	"github.com/avanibbles/flowflow/internal/flowflow"
	"github.com/avanibbles/flowflow/internal/services"
	"github.com/avanibbles/flowflow/pkg/config"

	_ "github.com/joho/godotenv/autoload"
)

var dependencyFactory services.DependencyFactory
var dfLck sync.Mutex
var dfOnce sync.Once

func GetTestDependencyFactory() services.DependencyFactory {
	dfOnce.Do(func() {
		dfLck.Lock()
		defer dfLck.Unlock()

		df, err := buildTestDependencyFactory()
		if err != nil {
			panic(err)
		}

		dependencyFactory = df
	})

	return dependencyFactory
}

func buildTestDependencyFactory() (services.DependencyFactory, error) {
	if err := flowflow.SetupTestingConfig(); err != nil {
		return nil, err
	}

	cfg, err := config.LoadConfigFromViper()
	if err != nil {
		return nil, err
	}

	dc := &services.DependencyConfig{
		Logger: internal.GetLogger(),
		Config: cfg,
		Wg:     &sync.WaitGroup{},
	}

	df, err := services.NewDependencyFactory(dc)
	if err != nil {
		return nil, err
	}

	if err = initDb(df); err != nil {
		return nil, err
	}

	return df, nil
}

func initDb(df services.DependencyFactory) error {
	ms := df.GetDomain().NewMaintenanceService()

	if err := ms.PreServiceStart(); err != nil {
		return err
	}

	if err := dbDataPrep(df); err != nil {
		return err
	}

	return nil
}
