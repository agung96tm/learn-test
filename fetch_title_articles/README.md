Fetch Title Articles
=================================================


#### Question
```
Write an HTTP GET articles.

BaseUrl & Query:
https://jsonmock.hackerrank.com/api/articles?author=&page=

Response:
    page
    per_page
    total
    total_pages
    data: [
        {
            title
            url
            author
            num_comments
            story_id
            story_title
            story_url
            parent_id
            created_at
        }
    ]

Provide Code to return Titles of articles by author

With Condition:
    if title is empty
        then return story as title
 
    if title and story are null
        then ignore article
```


#### Example
```
epega // input
A Message to Our Customers
Apple's declining software quality
```