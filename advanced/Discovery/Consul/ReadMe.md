# CONSUL

## What's this about?
This readme is part of the Packt course on GoLang Microservices and will outline my experiences with Consul following a quick write up on how Consul works, what benefits it brings to the table and how it integrates into a microservice environment.

### What is CONSUL?

### How does CONSUL work?

### What do I use CONSUL for?

### What benefits compared to other xyz software to handle xyz?
How is CONSUL better than other software that solves the same thing?

## Cheat sheet:

### Good commands to know:
`consul` is the package you will use if you wish to interact with a running consul instance. 

#### Add a key-value pair:
```
consul kv put <key> <value>
```

### Good API endpoints: (** = IP and Port)

----

##### Get running services

GET: `**/v1/agent/services`

----

##### Register a service in CONSUL

POST: `**/v1/agent/service/register`

Example body:
```
{
	"ID" : "gin-web-02",
	"Name" : "gin-web",
	"Tags" : [
		"cloud-native-go",
		"v1"
		],
	"Address" : "gin-web-01",
	"Port" : 9090,
	"EnableTagOverride" : false,
	"check" : {
			"id" : "ping",
			"name" : "HTTP API on port 9090",
			"http" : "http://gin-web-01:9090/ping",
			"interval" : "5s",
			"timeout" : "1s"
		}
}
```
-----

##### Health-check(s):
General health of a service:
GET: `**/v1/health/service/name-of-service`

Query Parameters:

* ?passing => returns all services with that name that is passing its health check.
* ?critical => returns all services with that name that is in a critical state.

----

##### Get Key/Value pairs:
Returns any key-value pairs in a running Consul instance:

GET: `**/v1/kv/<optional segments>`

Example: `**/v1/kv/gin-web/message/`

Query Parameters: 

* ?keys => returns all keys in the instance.
* ?raw => returns the value in `raw` form.
----

##### Add Key/Value pair:
Returns `true` if the API call is correct:

PUT: `**/v1/kv/<path>`

<path> => any legal RESTful endpoint where segments are separated by `/`. 

Example: `**/vi/kv/gin-web/message`.

The value would be taken from the body.

----

##### Delete Key/Value pair:
Deletes a Key/Value pair if it exists.

DELETE: `**/v1/kv/<path-to-value>`

----

