# Print first 13 fibonacci sequence,

def fib(a, b):
    x = b # x is a
    y = a + b # y is b
    print(x, y)
    if x < 100:
        return fib(x, y)


if __name__ == "__main__":
    fib(0, 1)
