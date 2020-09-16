# rest-auth-server

An example HTTP REST server with authentication and authorization.

See the blog post on [RESTful authentication in Go](https://www.maragu.dk/blog/restful-authentication-in-go/) for background.

## Usage

Start the server:

```shell script
make start
```

### Demo

Try out the endpoints with [HTTPie](https://httpie.org):

```shell script
$ make demo
http -v --session=./session.json --verify=no https://localhost:8080/check
GET /check HTTP/1.1
Accept: */*
Accept-Encoding: gzip, deflate
Connection: keep-alive
Host: localhost:8080
User-Agent: HTTPie/2.2.0



HTTP/1.1 401 Unauthorized
Content-Length: 27
Content-Type: text/plain; charset=utf-8
Date: Wed, 16 Sep 2020 08:18:18 GMT
X-Content-Type-Options: nosniff

unauthorized, please login

http -v --session=./session.json --form --verify=no post https://localhost:8080/signup name=demo password=1234567890
POST /signup HTTP/1.1
Accept: */*
Accept-Encoding: gzip, deflate
Connection: keep-alive
Content-Length: 29
Content-Type: application/x-www-form-urlencoded; charset=utf-8
Host: localhost:8080
User-Agent: HTTPie/2.2.0

name=demo&password=1234567890

HTTP/1.1 200 OK
Content-Length: 0
Date: Wed, 16 Sep 2020 08:18:18 GMT



http -v --session=./session.json --form --verify=no post https://localhost:8080/login name=demo password=1234567890
POST /login HTTP/1.1
Accept: */*
Accept-Encoding: gzip, deflate
Connection: keep-alive
Content-Length: 29
Content-Type: application/x-www-form-urlencoded; charset=utf-8
Host: localhost:8080
User-Agent: HTTPie/2.2.0

name=demo&password=1234567890

HTTP/1.1 200 OK
Cache-Control: no-cache="Set-Cookie"
Content-Length: 0
Date: Wed, 16 Sep 2020 08:18:19 GMT
Set-Cookie: session=IRygA46_XXlrTaqETg_eNhVoG6bfFKlWQdC9ATAsrPM; Path=/; Expires=Thu, 17 Sep 2020 08:18:20 GMT; Max-Age=86400; HttpOnly; Secure; SameSite=Lax
Vary: Cookie



http -v --session=./session.json --verify=no https://localhost:8080/check
GET /check HTTP/1.1
Accept: */*
Accept-Encoding: gzip, deflate
Connection: keep-alive
Cookie: session=IRygA46_XXlrTaqETg_eNhVoG6bfFKlWQdC9ATAsrPM
Host: localhost:8080
User-Agent: HTTPie/2.2.0



HTTP/1.1 200 OK
Content-Length: 31
Content-Type: text/plain; charset=utf-8
Date: Wed, 16 Sep 2020 08:18:19 GMT

{
    "Name": "demo",
    "Password": null
}

http -v --session=./session.json --form --verify=no post https://localhost:8080/logout
POST /logout HTTP/1.1
Accept: */*
Accept-Encoding: gzip, deflate
Connection: keep-alive
Content-Length: 0
Content-Type: application/x-www-form-urlencoded; charset=utf-8
Cookie: session=IRygA46_XXlrTaqETg_eNhVoG6bfFKlWQdC9ATAsrPM
Host: localhost:8080
User-Agent: HTTPie/2.2.0



HTTP/1.1 200 OK
Cache-Control: no-cache="Set-Cookie"
Content-Length: 0
Date: Wed, 16 Sep 2020 08:18:20 GMT
Set-Cookie: session=; Path=/; Expires=Thu, 01 Jan 1970 00:00:01 GMT; Max-Age=0; HttpOnly; Secure; SameSite=Lax
Vary: Cookie



http -v --session=./session.json --verify=no https://localhost:8080/check
GET /check HTTP/1.1
Accept: */*
Accept-Encoding: gzip, deflate
Connection: keep-alive
Host: localhost:8080
User-Agent: HTTPie/2.2.0



HTTP/1.1 401 Unauthorized
Content-Length: 27
Content-Type: text/plain; charset=utf-8
Date: Wed, 16 Sep 2020 08:18:20 GMT
X-Content-Type-Options: nosniff

unauthorized, please login
```
