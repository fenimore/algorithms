# Binary search
# Takes a sorted array


def binary_search(items, target, L=-1, H=-1, m=-1):
    """Pass in a list and a target, and it'll return
    the index of the item you're looking force
    """
    length = len(items)
    if L == -1:
            L = 0          # low bound index
            H = length -1  # index of Upper bound
            m = int(length/2)
    if target > items[H]:
        return -1
    if target == items[m]:
        return m # the index
    elif target < items[m]:
        H = m-1
        m = int(H+L/2)
    elif target > items[m]:
        L = m + 1
        m = int((H+L)/2)
    return binary_search(items, target, L, H, m)
