# Bookings and Reservations

I cloned and followed by the repository for [Building Modern Web Applications with Go](https://www.udemy.com/course/building-modern-web-applications-with-go/?referralCode=0415FB906223F10C6800).


- Built in Go version 1.16
- Uses the [chi router](github.com/go-chi/chi)
- Uses [alex edwards scs session management](github.com/alexedwards/scs)
- Uses [nosurf](github.com/justinas/nosurf)
- Uses [go-simple-mail](github.com/xhit/go-simple-mail)
    - [MailHog](https://github.com/mailhog/MailHog) - For testing at localhost
        - `brew services start mailhog`
    - [Foundation](https://get.foundation/emails.html) - Email templates
- Uses [RoyalUI-Free-Bootstrap-Admin-Template](https://github.com/BootstrapDash/RoyalUI-Free-Bootstrap-Admin-Template) - Admin dashboard template

## Start the Web Server

### Run the web server
```
$ source run.sh
```

## Start Test

### Run tests
```
$ cd cmd/web
$ go test -v

$ go test -v ./...
```

### Check coverage
```
$ cd cmd/web
$ go test -cover
$ go test -coverprofile=coverage.out && go tool cover -html=coverage.out
```

### Check coverage
```
$ cd cmd/web
$ go test -coverprofile=coverage.out && go tool cover -html=coverage.out
```

After starting server, open the browser, hit the `127.0.0.1:8080` in URL

## Database

### PostgreSQL

PostgreSQL is the DB be used in this course.

### Run DB through Docker

```bash
# Execute docker image
$ docker pull postgres:14
$ docker run --name postgres-go-course -p 5432:5432 -e POSTGRES_PASSWORD=12345678 -d postgres:14
# $ psql --host=127.0.0.1 --username=postgres --password
```

### Create Database

Create a Database called `bookings`

### Create tables

If you don't have migrations folder, follow the steps list bellow

1. [Set the database.yml conf in your root folder](https://gobuffalo.io/en/docs/db/configuration/)
1. [fizz-migrations](https://gobuffalo.io/en/docs/db/migrations/#fizz-migrations)
    ```bash
    # Create tables
    $ soda generate fizz CreateUserTable
    $ soda generate fizz CreateReservationTable
    $ soda generate fizz CreateRoomsTable
    $ soda generate fizz CreateRestrictionsTable
    $ soda generate fizz CreateRoomRestrictionsTable
    # Create foreign key
    $ soda generate fizz CreateFKForReservationsTable
    $ soda generate fizz CreateFKForRoomReservationsTable
    # Create Index
    $ soda generate fizz CreateUniqueIndexForUsersTable
    $ soda generate fizz CreateIndicesOnRoomRestrictions

    $ soda generate fizz AddFKAndIndicesToReservationTable
    $ soda generate fizz AddNotNullToReservationIDForRestrictions

    # Migration for seeding rooms
    $ soda generate sql SeedRoomsTable
    # Migration for seeding restrictions
    $ soda generate sql SeedRestrictionsTable
    ```
1. [Put setting into the create_create_user_tables.up.fizz file which is in migrations folder](https://gobuffalo.io/en/docs/db/fizz#create-a-table)

Generate the columns of users table

```bash
# This command will reference the create_create_user_tables.up.fizz to create the content and the table of users
$ soda migrate
```

### soda db management cli tool
> In this project, we are using the `soda` to manage database
> https://gobuffalo.io/en/docs/db/toolbox/
>
> After installing `soda` cli, check its path by `whcih soda`