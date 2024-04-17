
# Test Database Package

This package provides a simple and convenient way to set up a PostgreSQL instance using Docker for testing purposes in Go.


## Features

- Sets up a PostgreSQL with least settings.
- Tears down the instance after the test run.


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
	err := con.Setup()
	assert.NoError(t, err)

    m.Run()

	err = con.TearDown()
	assert.NoError(t, err)
}
```

## License

[MIT](https://choosealicense.com/licenses/mit/)

