/*
Implementation of RabbitMq Management HTTP Api in Go. Alpha status.

Currently supported api calls are:

    GET     /api/exchanges/
    GET     /api/exchanges/vhost/name
    PUT     /api/exchanges/vhost/name
    DELETE  /api/exchanges/vhost/name

    GET     /api/vhosts
    GET     /api/vhost/name
    PUT     /api/vhost/name
    DELETE  /api/vhost/name
    GET     /api/vhost/name/permissions

    GET     /api/users
    GET     /api/users/name
    PUT     /api/users/name
    DELETE  /api/users/name
    GET     /api/users/name/permissions

    GET     /api/permissions
    GET     /api/permissions/vhost/user
    PUT     /api/permissions/vhost/user
    DELETE  /api/permissions/vhost/user

Example code:

	r := rabbitapi.Auth("guest", "guest", "http://localhost:15672")
	vhosts, err := r.GetVhosts()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("vhosts:", vhosts)
	}

	err = r.CreateExchange("/", "rabbitapi", "topic", false, false, false, nil)
	if err != nil {
		fmt.Println(err)
	}

	exchange, err := r.GetExchange("/", "rabbitapi")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(exchange) // exchange.Type is 'topic'
*/
package rabbitapi
