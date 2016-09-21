# Arithmetic with only a plus operator
# Yup
import unittest


def minus(first, second):
    return first + -second

def multiply(multiplied, multiplier):
    result = 0
    for i in range(multiplier):
        result += multiplied

    return result

def divide(divided, divisor):
    """divide returns 0 if not whole number"""
    result = 0
    for i in range(divided):
        if i == 1 or i == 0:
            continue
        target = multiply(i, divisor)
        if target == divided:
            result = i
            break

    return result


class test_arithmetic(unittest.TestCase):
    def test_minus(self):
        self.assertEqual(minus(5, 3), 2)

    def test_multiply(self):
        self.assertEqual(multiply(5, 3), 15)

    def test_divide_1(self):
        self.assertEqual(divide(15, 3), 5)

    def test_divide_2(self):
        self.assertEqual(divide(20, 3), 0)

def tests():
    unittest.main()

if __name__ == '__main__':
    tests()
