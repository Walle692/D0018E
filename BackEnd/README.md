## To Test Version 1

navigate to version 1 and use "go run ."

then in cmd use 

curl -X POST http://localhost:5000/login -H "Content-Type: application/json" -d "{\"username\":\"yourusername\",\"password\":\"yourpassword\"}" -i