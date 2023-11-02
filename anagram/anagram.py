def are_anagrams_with_builtin_func(str1, str2):
    # Remove spaces and convert to lowercase for case-insensitive comparison
    str1 = str1.replace(" ", "").lower()
    str2 = str2.replace(" ", "").lower()

    # Check if the sorted characters of both strings are equal
    return sorted(str1) == sorted(str2)

def are_anagrams_manual(str1, str2):
    str1 = str1.replace(" ", "").lower()
    str2 = str2.replace(" ", "").lower()

    if len(str1) != len(str2):
        return False

    char_count1 = {}
    char_count2 = {}

    for char in str1:
        if char in char_count1:
            char_count1[char] += 1
        else:
            char_count1[char] = 1

    for char in str2:
        if char in char_count2:
            char_count2[char] += 1
        else:
            char_count2[char] = 1

    return char_count1 == char_count2

first_str = input("Input: ")
last_str = input("Input: ")

# choose, manual or built-in function
print("Is anagram? " + "Yes" if are_anagrams_manual(first_str, last_str) else "No")
# print("Is anagram? " + "Yes" if are_anagrams_with_builtin_func(first_str, last_str) else "No")