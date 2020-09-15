# rest-auth-server

An example HTTP REST server with authentication and authorization.

See the blog post on [RESTful authentication in Go](https://www.maragu.dk/blog/restful-authentication-in-go/) for background.

## Usage

Start the server:

```shell script
make start
```

Try out the endpoints with [HTTPie](https://httpie.org):

### Unauthenticated

```shell script
$ http -v --session=./session.json --verify=no https://localhost:8080/check
GET /check HTTP/1.1
Accept: */*
Accept-Encoding: gzip, deflate
Connection: keep-alive
Host: localhost:8080
User-Agent: HTTPie/2.2.0



HTTP/1.1 401 Unauthorized
Content-Length: 27
Content-Type: text/plain; charset=utf-8
Date: Tue, 15 Sep 2020 12:23:07 GMT
X-Content-Type-Options: nosniff

unauthorized, please login
```

### Sign up

```shell script
$ http -v --session=./session.json --form --verify=no post https://localhost:8080/signup name=markus password=1234567890
POST /signup HTTP/1.1
Accept: */*
Accept-Encoding: gzip, deflate
Connection: keep-alive
Content-Length: 31
Content-Type: application/x-www-form-urlencoded; charset=utf-8
Host: localhost:8080
User-Agent: HTTPie/2.2.0

name=markus&password=1234567890

HTTP/1.1 200 OK
Content-Length: 0
Date: Tue, 15 Sep 2020 12:24:48 GMT
```

### Log in

```shell script
$ http -v --session=./session.json --form --verify=no post https://localhost:8080/login name=markus password=1234567890
POST /login HTTP/1.1
Accept: */*
Accept-Encoding: gzip, deflate
Connection: keep-alive
Content-Length: 31
Content-Type: application/x-www-form-urlencoded; charset=utf-8
Host: localhost:8080
User-Agent: HTTPie/2.2.0

name=markus&password=1234567890

HTTP/1.1 200 OK
Cache-Control: no-cache="Set-Cookie"
Content-Length: 0
Date: Tue, 15 Sep 2020 12:25:13 GMT
Set-Cookie: session=k0R1eMO7bFM-FCWXb4dhMTjpV1aVAmdT8_cbtSyt5Kc; Path=/; Expires=Wed, 16 Sep 2020 12:25:14 GMT; Max-Age=86400; HttpOnly; SameSite=Lax
Vary: Cookie
```

### Authenticated

```shell script
$ http -v --session=./session.json --verify=no https://localhost:8080/check
GET /check HTTP/1.1
Accept: */*
Accept-Encoding: gzip, deflate
Connection: keep-alive
Cookie: session=k0R1eMO7bFM-FCWXb4dhMTjpV1aVAmdT8_cbtSyt5Kc
Host: localhost:8080
User-Agent: HTTPie/2.2.0



HTTP/1.1 200 OK
Content-Length: 33
Content-Type: text/plain; charset=utf-8
Date: Tue, 15 Sep 2020 12:25:40 GMT

{
    "Name": "markus",
    "Password": null
}
```

### Log out

```shell script
$ http -v --session=./session.json --form --verify=no post https://localhost:8080/logout
POST /logout HTTP/1.1
Accept: */*
Accept-Encoding: gzip, deflate
Connection: keep-alive
Content-Length: 0
Content-Type: application/x-www-form-urlencoded; charset=utf-8
Cookie: session=k0R1eMO7bFM-FCWXb4dhMTjpV1aVAmdT8_cbtSyt5Kc
Host: localhost:8080
User-Agent: HTTPie/2.2.0



HTTP/1.1 200 OK
Cache-Control: no-cache="Set-Cookie"
Content-Length: 0
Date: Tue, 15 Sep 2020 12:27:00 GMT
Set-Cookie: session=; Path=/; Expires=Thu, 01 Jan 1970 00:00:01 GMT; Max-Age=0; HttpOnly; SameSite=Lax
Vary: Cookie
```
