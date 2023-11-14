const BASE_URL = "https://jsonmock.hackerrank.com/api/articles";


const getArticles = async (author) => {
    let articles = [];
    let page = 1;

    while (true) {
        try {
            const params = new URLSearchParams({page: page});
            if (author !== null || author !== "") {
                params.append("author", author)
            }

            const response = await fetch(`${BASE_URL}?${params.toString()}`);
            const data = (await response.json())?.data || [];

            if (data.length === 0) {
                break
            }

            articles.push(...data);
            page++;
        } catch (e) {
            console.error("Something happens when fetch data: ", e);
            break;
        }
    }

    return articles;
}

const getTitleArticles = async (articles) => {
    const titles = [...new Set(
        articles.map(article => article?.title || article?.story_title).filter(title => title)
    )];
    return titles.sort();
}

const execute = async () => {
    try {
        const readline = require('readline');
        const rl = readline.createInterface({ input: process.stdin, output: process.stdout});

        rl.question("", async (authorName) => {
            const articles = await getArticles(authorName.trim() || null);
            const articleNames = await getTitleArticles(articles);
            console.log(articleNames.join("\n"));
            rl.close();
        });
    } catch (error) {
        console.error("Error:", error);
    }
}

execute();