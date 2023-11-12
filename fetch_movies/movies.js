const BASE_URL = "https://jsonmock.hackerrank.com/api/moviesdata/search";

const getMovies = async (title) => {
    let movies = [];
    let page = 1;

    while (true) {
        try {
            const response = await fetch(`${BASE_URL}?page=${page}&Title=${title}`);
            const data = await response.json();
            const moviesData = data.data || [];
            if (!moviesData.length) {
                break
            } else {
                movies.push(...moviesData);
                page++;
            }
        } catch (error) {
            console.log("Something happens when fetching data: ", error.message);
            break
        }
    }

    return movies
}

const getTitleMovies = (movies) => {
    const titleMovies = [...new Set(movies.map(movie => movie.Title))];
    return titleMovies.sort();
}

const execute = async () => {
    const title = await getUserInput();
    const movies = await getMovies(title);
    const moviesNames = getTitleMovies(movies);
    console.log(moviesNames.join("\n"));
}

const getUserInput = () => {
    return new Promise(resolve => {
        const readline = require('readline').createInterface({
            input: process.stdin,
            output: process.stdout
        });
        readline.question('', title => {
            readline.close();
            resolve(title);
        });
    });
}

if (require.main === module) {
    execute();
}