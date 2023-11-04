# initial
EDIT = "edit"
EXIT = "exit"


# helpers
def print_records_by_action(records, action, printable_actions, print_reversed=False):
    if action in printable_actions:
        records_items = reversed(records.items()) if print_reversed else records.items()
        for key, value in records_items:
            print(f"{key} {value}")

# input and process
def input_action(records):
    action, *values = input().split(" ")
    if action == EDIT and len(values) == 2:
        records[values[0]] = values[1]
    elif action == EXIT:
        return records, action
    else:
        print("invalid action or values")
    return records, action

def input_length_records():
    return int(input())

def input_records_by_length(length):
    result = {}

    for _ in range(length):
        save_data = input().split(" ")
        result[save_data[0]] = save_data[1]

    return result


def run():
    length_records = input_length_records()
    records = input_records_by_length(length_records)

    while True:
        records, action = input_action(records)
        print_records_by_action(
            records=records,
            action=action,
            printable_actions=[EDIT],
            print_reversed=True,
        )

        if action == EXIT:
            break


if __name__ == "__main__":
    run()
