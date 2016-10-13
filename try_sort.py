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

"""
MERGE SORT
"""

def merge(a, b):
    """merge two already sorted lists"""
    result = []
    a = deque(a)
    b = deque(b)
    while len(a) > 0 and len(b) > 0:
        if a[0] < b[0]:
            result.append(a.popleft())
        else:
            result.append(b.popleft())
            
    while len(a) > 0:
        result.append(a.popleft())
    while len(b) > 0:
        result.append(b.popleft())
        
    return result + list(a) + list(b)
        

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
    if len(array) <= 1:
        return array
    left = divide(array)[0]
    right = divide(array)[1]
    left = merge_sort(left)
    right = merge_sort(right)
    return merge(left, right)
    

        
if __name__ == "__main__":
    print("Test Subject 1", test_list)
    print(love_sort(test_list))  
    print("Merge")
    print(merge_sort(test_list))
