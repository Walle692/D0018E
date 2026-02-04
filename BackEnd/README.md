# To Test Version 2

## Step 1

Navigate into the version 2 folder

## Step 2

Run

```sh
go mod tidy
```

## Step 3

Run

```sh
go run .
```

## Step 4

Use the following command in cli

```sh
curl -X POST http://localhost:5000/login -H "Content-Type: application/json" -d "{\"username\":\"yourusername\",\"password\":\"yourpassword\"}" -i
```

Be aware of the differences between linux and windows curl

