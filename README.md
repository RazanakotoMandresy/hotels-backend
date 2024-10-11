# An restfull api of an hotels's reservation

## Configuring the database

<p> export your postgres url </p>

```bash
$ export POSTGRESQL_URL='postgres://postgres:yourpasssword@localhost:5432/the_name_of_the_db?sslmode=disable'
```

<p>  run migration files by executing the command: </p>

```bash
$ migrate -database ${POSTGRESQL_URL} -path migrations up
```

## To run the backend

<p> run it with the make commandes </p>

```bash
$ make run
```
