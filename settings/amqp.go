package settings

import "fmt"

type ConnectionSettings struct {
	host string
	port string
	vhost string
	user string
	password string
}

func (cs *ConnectionSettings) GetConnectionUrl() ConnectionUrl {
	if cs.vhost != "/" {
		cs.vhost = "/" + cs.vhost
	}

	return ConnectionUrl(fmt.Sprintf("amqp://%s:%s@%s:%s%s",
		cs.user,
		cs.password,
		cs.host,
		cs.port,
		cs.vhost))
}

type ConnectionUrl string


var AmqpConnectionSettings = &ConnectionSettings{
	host:     getEnvironment("AMQP_HOST", "localhost"),
	port:     getEnvironment("AMQP_PORT", "5672"),
	vhost:    getEnvironment("AMQP_VHOST", "/"),
	user:     getEnvironment("AMQP_USER", "guest"),
	password: getEnvironment("AMQP_PASSWORD", "guest"),
}

var ConsumerQueue = getEnvironment("CONSUMER_QUEUE", "request.v1")
var PublishQueue = getEnvironment("PUBLISH_QUEUE", "response.v1")
var ErrorQueue = getEnvironment("ERROR_QUEUE", "error.v1")

