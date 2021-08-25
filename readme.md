# Bookings and Reservations

I cloned and followed by the repository for [Building Modern Web Applications with Go](https://www.udemy.com/course/building-modern-web-applications-with-go/?referralCode=0415FB906223F10C6800).


- Built in Go version 1.16
- Uses the [chi router](github.com/go-chi/chi)
- Uses [alex edwards scs session management](github.com/alexedwards/scs)
- Uses [nosurf](github.com/justinas/nosurf)

## Start the Web Server

### Run the web server
```
$ go run cmd/web/*.go
```

## Start Test

### Run tests
```
$ cd cmd/web
$ go test -v
```

### Check coverage
```
$ cd cmd/web
$ go test -cover
```

### Check coverage
```
$ cd cmd/web
$ go test -coverprofile=coverage.out && go tool cover -html=coverage.out
```

After starting server, open the browser, hit the `127.0.0.1:8080` in URL