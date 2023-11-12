Fetch Movies
=================================================


#### Question
```
Write an HTTP GET to retrieve movies.

BaseUrl & Query:
https://jsonmock.hackerrank.com/api/moviesdata/search?Title=substr&page=pageNumber

Response:
page, per_page, total, total_pages
data: [
    {
        "Title", "Year", "imdbID"
    }
]

Provide Code to return Titles of movies by title (contains)
```


#### Example
```
spiderman // input
The Making of 'Waterworld'
Waterworld
Waterworld 4: History of the Islands
```