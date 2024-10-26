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

There are two distinct user roles within the system:

- Hotel Owner: Hotel Owners possess comprehensive permissions, enabling them to perform a wide range of actions on hotel listings. These include creating new hotel listings, modifying existing ones, removing listings, restoring previously deleted listings, and accessing detailed read-only information about all hotels.
Guest User: Guest Users have a more limited set of permissions. They can browse and view all available hotel listings, filter search results based on specific criteria, and make reservations for available rooms. Upon booking, a 10% deposit is required. Additionally, Guest Users can rate their stay at a hotel and cancel their reservations if necessary.
Hotel Management Features:

- Hotel Owners have the ability to:

Create new hotel listings: Add comprehensive details about a new hotel property.
Upload images: Include high-quality images to showcase the hotel's amenities and rooms.
Update hotel information: Modify existing hotel listings with updated details such as pricing, availability, and descriptions.
Soft delete and restore hotels: Remove hotel listings temporarily, with the option to restore them at a later date.
