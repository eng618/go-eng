/*
package main is an example of how to write a web service using gin.

This example follows the tutorial: https://go.dev/doc/tutorial/web-service-gin

To test this code locally, you must clone the repository, `cd` into it, and run the following:

```shell
go run ./examples/web-service-gin
```

The above command will run the server locally and you can interact easily via `curl` or other methods.

Available endpoints:

1. GET /albums
  - Responds with the list of all albums as JSON.

2. GET /albums/:id
  - Responds with the album whose ID value matches the id parameter sent by the client.

3. POST /albums
  - Adds an album from JSON received in the request body.
*/
package main
