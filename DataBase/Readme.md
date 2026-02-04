# Database command guide

## Loading a schema into postgres

### Step 1

Make the DataBase folder your working directory (IN THE VM)

### Step 2

Run 
```sh
export RDSHOST="d0018e.cx6ecgasovnp.eu-north-1.rds.amazonaws.com
```

### Step 3

Run
```sh
psql -f schema.sql "host=$RDSHOST port=5432 dbname=postgres user=postgres sslmode=verify-full sslrootcert=/certs/global-bundle.pem password=<Enter_DB_Password>"
```
Make sure that sslrootcert path is correct and you have downloaded global-bundle.pem from aws.

## Adding a user to the database

to add a user to the db use this

```sh
INSERT INTO myschema.users (username, password,role) VALUES ('admin','admin','admin');
```

### Show tables in schema

```sh
\dt myschema.*
```

### Show data in table

```sh
SELECT * FROM myschema.users;
```