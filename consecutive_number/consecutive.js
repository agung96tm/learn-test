const readline = require('readline');
const rl = readline.createInterface({ input: process.stdin, output: process.stdout });

function findConsecutiveNumbers(targetSum) {
  const resultList = [];
  let start = 1;
  let end = 1;
  let currentSum = start;

  while (start <= Math.floor(targetSum / 2)) {
    if (currentSum === targetSum) {
      if (start !== end) {  // Avoid single numbers
        const sequence = Array.from({ length: end - start + 1 }, (_, i) => start + i);
        resultList.push(sequence);
      }
      end += 1;
      currentSum += end;
    } else if (currentSum < targetSum) {
      end += 1;
      currentSum += end;
    } else {
      currentSum -= start;
      start += 1;
    }
  }

  return resultList;
}

rl.question("input: ", function (user_input) {
  const result = findConsecutiveNumbers(parseInt(user_input));
  const resultStrings = result.map(sequence => `[${sequence.join(', ')}]`);
  console.log(`output: ${resultStrings.join(', ')}`);
  rl.close();
});