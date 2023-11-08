const readline = require('readline');

// function removeItemsInList(originItems, removeItems) {
//     return originItems.filter(item => !removeItems.includes(item));
// }

function removeItemsInList(originItems, removeItems) {
    const result = [];

    for (let i = 0; i < originItems.length; i++) {
        const item = originItems[i];
        if (!removeItems.includes(item)) {
            result.push(item);
        }
    }

    return result;
}

const rl = readline.createInterface({
    input: process.stdin,
    output: process.stdout
});

rl.question("Items: ", itemsInput => {
    rl.question("Remove items: ", removeInput => {
        const originItems = itemsInput.split(' ').map(Number);
        const removeItems = removeInput.split(' ').map(Number);

        const newList = removeItemsInList(originItems, removeItems);
        console.log(newList.join(' '));

        rl.close();
    });
});