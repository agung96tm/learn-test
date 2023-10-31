function calculateTax(income, taxBrackets) {
    let taxResult = 0;
    let remainingIncome = income;
    const calculationList = [];

    for (const bracket of taxBrackets) {
        const [bracketLimit, taxRate] = bracket;
        if (remainingIncome > 0) {
            const taxableAmount = Math.min(remainingIncome, bracketLimit);
            taxResult += taxableAmount * taxRate;
            remainingIncome -= taxableAmount;

            // Printable format
            const formattedTaxable = taxableAmount.toLocaleString('en-US', { useGrouping: false });
            const formattedTaxRate = (taxRate * 100).toFixed(0) + '%';
            calculationList.push(`${formattedTaxable.replace(',', '.') + ' * ' + formattedTaxRate}`);
        } else {
            break;
        }
    }

    return [calculationList, taxResult];
}

const taxBracketsList = [[30000000, 0.05], [100000000, 0.10], [Infinity, 0.20]];

const readline = require('readline');
const rl = readline.createInterface({
    input: process.stdin,
    output: process.stdout
});

rl.question("input: ", (incomeInput) => {
    const income = parseFloat(incomeInput);
    const [printableCalculation, tax] = calculateTax(income, taxBracketsList);
    const formattedTax = tax.toLocaleString('en-US', { useGrouping: false });

    console.log("output: " + printableCalculation.join(' + ') + ' = ' + formattedTax.replace(',', '.'));
    rl.close();
});
