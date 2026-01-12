
# Config file for Product

## nats

Save this in `config/mq/nats-server.conf`

```json
port 4221

authorization {
  user: $NATS_USER
  password: $NATS_PASSWORD
}

jetstream {
  store_dir: /jetstream/data
  max_mem: 1GiB
  max_file: 100GiB
}

websocket {
  port: 8081
  no_tls: true
}
```

## RabbitMQ

Save this in `config/mq/definition.json`

```json

{
  "users": [
    {
      "name": "user",
      "password": "password",
      "tags": "administrator"
    }
  ],
  "vhosts": [
    {
      "name": "/"
    }
  ],
  "permissions": [
    {
      "user": "guest",
      "vhost": "/",
      "configure": ".*",
      "write": ".*",
      "read": ".*"
    }
  ],
  "queues": [
    {
      "name": "my_queue",
      "vhost": "/",
      "durable": true,
      "auto_delete": false,
      "arguments": {}
    }
  ],
  "exchanges": [
    {
      "name": "my_exchange",
      "vhost": "/",
      "type": "direct",
      "durable": true,
      "auto_delete": false,
      "internal": false,
      "arguments": {}
    }
  ],
  "bindings": [
    {
      "source": "my_exchange",
      "vhost": "/",
      "destination": "my_queue",
      "destination_type": "queue",
      "routing_key": "my_routing_key",
      "arguments": {}
    }
  ]
}
```

### Kafka

Save this at `config/mq/jaas.conf`

```json
KafkaServer {
  org.apache.kafka.common.security.plain.PlainLoginModule required
  username="admin"
  password="admin-secret"
  user_admin="admin-secret"
  user_app="app-secret";
};
```

### Valkey

Save this at `/config/cache/valkey.acl`

```json
user default off
user username on >password allcommands allkeys
```

### SeaweedFS

Save this at `config/storage/s3.json`

```json
{
  "identities": [
    {
      "name": "superclient",
      "credentials": [
        {
          "accessKey": "ROOTUSER",
          "secretKey": "CHANGEME123"
        }
      ],
      "actions": [
        "Read",
        "Write",
        "List",
        "Tagging",
        "Admin"
      ]
    }
  ]
}
```
