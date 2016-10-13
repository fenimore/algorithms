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

            

if __name__ == "__main__":
    print("Test Subject 1", test_list)
    print(love_sort(test_list))  
