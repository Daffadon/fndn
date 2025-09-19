package main_template

const BootStrapTemplate = `
package bootstrap

import "go.uber.org/dig"

func Run() *dig.Container {
	env.Load()
	container := di.BuildContainer()
	return container
}
`
