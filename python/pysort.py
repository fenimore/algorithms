# Reverse implementation
import unittest

word = "Reverse Me"
palidrome = "step on no pets"

def reverse(word):
    result = ""
    reverse_pos = 0
    length = int(len(word) - 1)
    while length >= 0:
        result += word[length]
        length -= 1
        
    return result


class test_reverse(unittest.TestCase):
    def test_word(self):
        self.assertEqual(reverse(word), "eM esreveR")
    def test_palidrome(self):
        self.assertEqual(reverse(palidrome), "step on no pets")

if __name__ == '__main__':
    unittest.main()
