# Integration test with ginkgo :)  
This project is an example for integration test with ginkgo.  
See integration [test code](./integration)  

Assume that we are serving account crud api and integration test scenario like below  

> ### APIs  

**Note**: You can test simply by using [request.http](./tools/http/request.http)  

- `GET /v1/accounts` : get all accounts
- `POST /v1/account` : create a new account
- `GET /v1/account/:id` : get an account
- `PUT /v1/account/:id` : update an account
- `DELETE /v1/account/:id` : delete an account  

> ### Scenario  
This integration test focus on request endpoint i.e HTTP Method + URI such as `GET /v1/accounts`

- Describe("Account API")
  - Context("GET /v1/accounts")
    - When("empty accounts")
      - It("returns empty")
    - When("exist account")
      - It("returns accounts)
  - Context("POST /v1/account")
    - When("request invalid")
      - It("returns bad request with email")
      - It("returns bad request with username")
      - It("returns duplicate email")
    - When("request valid")
      - It("returns ok")
  - Context("GET /v1/account/:accountID")
    - When("request invalid")
      - It("returns bad request")
      - It("returns not found request")
    - When("request valid")
      - It("returns a account")
  - Context("PUT /v1/account/:accountID")
    - When("request invalid")
      - It("returns bad request with id")
      - It("returns bad request with username")
      - It("returns not found request")
    - When("request valid")
      - It("returns updated status")
  - Context("DELETE /v1/account/:accountID")
    - When("request invalid")
      - It("returns bad request with id")
      - It("returns not found request")
    - When("request valid")
      - It("returns updated status")

---  

## Getting started  

> ### Build a docker image

```bash
$ make build
```

> ### Run docker-compose (api server + database)  

```bash
// should down with -v flag to remove docker volume i.e docker-compose down -v or make compose-down
$ make compose
```  

> ### Run integration test

```bash
$ make integration
go test ./integration -v
=== RUN   TestIntegration
Running Suite: Integration Suite
================================
Random Seed: 1600011660
Will run 16 of 16 specs

••••••••••••••••
Ran 16 of 16 Specs in 0.226 seconds
SUCCESS! -- 16 Passed | 0 Failed | 0 Pending | 0 Skipped
--- PASS: TestIntegration (0.23s)
PASS
ok      integration-ginkgo-example/integration  0.878s
```

---  

## Reference  

- https://github.com/onsi/ginkgo
- https://semaphoreci.com/community/tutorials/getting-started-with-bdd-in-go-using-ginkgo