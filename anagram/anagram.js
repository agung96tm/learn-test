function areAnagramsWithBuiltinFunc(str1, str2) {
  // Remove spaces and convert to lowercase for case-insensitive comparison
  str1 = str1.replace(/ /g, '').toLowerCase();
  str2 = str2.replace(/ /g, '').toLowerCase();

  // Check if the sorted characters of both strings are equal
  return str1.split('').sort().join('') === str2.split('').sort().join('');
}

function areAnagramsManual(str1, str2) {
  // Remove spaces and convert to lowercase for case-insensitive comparison
  str1 = str1.replace(/ /g, '').toLowerCase();
  str2 = str2.replace(/ /g, '').toLowerCase();

  if (str1.length !== str2.length) {
    return false;
  }

  const charCount1 = {};
  const charCount2 = {};

  for (const char of str1) {
    charCount1[char] = (charCount1[char] || 0) + 1;
  }

  for (const char of str2) {
    charCount2[char] = (charCount2[char] || 0) + 1;
  }

  for (const char in charCount1) {
    if (charCount1[char] !== charCount2[char]) {
      return false;
    }
  }

  return true;
}

const readline = require('readline');

const rl = readline.createInterface({
  input: process.stdin,
  output: process.stdout
});

rl.question("Input: ", (firstStr) => {
  rl.question("Input: ", (lastStr) => {
    const isAnagram = areAnagramsManual(firstStr, lastStr);
    // const isAnagram = areAnagramsWithBuiltinFunc(firstStr, lastStr);
    console.log("Is anagram? " + (isAnagram ? "Yes" : "No"));
    rl.close();
  });
});