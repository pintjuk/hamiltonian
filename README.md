
# Routemaster

API that provides a list of the closest pickup spots to the customer

![](https://t3.ftcdn.net/jpg/00/82/84/72/360_F_82847279_fW7TCAujmT5B1G3HwSK7uPzrb0dgXGo2.jpg)



## Demo
demo server is running on ...
to try it out run
`curl`

## How to run localy

### With go
0) [install go](https://go.dev/doc/install)
1) [install git](https://git-scm.com/book/en/v2/Getting-Started-Installing-Git)
2) `git clone git@github.com:pintjuk/routemaster.git`
3) `cd routemaster`
4) `make run`
5) `curl `
### how to run tests
make sure you run steps 0-3 above then `make test` to generate coverage report run `make cover`

## How to run in docker
if you don't have golang installed but have docker, or if it did not work you may try running the server in docker:
1) make sure you install docker
2) `make docker-build`
3) `make docker run`

## Operational tasks
### View Logs
go to link
### View Dshboards
go to link
### Deploy
NOTE to self: we should idealy deploy on merges to master, but this is such a small project right now that this is fine

1) Ask Daniil to grant your user accont the needed role to deploy infra
2) Make changes to the terraform files in 'infra/', for example bump the docker tag
3) run `terraform diff`
4) make sure that you understand the changes that will be performed by terraform
5) run `terraform apply`

# TODO:
1) [ ] debug
2) [ ] make sure the unit test makes sense
3) [ ] docker file
4) [ ] terraform
5) [ ] finish readme
6) [ ] review code again to make sure it is readable
