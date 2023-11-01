function mod(dividend, divisor) {
  return ((dividend % divisor) + divisor) % divisor;
}

function createPyramid(height, rotation) {
    const pyramid = [];
    rotation = mod(rotation, 4);

    if (rotation === 0) {
        for (let i = 1; i <= height; i++) {
            const spaces = ' '.repeat(height - i);
            const stars = '*'.repeat(2 * i - 1);
            pyramid.push(spaces + stars);
        }
    } else if (rotation === 1) {
        for (let i = 1; i <= height; i++) {
            pyramid.push('*'.repeat(i));
        }
        for (let i = height - 1; i > 0; i--) {
            pyramid.push('*'.repeat(i));
        }
    } else if (rotation === 2) {
        for (let i = height; i > 0; i--) {
            const spaces = ' '.repeat(height - i);
            const stars = '*'.repeat(2 * i - 1);
            pyramid.push(spaces + stars);
        }
    } else if (rotation === 3) {
        for (let i = 1; i <= height; i++) {
            const spaces = ' '.repeat(height - i);
            const stars = '*'.repeat(i);
            pyramid.push(spaces + stars);
        }
        for (let i = height - 1; i > 0; i--) {
            const spaces = ' '.repeat(height - i);
            const stars = '*'.repeat(i);
            pyramid.push(spaces + stars);
        }
    }

    if (pyramid.length === 0) {
        console.log("Invalid input");
    }

    for (const p of pyramid) {
        console.log(p);
    }
}

const readline = require('readline');
const rl = readline.createInterface({
    input: process.stdin,
    output: process.stdout
});

rl.question("Height: ", (height) => {
    rl.question("Rotation: ", (rotation) => {
        createPyramid(Number(height), Number(rotation));
        rl.close();
    });
});