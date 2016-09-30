class Node(object):
    """Node is the basic object in a linked list"""
    value = ""
    reference = [] # Why is this an Empty node?
    parent = []

    def __init__(self, v):
        """constructor"""
        self.value = v
        # Add empty Node value?

    def __str__(self):
        """print string"""
        represent = self.value # TODO print reference
        return represent
        
    def add(self, v):
        """add"""
        if not self.reference:
            self.reference = Node(v)
            self.reference.parent = self
        else:
            self.reference.add(v)

    def rem(self):
        """remove next node"""
        new_next = self.reference.reference
        self.reference = new_next
        new_next.parent = self

    def ins(self, v):
        """insert"""
        old_reference = self.reference
        self.reference = Node(v)
        self.reference.parent = self
        self.reference.reference = old_reference
        old_reference.parent = self.reference
        
    def list_nodes(self):
        """list all nodes"""
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
    n.ins("The New Second")
    print(n.list_nodes())
    print(n.reference.parent)
    
