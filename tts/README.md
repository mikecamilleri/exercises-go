# Michael Camilleri - Software Engineering Coding challenge

This is my solution to the _Software Engineering Coding Challenge_. I have tried to reasonably replicate something that might be used in production, but have taken certain shortcuts which I note here and in the code comments. 

## Tools Used and Rationale

### Language

I chose to implement this program in Go because it is the language that I am most comfortable with. Go is a good choice for a project like this because it is a general purpose programming language with a focus on networking, concurrency, and readability.

### Database

Go provides a consistent [interface](https://golang.org/pkg/database/sql/) for interacting with SQL databases via [drivers](https://github.com/golang/go/wiki/SQLDrivers). Because of this, various relational database management systems are (almost) drop in replacements for each other, at least for a project such as this where RDBMS-specific features and performance are not concerns. Instead of using Postgres, I'm using SQLite. SQLite is an excellent choice for a project like this because its databases exist as single files stored on disk or in memory. This makes configuration and testing easier than it would be if I used Postgres, but still allows me to demonstrate my ability to work with a relational database. 

### HTTP API

I've used the [`gorilla/mux`](https://github.com/gorilla/mux) package the help build the HTTP API. I could have done it with the built in HTTP package, but I wanted to take this opportunity to become more familiar with the Gorilla Toolkit. 

### Logging

Logging is done using Go's built in `log` package. The features of the built in `log` package are pretty minimal (e.g. no levels), but adequate for this project. In an actual production environment, a package that builds upon the built in logger such as [Logrus](https://github.com/sirupsen/logrus) or [Glog](https://github.com/golang/glog) would almost certainly be used. 

### Testing

In order to save time, I've implemented only a couple of unit tests. In actual production code, unit tests for each function should cover successful conditions including edge cases, and all error conditions. 100% code coverage is a useful goal. Additional tests such as integration and regression tests will likely also be required in a production environment. 

## Program Usage

### HTTP Interface

The interface to this program is REST-like and speaks JSON. I use [HTTPie](https://httpie.org) in the following examples. HTTPie uses JSON by default and builds it from the arguments passed.

New shoes are added to the database using POST requests to the `/shoes` endpoint. 

```
$ http POST http://localhost:8080/api/v0/shoes brand="Nike" name="Kyrie 5 Spongebob Squarepants"
HTTP/1.1 201 Created
Content-Length: 0
Date: Tue, 15 Oct 2019 21:39:07 GMT

$ http POST http://localhost:8080/api/v0/shoes brand="Timberland" name="World Hiker Front Country Boot Supreme Orange"
HTTP/1.1 201 Created
Content-Length: 0
Date: Tue, 15 Oct 2019 21:39:15 GMT

```

A list of shoes in the database can be retrieved by making a GET request to the `/shoes` endpoint. I am aware that the returned JSON isn't valid because the outer most element is an array (`[]`), not a object (`{}`). In a production system this design problem should be corrected.

```
$ http GET http://localhost:8080/api/v0/shoes
HTTP/1.1 200 OK
Content-Length: 150
Content-Type: application/json
Date: Tue, 15 Oct 2019 21:41:27 GMT

[
    {
        "brand": "Nike",
        "id": 1,
        "name": "Kyrie 5 Spongebob Squarepants"
    },
    {
        "brand": "Timberland",
        "id": 2,
        "name": "World Hiker Front Country Boot Supreme Orange"
    }
]

```

A new TTS value can be added for a shoe by making a POST request to `/shoes/<id>/tts`:

```
$ http POST http://localhost:8080/api/v0/shoes/1/tts value:=1
HTTP/1.1 201 Created
Content-Length: 0
Date: Tue, 15 Oct 2019 21:46:35 GMT

$ http POST http://localhost:8080/api/v0/shoes/1/tts value:=2
HTTP/1.1 201 Created
Content-Length: 0
Date: Tue, 15 Oct 2019 21:46:39 GMT

$ http POST http://localhost:8080/api/v0/shoes/1/tts value:=2 
HTTP/1.1 201 Created
Content-Length: 0
Date: Tue, 15 Oct 2019 21:46:42 GMT

$ http POST http://localhost:8080/api/v0/shoes/1/tts value:=5
HTTP/1.1 201 Created
Content-Length: 0
Date: Tue, 15 Oct 2019 21:46:49 GMT

```

The average of TTS values for a shoe can be retrieved by making a get request to  `/shoes/<id>/tts`:

```
$ http GET http://localhost:8080/api/v0/shoes/1/tts                                                                 
HTTP/1.1 200 OK
Content-Length: 25
Content-Type: application/json
Date: Tue, 15 Oct 2019 21:51:45 GMT

{
    "shoeID": 1,
    "value": 2.5
}

```

### Getting This Running.

## Installing Prerequisites

If you are running MacOS Catalina, the `tts` binary included in this archive might work, so skip down to _Running the Server_ and give it a try. This program was built and tested on a Mac, but should work on any Unix-like OS and might work on Windows. 

In order to build this program on your machine, you will need to have Go installed. 

```
mike@Darwin-D tts % go version
go version go1.12.7 darwin/amd64
```

I built and tested this program with the version of Go listed above. If your version is significantly older, or you don't have Go installed, follow the instructions at https://golang.org/doc/install. On the Mac, I prefer to install using Homebrew. 

Once Go is installed, run the following commands to get some requirements:

```
CGO_ENABLED=1 go get github.com/mattn/go-sqlite3
go get -u github.com/gorilla/mux
```

`mattn/go-sqlite3` requires CGO be enabled since it uses C bindings. It also requires that GCC be installed. 

## Running the Tests

The test suite may be run with `go test -v`: 

```
mike@Darwin-D tts % go test -v 
=== RUN   TestNewShoeAndGetShoes
--- PASS: TestNewShoeAndGetShoes (0.00s)
=== RUN   TestCreateShoeAndReadShoes
--- PASS: TestCreateShoeAndReadShoes (0.00s)
PASS
ok  	github.com/mikecamilleri/tts	0.015s
```

## Running the Server

The included `config.json` file contains the following defaults. I personally don't like configuration files written in JSON (something like YAML is better for config), but JSON was convenient for this simple project since it's used elsewhere:

```
{
    "dbPath": ":memory:",
    "apiBasePath": "/api/v0",
    "apiPort": "8080"
}
```

`debPath` is set to use an in memory database by default which of course will disappear when the server is exited. Setting a path such as `./test.db` will allow the database to be persistent.

Version numbers are nice to have in an API's path. In this case, I use `v0` in `apiBasePath` to represent an unstable API currently under development.

To run the server, first build it with `go build`, then run it as follows. Logs are output on `stderr`, so let's redirect them to `stdout`:

```
mike@Darwin-D tts % ./tts 2>&1
2019/10/15 15:22:57 INFO: initializing database ...
2019/10/15 15:22:57 INFO: database successfully initialized
2019/10/15 15:22:57 INFO: serving http ...

```

The server is now running.
