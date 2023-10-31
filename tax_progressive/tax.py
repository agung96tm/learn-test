def calculate_tax(income, tax_brackets):
    tax_result = 0
    remaining_income = income
    calculation_list = []

    for bracket in tax_brackets:
        bracket_limit, tax_rate = bracket
        if remaining_income > 0:
            taxable_amount = min(remaining_income, bracket_limit)
            tax_result += taxable_amount * tax_rate
            remaining_income -= taxable_amount

            # printable format
            formatted_taxable = format(int(taxable_amount), ",").replace(",", ".")
            formatted_tax_rate = "{:.0%}".format(tax_rate)
            calculation_list.append(f"{formatted_taxable} * {formatted_tax_rate}")
        else:
            break

    return calculation_list, tax_result

tax_brackets_list = [(30_000_000, 0.05), (100_000_000, 0.10), (float('inf'), 0.20)]

# Input the income
income_input = float(input("input: "))
printable_calculation, tax = calculate_tax(income_input, tax_brackets_list)
formatted_tax = format(int(tax), ",").replace(",", ".")
print("output: ", " + ".join(printable_calculation), "=", formatted_tax)