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


routes for the app : 
for user creator :  
- register : method get 
```
http://localhost:<yourport>/users/register
```
- login : method get
```
http://localhost:<yourport>/users/login
```
<br> for the body json you need </br>
``` json
{
  "name":"name",
  "passwords":"passwords",
  "mail":"email@gmail.com"
}
```
- get all hotels
```
http://localhost:<yourport>/hotels
```
- get hotels but filtred by , place , services , opening date ,  max and min budget
 ```
http://localhost:<yourport>/hotels/filter
```
<br> for the body json you need  </br>
``` json
{
  "name":"name",
  "ouverture":"opening date",
  "place":"place",
  "min_budget":min_budget,
  "service":["",""],
}
```
- get an hotels by his uuid uuid
```
http://localhost:<yourport>/hotels/<uuid_hotels>
```
- search hotels by querry
```
http://localhost:<yourport>/hotels/search?hotels=<hotel_name>
```
### for these routes you need to be authentified

- Create an hotels : post method
```
http://localhost:<yourport>/hotels
```
- Update an hotels : put method
```
http://localhost:<yourport>/hotels/<uuid_hotels>
```
- Delete an hotels : delete method
```
http://localhost:<yourport>/hotels/<uuid_hotels>
``` 
- Restore deleted hotels : method put
```
http://localhost:<yourport>/hotels/<uuid_hotels>
```
<h2 style="color: red "> notes : this is app is not finished yet , there are many functionnality that i haven't implement yet :like oauth2 ,anti brut force...  </h2>















