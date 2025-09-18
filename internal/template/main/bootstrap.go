package main_template

const BootStrapTemplate = `
package bootstrap

func Run() *dig.Container {
	env.Load()
	container := di.BuildContainer()
	return container
}
`
