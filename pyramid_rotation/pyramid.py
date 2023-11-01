def create_pyramid(_height, _rotation):
    pyramid = []
    _rotation = _rotation % 4

    if _rotation == 0:
        for i in range(1, _height + 1):
            spaces = ' ' * (_height - i)
            stars = '*' * (2 * i - 1)
            pyramid.append(spaces + stars)
    elif _rotation == 1:
        for i in range(1, _height + 1):
            pyramid.append('*' * i)
        for i in range(_height - 1, 0, -1):
            pyramid.append('*' * i)
    elif _rotation == 2:
        for i in range(_height, 0, -1):
            spaces = ' ' * (_height - i)
            stars = '*' * (2 * i - 1)
            pyramid.append(spaces + stars)
    elif _rotation == 3:
        for i in range(1, _height + 1):
            spaces = ' ' * (_height - i)
            stars = '*' * i
            pyramid.append(spaces + stars)
        for i in range(height - 1, 0, -1):
            spaces = ' ' * (height - i)
            stars = '*' * i
            pyramid.append(spaces + stars)

    if len(pyramid) == 0:
        print("invalid input")

    for p in pyramid:
        print(p)


height = int(input("height: "))
rotation = int(input("rotation: "))

create_pyramid(height, rotation)
