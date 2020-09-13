package integration_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"integration-ginkgo-example/integration/client"
)

// TODO : impl

var _ = Describe("Account", func() {
	var (
		accountCli   = client.AccountClient{Endpoint: "http://localhost:3000"}
		username     = "user1"
		email        = "user1@email.com"
		savedAccount *client.Account
		err          error
	)

	Describe("GET /accounts", func() {
		Context("when not exist accounts", func() {
			It("returns empty accounts", func() {
				accounts, err := accountCli.GetAccounts()
				Expect(err).Should(BeNil())
				Expect(accounts).Should(BeEmpty())
			})
		})

		Context("when exist accounts", func() {
			savedAccount, err = accountCli.SaveAccount(username, email)
			Expect(err).Should(BeNil())

			It("returns accounts", func() {
				accounts, err := accountCli.GetAccounts()
				Expect(err).Should(BeNil())
				Expect(accounts).Should(HaveLen(1))
				Expect(accounts[0]).Should(Equal(savedAccount))
			})
		})
	})

	//Context("initially", func() {
	//	It("should be empty", func() {
	//		accounts, err := accountCli.GetAccounts()
	//		Expect(err).Should(BeNil())
	//		Expect(accounts).Should(BeEmpty())
	//	})
	//})
	//
	//Context("create a new account", func() {
	//
	//})
	//
	//Context("create a new account with invalid request", func() {
	//	It("should be fail", func() {
	//		_, err := accountCli.SaveAccount("", email)
	//		Expect(err).Should(HaveOccurred())
	//		Expect(err.Error()).Should(ContainSubstring("Username"))
	//	})
	//	It("should be fail", func() {
	//		_, err := accountCli.SaveAccount(username, "")
	//		Expect(err).Should(HaveOccurred())
	//		Expect(err.Error()).Should(ContainSubstring("Email"))
	//	})
	//})
	//
	//Context("create a new account", func() {
	//	//It("should be success", func() {
	//	//	account, err := accountCli.SaveAccount(username, email)
	//	//	Expect(err).Should(BeNil())
	//	//	Expect(account.Username).Should(BeEquivalentTo(username))
	//	//	Expect(account.Email).Should(BeEquivalentTo(email))
	//	//	//savedAccount = account
	//	//})
	//})
})
