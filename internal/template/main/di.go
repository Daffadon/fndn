package main_template

const DITemplate string = `
package di

func BuildContainer() *dig.Container {
	container := dig.New()

	// here is your dependency injection using dig with order matter,
	// but for the very first time, i will give you a default one.
	// you can change this anytime

	// logger
	if err := container.Provide(logger.NewLogger); err != nil {
		panic("Failed to provide logger: " + err.Error())
	}
	// minio connection
	if err := container.Provide(storage.NewMinioConnection); err != nil {
		panic("Failed to provide minio connection: " + err.Error())
	}
	// nats client connection
	if err := container.Provide(mq.NewNatsConnection); err != nil {
		panic("Failed to provide nats connection: " + err.Error())
	}
	// postgresql pool connection
	if err := container.Provide(storage.NewSQLConn); err != nil {
		panic("Failed to provide postgresql pool connection: " + err.Error())
	}
	// jetstream connection
	if err := container.Provide(jetstream.New); err != nil {
		panic("Failed to provide jetstream instance: " + err.Error())
	}
	// redis connection
	if err := container.Provide(cache.NewRedisConnection); err != nil {
		panic("Failed to provide redis connection: " + err.Error())
	}

	// you can add your own handler, service, or repository here 
	// and invoke in the /cmd/server/http_server.go to register
	// the route to http server

	// http server (gin)
	if err := container.Provide(router.NewHTTP); err != nil {
		panic("Failed to provide gin http server: " + err.Error())
	}
	return container
}
`
