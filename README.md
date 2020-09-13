# Integration test with ginkgo :)  
This project is an example for integration test with ginkgo.  
See integration [test code](./integration)(WORKING!!)  

Assume that we are serving account crud api and integration test scenario like below  

> ### APIs  

- `GET /v1/accounts` : get all accounts
- `POST /v1/account` : create a new account
- `GET /v1/account/:id` : get an account
- `PUT /v1/account/:id` : update an account
- `DELETE /v1/account/:id` : delete an account  

> ### Scenario  
; TBD  

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

> ### Run integration test (TBD)  

```bash
$ make integration
```

---  

## Reference  

- https://github.com/onsi/ginkgo
- https://semaphoreci.com/community/tutorials/getting-started-with-bdd-in-go-using-ginkgo