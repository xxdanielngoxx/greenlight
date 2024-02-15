# Greenlight
The purpose of this project is for learning go language.

## Endpoints
|Method|URL Pattern|Handler|Action|
|:-----|:---------:|:-----:|-----:|
|GET|/v1/healthcheck|healthcheckHandler|Show application information
|POST|/v1/movies|createMovieHandler|Create a new movie
|GET|/v1/movies/:id|showMovieHandler|Show the details of a specific movie