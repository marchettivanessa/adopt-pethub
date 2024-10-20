package database

import (
	"adopt-pethub/backend/config"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestDatabase(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Database Suite")
}

var _ = Describe("Adopt Pethub Database", func() {
	Context("build connection string", func() {
		It("builds the data source name according to config", func() {
			c := config.DatabaseConfig{
				Name:           "adopt_pethub",
				Host:           "any_host",
				Password:       "any_password",
				Username:       "any_username",
				Port:           1223,
				ConnectTimeout: 20,
				Schema:         "adopt_pethub",
			}

			exercise := buildConnectionString(c)
			expectedString := "host=any_host port=1223 user=any_username password=any_password dbname=adopt_pethub connect_timeout=20 search_path=public,adopt_pethub sslmode=disable"
			Expect(exercise).To(Equal(expectedString))
		})
	})
})
