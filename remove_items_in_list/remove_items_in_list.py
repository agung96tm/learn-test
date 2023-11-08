# def remove_items_in_list(origin_items: list, remove_items: list) -> list:
#     return list(filter(lambda item: item not in remove_items, origin_items))

def remove_items_in_list(origin_items: list, remove_items: list) -> list:
    result = []
    for item in origin_items:
        if item not in remove_items:
            result.append(item)
    return result


ori_items = [int(v) for v in input("items: ").split()]
rm_items = [int(v) for v in input("remove items: ").split()]

new_list = remove_items_in_list(ori_items, rm_items)
print(" ".join([str(v) for v in new_list]))
