# Bibit Assessment - No 2

## Prerequisites
please before you run this application please have installed some this requirement:

1. Go 1.12 and above, this project make is using go version 1.16, if you have dont go installed in your computer, download [here](https://golang.org/dl/)

## Usage

### Run Test
To run a test you can simply run the `make test` command and to see the coverage of the test you can run` make cover` (make sure your computer can run Makefile). Don't forget to go into each directory to do the test.

### Run Application - Without Docker Compose

#### Movies Service
First, you must running movie service before running API Gateway.
before running it, make sure you go into movies directory using `cd movies`, then set the environment variable for this application, edit `.env.example` to `.env` then set:

```env
OMDB_URL= [omdb url]
OMDB_API_KEY= [your main omdb api key]
DB_HOST= [your main database host]
DB_PORT= [your main database port]
DB_USER= [your main database username]
DB_PASS= [your main database password]
DB_NAME= [your main database name which will be used]
DB_SSL_MODE= [SSL MODE (this is only for the postgres database)]
```

after complete setting environment variables then you can run the command `make run` (make sure your computer can run Makefile) to build the application and wait for the build process done.
make sure you have run postgres database.
#### API Gateway
After running movies service we can run API Gateway for handle request from client, to running it, make sure you go into gateway directory using `cd gateway`, then set the environment variable for this application, edit `.env.example` to `.env` then set:

```env
PORT=8080
MOVIE_PORT=8081
```

after complete setting environment variables then you can run the command `make run` (make sure your computer can run Makefile) to build the application and wait for the build process done. then the application will be running.

### Run Application - With Docker Compose
By using docker compose you can easily run this application. make sure you have installed Docker on your computer or laptop.

first build a docker image for this application by running `make docker-build` in each directory service (e.g movie and gateway). After building the docker image, you can immediately run docker compose by executing the command `docker-compose up`, please wait for the application to run.

here below is `docker-compose.yaml` file:

```yaml
version: "3.9"
services:

  # postgres db
  postgres:
    image: postgres:13.3-alpine
    restart: always
    ports: 
      - 5432:5432
    environment:
      POSTGRES_PASSWORD: root
      POSTGRES_DB: bibit
    healthcheck:
      test: ["CMD-SHELL", "pg_isready -U postgres"]
      interval: 10s
      timeout: 5s
      retries: 5

  movie:
    image: bibit-movie:1.0.0
    restart: always
    ports:
      - 8081:8081
    depends_on:
      postgres:
        condition: service_healthy
    environment:
      OMDB_URL: http://www.omdbapi.com/
      OMDB_API_KEY: faf7e5bb
      DB_HOST: postgres
      DB_PORT: 5432
      DB_USER: postgres
      DB_PASS: root
      DB_NAME: bibit
      DB_SSL_MODE: disable
      PORT: 8081

  gateway:
    image: bibit-gateway:1.0.0
    restart: always
    ports:
      - 8080:8080
    depends_on:
      - movie
    environment:
      PORT: 8080
      MOVIE_PORT: 8081
      MOVIE_SERVICE_HOST: movie
```

## HTTP Endpoint

### Search Movies

- URL:

    `http://localhost:8080/api/v1/movies/search?s=Batman&page=2`
    
    (adjust the port you use)

- Method:

    GET

- Example Response:

    ```json
    {
        "movies": [
            {
                "Title": "Batman: The Killing Joke",
                "Year": "2016",
                "imdbID": "tt4853102",
                "Type": "movie",
                "Poster": "https://m.media-amazon.com/images/M/MV5BMTdjZTliODYtNWExMi00NjQ1LWIzN2MtN2Q5NTg5NTk3NzliL2ltYWdlXkEyXkFqcGdeQXVyNTAyODkwOQ@@._V1_SX300.jpg"
            },
            {
                "Title": "Batman: The Dark Knight Returns, Part 2",
                "Year": "2013",
                "imdbID": "tt2166834",
                "Type": "movie",
                "Poster": "https://m.media-amazon.com/images/M/MV5BYTEzMmE0ZDYtYWNmYi00ZWM4LWJjOTUtYTE0ZmQyYWM3ZjA0XkEyXkFqcGdeQXVyNTA4NzY1MzY@._V1_SX300.jpg"
            },
            {
                "Title": "Batman: Mask of the Phantasm",
                "Year": "1993",
                "imdbID": "tt0106364",
                "Type": "movie",
                "Poster": "https://m.media-amazon.com/images/M/MV5BYTRiMWM3MGItNjAxZC00M2E3LThhODgtM2QwOGNmZGU4OWZhXkEyXkFqcGdeQXVyNjExODE1MDc@._V1_SX300.jpg"
            },
        ]
    }
    ```

### Get Detail Movies

- URL:

    `http://localhost:8080/api/v1/movie/tt4853102`
    
    (adjust the port you use)

- Method:

    GET

- Example Response:

    ```json
    {
        "Title": "Batman: The Killing Joke",
        "Year": "2016",
        "imdbID": "tt4853102",
        "Type": "movie",
        "Poster": "https://m.media-amazon.com/images/M/MV5BMTdjZTliODYtNWExMi00NjQ1LWIzN2MtN2Q5NTg5NTk3NzliL2ltYWdlXkEyXkFqcGdeQXVyNTAyODkwOQ@@._V1_SX300.jpg",
        "Rated": "R",
        "Released": "25 Jul 2016",
        "Runtime": "76 min",
        "Genre": "Animation, Action, Crime",
        "Director": "Sam Liu",
        "Writer": "Brian Azzarello, Brian Bolland, Bob Kane",
        "Actors": "Kevin Conroy, Mark Hamill, Tara Strong",
        "Plot": "As Batman hunts for the escaped Joker, the Clown Prince of Crime attacks the Gordon family to prove a diabolical point mirroring his own fall into madness.",
        "Language": "English",
        "Country": "United States",
        "Awards": "1 win & 2 nominations",
        "ImdbRating": "6.4"
    }
    ```

## Hope

I hope what i do is satisfactory to you. and i hope i can join the Bibit Company. start contributing there as much as possible and improve my skills there. Thank You,

## Creator

Risky Feryansyah Pribadi