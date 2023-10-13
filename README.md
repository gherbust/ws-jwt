# Work Shop Uso de JWT-Middleware

Taller para explicar implementacion de JWT con una palabra secreta.

#### Configuracion Archivo launch.json para VSCode

``{
    // Use IntelliSense para saber los atributos posibles.
    // Mantenga el puntero para ver las descripciones de los existentes atributos.
    // Para más información, visite: https://go.microsoft.com/fwlink/?linkid=830387
    "version": "0.2.0",
    "configurations": [
        {
            "name": "GetTocken",
            "type": "go",
            "request": "launch",
            "mode": "debug",
            "program": "gentoken.go"
        },
        {
            "name": "WebApi",
            "type": "go",
            "request": "launch",
            "mode": "debug",
            "program": "cmd/api/main.go"
        }
    ]
}``


#### Endpoints
Crear token:
``
curl --location 'http://localhost:8080/api/v1/login' \
--header 'Content-Type: application/json' \
--data '{
    "name":"gerry",
    "password":"Pass123$",
    "client_key":"cliente 1"
}' 
``

Validar
``
curl --location 'http://localhost:8080/api/v1/valid' \
--header 'Authorization: tu JWT'
``

Get Pets
``
curl --location 'http://localhost:8080/api/v1/pets' \
--header 'Authorization: tu JWT'
``

Bibligrafias:
- go jwt https://pkg.go.dev/github.com/golang-jwt/jwt/v5@v5.0.0
 - gin Web framework middleware https://gin-gonic.com/zh-tw/docs/examples/using-middleware/
