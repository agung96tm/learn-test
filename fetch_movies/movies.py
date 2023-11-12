import requests

BASE_URL = "https://jsonmock.hackerrank.com/api/moviesdata/search"

def get_movies(title):
    movies = []
    page = 1

    while True:
        try:
            response = requests.get(f"{BASE_URL}?page={page}&Title={title}")
            data = response.json().get('data', [])
            if not data:
                break
            movies.extend(data)
            page += 1
        except Exception as e:
            print("something happens when fetching data: ", str(e))
            break

    return movies

def get_title_movies(movies):
    title_movies = list(set(movie["Title"] for movie in movies))
    return sorted(title_movies)

def execute():
    title = input("")
    movies = get_movies(title)
    movies_names = get_title_movies(movies)
    print("\n".join(movies_names))

if __name__ == "__main__":
    execute()
