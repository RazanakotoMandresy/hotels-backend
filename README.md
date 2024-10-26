<h1 align="center"> <i>An rest api of an hotels's reservation </i></h1>

## if it's the the first time you want to run the app

#### using the script sh

```bash
$ sh first_running.sh
```

#### manual configuration

<i> export your postgres url </i>

```bash
$ export POSTGRESQL_URL='postgres://postgres:yourpasssword@localhost:5432/the_name_of_the_db?sslmode=disable'
```

<p>  run migration files by executing the command: </p>

```bash
$ migrate -database ${POSTGRESQL_URL} -path migrations up
```

#### To run the backend

<p> run it with the make commandes </p>

```bash
$ make run
```

<h2 align="center"> functionnality </h2>

#### user :

There are two types of users:

* Owner User: Owner users can create, update, delete, read, in short, manipulate hotels.
Simple User: Can view all hotels, and book a hotel if it is still available with a 10% deposit upon booking, can rate hotels, can cancel reservations.
Hotels:
* When an owner user creates a hotel, the functionalities they can assign to it are:

- Create hotels
- Add images
- Update hotel information
- Delete and restore deleted hotels

From the simple user's point of view once logged in:

- View all hotels
- Make reservations
- Filter results
