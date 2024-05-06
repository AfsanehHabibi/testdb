
# Test Database Package

This package provides a simple and convenient way to set up a PostgreSQL instance using Docker for testing purposes in Go.


## Features

- Sets up a PostgreSQL with least settings.
- Tears down the instance after the test run.

## Installation

To install TestDB, use `go get`:

```bash
go get github.com/AfsanehHabibi/testdb
```

## Usage/Examples
To use this package in your test suite, follow these steps:
1. Import the package:

```go
import (
    "github.com/AfsanehHabibi/testdb/pgsql"
    "github.com/stretchr/testify"
)
```

2. Use the Setup function to set up the PostgreSQL instance before your test suite:

```go
func TestMain(m *testing.M) {
    con := TestPG{}
	err := con.Setup(nil)
    // wait for postgres to start
    time.Sleep(3*time.Second)

    m.Run()

	err = con.TearDown()
}
```
### Config

The `Config` struct is used to configure the setup of the PostgreSQL container. It includes the following fields:

- `ImageName`: The name of the Docker image to use. Default value is "postgres:latest".
- `DBName`: The name of the database to be created. Default value is "default".
- `DBUser`: The username for the database. Default value is "root".
- `DBPassword`: The password for the database user. In default state database is set to trust.
- `PORT`: The port on which the PostgreSQL container will expose. Default value is 5432.

It is not necessary for you to set any of these values, if you leave them unset they will have a default value.

#### Example Usage

```go
config := &testdb.Config{
    ImageName: "postgres:latest",
    DBName: "mydatabase",
    DBUser: "myuser",
    DBPassword: "mypassword",
    PORT: 5432,
}
con := TestPG{}
err := con.Setup(config)
```
### Execute
The Execute function allows you to run a SQL query inside the PostgreSQL container. It takes a query string as a parameter and executes it within the container. It is recommended to connect to database with whatever library you prefer using database information instead of using this funcion directly. Wait a few seconds before invoking Execute to ensure the database is fully operational.

#### Example Usage

```go
err := con.Execute("SELECT * FROM users;")
if err != nil {
    log.Fatal(err)
}
```

## License

[MIT](https://choosealicense.com/licenses/mit/)


