### Database command guide

## Loading a schema into postgres

# Step 1

Make the DataBase folder your working directory (IN THE VM)

# Step 2

Run "export RDSHOST="d0018e.cx6ecgasovnp.eu-north-1.rds.amazonaws.com" 

# Step 3

Run "psql -f schema.sql "host=$RDSHOST port=5432 dbname=postgres user=postgres sslmode=verify-full sslrootcert=/certs/global-bundle.pem password=<Enter_DB_Password>""


## Adding a user to the database

to add a user to the db use this

"INSERT INTO myschema.users (username, password,role) VALUES ('admin','admin','admin');"