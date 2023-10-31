def find_consecutive_numbers(target_sum):
    result_list = []
    start = 1
    end = 1
    current_sum = start

    while start <= target_sum // 2:
        if current_sum == target_sum:
            if start != end:  # Avoid single numbers
                result_list.append(list(range(start, end + 1)))
            end += 1
            current_sum += end
        elif current_sum < target_sum:
            end += 1
            current_sum += end
        else:
            current_sum -= start
            start += 1

    return result_list

user_input = input("input\t: ")
result = find_consecutive_numbers(int(user_input))
result = [str(r) for r in result]

print(f"output\t: " + ", ".join(result))