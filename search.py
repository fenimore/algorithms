# Binary search
# Takes a sorted array

def binary_search_recursive(l, target):
    """Finds if a value is in a sorted list. """
    half = int(len(l)/2)
    middle = l[half]
    if (len(l) < 3) and (l[half] != target):
        return false
    if target == middle:
        return true
    else:
        if target < middle:
            return binary_search(l[:half], target)
        elif target > middle:
            return binary_search(l[half:], target)


def binary_search(items, target):
    length = len(items)
    L = 0
    R = length -1  # index of Upper bound
    m = length/2
    if target == items[m]:
        return m # the index
