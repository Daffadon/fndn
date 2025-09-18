package config_template

const YamlConfigMessageTemplate string = `
# this is the single source of truth for configuration.
# to make it easier to config and not hardcoded, always put the variable here
# and call it using viper in your app

# for production, copy this file and change the name to config.yaml
# for test, copy this file and change the name to config.test.yaml
`

const AppYamlConfigTemplate string = `
app:
  # name: "your_app_name"
  http:
    port: 8080
# this is an example. uncomment if your app connect to other services via grpc
  # grpc:
    # service:
      # service_a: service_a:50051
      # service_b: service_b:50051
`


