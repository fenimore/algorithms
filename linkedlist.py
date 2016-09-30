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
        
    def list_nodes(self):
        print(self)
        nxt = self.reference
        print(nxt)
        while True:
            if not nxt.reference:
                break
            else:
                print(nxt.reference)
                nxt = nxt.reference


if __name__ == "__main__":
    n = Node("The First")
    n.add("The Second")
    n.add("The Third")
    n.add("The Fourth")
    print(n.list_nodes())
    n.rem()
    print(n.list_nodes())
    
