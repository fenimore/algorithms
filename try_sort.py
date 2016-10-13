# python 3.5
# Sorting
from collections import deque

test_list = [3, 2,  5, 12, 7, 1, 44, 666]
test_2_list = [1, 4, 5, 7, 8]


          
def love_sort(input_array):
    """My attempt at a sort algorithm.
    
    The idea is to insert an item into the 
    beginning of a list unless it is greater
    than the first item.
    """
    items = deque(input_array) # input
    array = deque([])          # result
    array.append(items.popleft())
    
    for item in items:
        idx = 0
        if item < array[idx]:
            array.insert(idx, item)
        elif item > array[idx]:
            while item > array[idx]:
                idx += 1
                if idx+1 > len(array):
                    break
            array.insert(idx, item)

    return list(array)

def merge(a, b):
    """Takes two values, and returns a list of them
    merged...Or rather, with the first values of each
    set to a proper position?
    """
    y = a[0]
    z = b[0]
    a.remove(a[0])
    b.remove(b[0])
    if y < z:
        x = [y, z]
    else:
        x = [z, y]
    return x + a + b

def divide(array):
    """Divides an array into a list of two lists
    where the input is cut in half"""
    a = array[:int(len(array)/2)]
    b = array[int(len(array)/2):]
    return [a, b]

def merge_sort(array):
    """From Wikipedia:
    First divide the list into the smallest unit (1 element),
    then compare each element with the adjacent list to sort 
    and merge the two adjacent lists. 
    Finally all the elements are sorted and merge.
    """
    # : first means first half
    if len(array) < 1:
        return array
    a = divide(array)[0]
    b = divide(array)[1]
    new = merge(a, b)
    return new

def well(array):
    size = len(array)
    for _ in range(size):
        array = merge_sort(array)
    return array
        
if __name__ == "__main__":
    print("Test Subject 1", test_list)
    print(love_sort(test_list))  
    print("Merge")
    print(merge_sort(test_list))
    print(well(test_list))
