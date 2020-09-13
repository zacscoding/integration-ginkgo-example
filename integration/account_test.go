package integration

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"integration-ginkgo-example/integration/client"
	"integration-ginkgo-example/internal/account/model"
	"net/http"
	"strconv"
)

var _ = Describe("Account API", func() {
	var (
		accountCli = client.AccountClient{Endpoint: "http://localhost:3000"}
	)

	BeforeEach(func() {
		Expect(db.Delete(&model.Account{}).Error).Should(BeNil())
	})

	Context("GET /v1/accounts", func() {
		When("empty accounts", func() {
			It("returns empty", func() {
				accounts, code, err := accountCli.GetAccounts()
				Expect(err).NotTo(HaveOccurred())
				Expect(code).Should(Equal(http.StatusOK))
				Expect(accounts).Should(BeEmpty())
			})
		})

		When("exist account", func() {
			It("returns accounts", func() {
				var (
					username = "user1"
					email    = "user1@email.com"
					code     int
					err      error
				)
				savedAccount, code, err := accountCli.SaveAccount(username, email)
				Expect(err).NotTo(HaveOccurred())
				Expect(code).Should(Equal(http.StatusOK))

				accounts, code, err := accountCli.GetAccounts()
				Expect(err).NotTo(HaveOccurred())
				Expect(code).Should(Equal(http.StatusOK))
				Expect(accounts).Should(HaveLen(1))
				Expect(accounts[0].ID).Should(Equal(savedAccount.ID))
				Expect(accounts[0].Username).Should(Equal(savedAccount.Username))
				Expect(accounts[0].Email).Should(Equal(savedAccount.Email))
				Expect(accounts[0].CreatedAt).ShouldNot(BeNil())
				Expect(accounts[0].UpdatedAt).ShouldNot(BeNil())
			})
		})
	})

	Context("POST /v1/account", func() {
		BeforeEach(func() {
			Expect(db.Delete(&model.Account{}).Error).Should(BeNil())
		})

		When("request invalid", func() {
			It("returns bad request with email", func() {
				_, code, err := accountCli.SaveAccount("user1", "")
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).Should(ContainSubstring("Email"))
				Expect(code).Should(Equal(http.StatusBadRequest))
			})
			It("returns bad request with username", func() {
				_, code, err := accountCli.SaveAccount("", "user1@email.com")
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).Should(ContainSubstring("Username"))
				Expect(code).Should(Equal(http.StatusBadRequest))
			})
			It("returns duplicate email", func() {
				var (
					username = "user1"
					email    = "user1@email.com"
				)
				_, code, err := accountCli.SaveAccount(username, email)
				Expect(err).NotTo(HaveOccurred())
				Expect(code).To(Equal(http.StatusOK))

				_, code, err = accountCli.SaveAccount(username, email)
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(ContainSubstring("duplicate email"))
				Expect(code).To(Equal(http.StatusBadRequest))
			})
		})

		When("request valid", func() {
			It("returns ok", func() {
				var (
					username = "user1"
					email    = "user1@email.com"
				)
				account, code, err := accountCli.SaveAccount(username, email)
				Expect(err).NotTo(HaveOccurred())
				Expect(code).Should(Equal(http.StatusOK))
				Expect(account.ID).ShouldNot(BeZero())
				Expect(account.Username).Should(Equal(username))
				Expect(account.Email).Should(Equal(email))
				Expect(account.CreatedAt).ShouldNot(BeNil())
				Expect(account.UpdatedAt).ShouldNot(BeNil())
			})
		})
	})

	Context("GET /v1/account/:accountID", func() {
		When("request invalid", func() {
			It("returns bad request", func() {
				_, code, err := accountCli.GetAccount("aa")
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).Should(ContainSubstring("AccountID"))
				Expect(code).To(Equal(http.StatusBadRequest))
			})

			It("returns not found request", func() {
				_, code, err := accountCli.GetAccount("100000")
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).Should(ContainSubstring("not found"))
				Expect(code).To(Equal(http.StatusNotFound))
			})
		})

		When("request valid", func() {
			It("returns a account", func() {
				var (
					username = "user1"
					email    = "user1@email.com"
				)
				account, code, err := accountCli.SaveAccount(username, email)
				Expect(err).NotTo(HaveOccurred())
				Expect(code).Should(Equal(http.StatusOK))

				find, code, err := accountCli.GetAccount(strconv.FormatUint(uint64(account.ID), 10))
				Expect(err).NotTo(HaveOccurred())
				Expect(find.ID).Should(Equal(account.ID))
				Expect(find.Username).Should(Equal(account.Username))
				Expect(find.Email).Should(Equal(account.Email))
				Expect(find.CreatedAt).ShouldNot(BeNil())
				Expect(find.UpdatedAt).ShouldNot(BeNil())
			})
		})
	})

	Context("PUT /v1/account/:accountID", func() {
		When("request invalid", func() {
			It("returns bad request with id", func() {
				_, code, err := accountCli.UpdateAccount("aa", "1")
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).Should(ContainSubstring("AccountID"))
				Expect(code).To(Equal(http.StatusBadRequest))
			})

			It("returns bad request with username", func() {
				_, code, err := accountCli.UpdateAccount("1", "")
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).Should(ContainSubstring("Username"))
				Expect(code).To(Equal(http.StatusBadRequest))
			})

			It("returns not found request", func() {
				_, code, err := accountCli.UpdateAccount("1", "user1")
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).Should(ContainSubstring("not found"))
				Expect(code).To(Equal(http.StatusNotFound))
			})
		})

		When("request valid", func() {
			It("returns updated status", func() {
				var (
					username    = "user1"
					email       = "user1@email.com"
					updatedName = "updatedUser1"
				)
				account, code, err := accountCli.SaveAccount(username, email)
				Expect(err).NotTo(HaveOccurred())
				Expect(code).Should(Equal(http.StatusOK))

				status, code, err := accountCli.UpdateAccount(strconv.FormatUint(uint64(account.ID), 10), updatedName)
				Expect(err).NotTo(HaveOccurred())
				Expect(code).Should(Equal(http.StatusOK))
				Expect(status.Status).Should(Equal("updated"))

				find, code, err := accountCli.GetAccount(strconv.FormatUint(uint64(account.ID), 10))
				Expect(err).NotTo(HaveOccurred())
				Expect(code).Should(Equal(http.StatusOK))
				Expect(find.Username).Should(Equal(updatedName))
			})
		})
	})

	Context("DELETE /v1/account/:accountID", func() {
		When("request invalid", func() {
			It("returns bad request with id", func() {
				_, code, err := accountCli.DeleteAccount("aa")
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).Should(ContainSubstring("AccountID"))
				Expect(code).To(Equal(http.StatusBadRequest))
			})

			It("returns not found request", func() {
				_, code, err := accountCli.DeleteAccount("1")
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).Should(ContainSubstring("not found"))
				Expect(code).To(Equal(http.StatusNotFound))
			})
		})

		When("request valid", func() {
			It("returns updated status", func() {
				var (
					username = "user1"
					email    = "user1@email.com"
				)
				account, code, err := accountCli.SaveAccount(username, email)
				Expect(err).NotTo(HaveOccurred())
				Expect(code).Should(Equal(http.StatusOK))

				status, code, err := accountCli.DeleteAccount(strconv.FormatUint(uint64(account.ID), 10))
				Expect(err).NotTo(HaveOccurred())
				Expect(code).Should(Equal(http.StatusOK))
				Expect(status.Status).Should(Equal("deleted"))

				_, code, _ = accountCli.GetAccount(strconv.FormatUint(uint64(account.ID), 10))
				Expect(code).Should(Equal(http.StatusNotFound))
			})
		})
	})
})
