package main_template

const DITemplate string = `
package di


import (
	"go.uber.org/dig"
	"github.com/nats-io/nats.go/jetstream"
)

func BuildContainer() *dig.Container {
	container := dig.New()

	// here is your dependency injection using dig with order matter,
	// but for the very first time, i will give you a default one.
	// you can change this anytime

	// logger
	if err := container.Provide(logger.NewLogger); err != nil {
		panic("Failed to provide logger: " + err.Error())
	}
	// object storage connection
	if err := container.Provide(storage.{{.OSConnection}}); err != nil {
		panic("Failed to provide object storage connection: " + err.Error())
	}

	{{.MQInit}}

	// db connection
	if err := container.Provide(storage.{{.DBConnection}}); err != nil {
		panic("Failed to provide db connection: " + err.Error())
	}
	//  connection
	if err := container.Provide(cache.{{.CacheConnection}}); err != nil {
		panic("Failed to provide cache connection: " + err.Error())
	}

	// you can add your own handler, service, repository,infra, or even 
	// your own defined config here and invoke in the /cmd/server/http_server.go 
	
	// infra
	if err := container.Provide(cache_infra.{{.CacheInfra}}); err != nil {
		panic("Failed to provide cache infra: " + err.Error())
	}	
	if err := container.Provide(mq_infra.{{.MQInfra}}); err != nil {
		panic("Failed to provide MQ infra: " + err.Error())
	}	
	if err := container.Provide(storage_infra.{{.OSInfra}}); err != nil {
		panic("Failed to provide object storage infra: " + err.Error())
	}	
	if err := container.Provide(storage_infra.NewQuerier); err != nil {
		panic("Failed to provide querier infra: " + err.Error())
	}

	// repo
	if err := container.Provide(repository.NewTodoRepository); err != nil {
		panic("Failed to provide todo repository: " + err.Error())
	}
	// service
	if err := container.Provide(service.NewTodoService); err != nil {
		panic("Failed to provide todo service: " + err.Error())
	}
	// handler
	if err := container.Provide(handler.NewTodoHandler); err != nil {
		panic("Failed to provide todo handler: " + err.Error())
	}

	// http server
	if err := container.Provide(router.NewHTTP); err != nil {
		panic("Failed to provide http server: " + err.Error())
	}
	return container
}
`
