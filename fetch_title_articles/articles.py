import requests

BASE_URL = "https://jsonmock.hackerrank.com/api/articles"


def get_articles(author_name):
    article = []
    page = 1

    while True:
        try:
            # params
            params = {"page": page}
            if author_name is not None:
                params["author"] = author_name

            response = requests.get(BASE_URL, params=params)
            data = response.json().get('data', [])
            if not data:
                break

            article.extend(data)
            page += 1
        except Exception as e:
            print("something happens when fetching data: ", str(e))
            break

    return article

def get_title_articles(article):
    title_article = list(
        set(
            article["title"] or article["story_title"]
            for article in article
            if article["title"] or article["story_title"]
        )
    )
    return sorted(title_article)

def execute():
    author_name = input("")
    articles = get_articles(author_name)
    article_names = get_title_articles(articles)
    print("\n".join(article_names))

if __name__ == "__main__":
    execute()
