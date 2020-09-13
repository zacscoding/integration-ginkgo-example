package integration

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var (
	db         *gorm.DB
)

var _ = BeforeSuite(func() {
	var err error
	db, err = gorm.Open("mysql", "root:password@(localhost:3306)/test_db?parseTime=True&charset=utf8mb4&loc=Local")
	if err != nil {
		Fail(err.Error())
	}
})

var _ = AfterSuite(func() {
	if db != nil {
		db.Close()
	}
})

func TestIntegration(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Integration Suite")
}
