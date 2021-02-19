package controllers_test

import (
	"log"
	"testing"

	"github.com/nimblehq/test-gin-templates/test"
	"github.com/spf13/viper"
	"gorm.io/gorm"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestControllers(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Controllers Suite")
}

var _ = BeforeSuite(func() {
	test.SetupTestEnvironment()
})

var _ = AfterSuite(func() {
	gormCnx := viper.Get("database").(*gorm.DB)
	dbCnx, err := gormCnx.DB()

	log.Println("========================= ", err)

	dbCnx.Close()
})
