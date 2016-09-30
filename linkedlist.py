class Node(object):
    value = ""
    reference = []

    def __init__(self, v):
        self.value = v
        # Add empty Node value?

    def __str__(self):
        represent = self.value # TODO print reference
        return represent
        
    def add(self, v):
        if not self.reference:
            self.reference = Node(v)
        else:
            self.reference.add(v)

    def rem(self):
        new_next = self.reference.reference
        self.reference = new_next
        

def NewNode(value):
    return Node(value)


if __name__ == "__main__":
    n = Node("The first")
    print(n)
    n.add("The Second")
    print(n.reference)
    n.add("The Third")
    print(n.reference.reference)
    n.add("The Fourth")
    print(n.reference.reference.reference)
    n.rem()
    print(n, n.reference, n.reference.reference)
    
