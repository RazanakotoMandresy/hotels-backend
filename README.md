<h1 align="center"> <i>An restfull api of an hotels's reservation </i></h1>

# if it's the the first time you want to run the app

## using the script sh

```bash
$ sh first_running.sh
```

## manual configuration

<i> export your postgres url </i>

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
