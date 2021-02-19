package helpers_test

import (
	"testing"

	"github.com/nimblehq/test-gin-templates/test"
	"github.com/spf13/viper"
	"gorm.io/gorm"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestHelpers(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Helpers Suite")
}

var _ = BeforeSuite(func() {
	test.SetupTestEnvironment()
})

var _ = AfterSuite(func() {
	gormCnx := viper.Get("database").(*gorm.DB)
	dbCnx, err := gormCnx.DB()
	if err != nil {
		Fail("Failed to get database instance: " + err.Error())
	}

	dbCnx.Close()
})
