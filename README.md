# dorayaki-backend

Welcome to my dorayaki stock management backend app repository for my programming lab assistant test submission!

### Deployment Link: [http://165.22.58.239:8080/](http://165.22.58.239:8080/)

## How to Run

Make sure you have [docker desktop](https://www.docker.com/products/docker-desktop) installed. Clone the repository, then run the app with:

```bash
docker-compose up
```
The app will run in [http://localhost:8080/]()

Stop the app with:
```bash
docker-compose down
```

## API Request
### Get
* api/v1/toko/ = find all store
* api/v1/toko/:id = find store with id
* api/v1/toko/query = find store with query (ex: query?query=Jakarta)
* api/v1/dorayaki/ = find all dorayaki
* api/v1/dorayaki/:id = find dorayaki with id
* api/v1/stock/ = find all stock
* api/v1/stock/toko/:id = find stock by store id
### Post
* api/v1/toko/ = create store
* api/v1/dorayaki/ = create dorayaki
* api/v1/stock/ = create stock
### Put
* api/v1/toko/:id = update store with id
* api/v1/dorayaki/:id = update dorayaki with id
* api/v1/stock/:id = update stock with id
### Delete
* api/v1/toko/:id = delete store with id
* api/v1/dorayaki/:id = delete dorayaki with id
* api/v1/stock/:id = delete stock with id


## Programming Language and Libraries
1. [Go](https://golang.org/)
2. [Gin](https://github.com/gin-gonic/gin)
3. [Gorm](https://gorm.io/index.html)
4. [smapping](https://github.com/mashingan/smapping)
5. [godotenv](https://github.com/joho/godotenv)
## Author
Afif Fahreza
18219058
