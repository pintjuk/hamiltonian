
# Routemaster

API that provides a list of the closest pickup spots to the customer

![](https://t3.ftcdn.net/jpg/00/82/84/72/360_F_82847279_fW7TCAujmT5B1G3HwSK7uPzrb0dgXGo2.jpg)



## Demo
Demo server is running on: https://routemaster-gfpakxausa-lz.a.run.app


try it out with
https://routemaster-gfpakxausa-lz.a.run.app/routes?src=12.388860,52.517037&dst=11.397634,52.529407&dst=13.428555,52.523219

# Project structure

- /routes endpoint is implemented in [src/http_resources/routes.go](src/http_resources/routes.go)
- OSRM client is in [src/integrations/osrm/osrm.go](src/integrations/osrm/osrm.go)
- [api/api.go](api/api.go) contains DTOs for routemaster
- [infra/main.tf] contains infra as code
- [src/route/route.go](src/route/route.go) containes two functions TestGetClosestRouteWithDurationAndDistance and TestGetClosestRouteWithDurationAndDistanceAsinc, that implement teh core logic of the application the only difference is that one is synchronous and te other is not. both are tested in [src/route/route_test.go](src/route/route_test.go). curently teh synchronous one is used since its simpler and there was no big advantage with the asynchronous when running against the demo osrm server
## How to run locally

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
### Release
push a semvar tag:
1) `git tag -a v0.0.2`
2) `git push --tags`

### Deploy
if you released version v0.0.2 and now wish to deploy it:
1) Ask Daniil to grant your user accont the needed role to deploy infra
2) `cd infra`
3) `terraform init` (if you have not done so previously)
3) run `terraform apply -var service_version=v0.0.2`
4) make sure that you understand the changes that will be performed by terraform
5) run `terraform apply -var service_version=v0.0.2`


