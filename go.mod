module github.com/ONSdigital/dp-publishing-dataset-controller

go 1.16

// fixes vulnerabilities in github.com/hashicorp/consul/api@v1.1.0 and github.com/hashicorp/consul/sdk@v0.1.1 dependencies
replace github.com/spf13/cobra => github.com/spf13/cobra v1.4.0

require (
	github.com/ONSdigital/dp-api-clients-go/v2 v2.183.0
	github.com/ONSdigital/dp-healthcheck v1.3.0
	github.com/ONSdigital/dp-net v1.5.0
	github.com/ONSdigital/log.go/v2 v2.2.0
	github.com/golang/mock v1.6.0
	github.com/gorilla/mux v1.8.0
	github.com/kelseyhightower/envconfig v1.4.0
	github.com/pkg/errors v0.9.1
	github.com/smartystreets/goconvey v1.7.2
)
