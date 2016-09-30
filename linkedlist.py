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
        if new_next:
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

    def get_next(self):
        if not self.is_last():
            return self.reference
                
    def get_list(self):
        """get_list returns a list of all nodes """
        li = []
        li.append(self)
        look_backwards = not self.is_first()
        # First get Forward
        if not self.is_last():
            nxt = self.reference
        else:
            return li
        
        while True:
            li.append(nxt)
            if nxt.is_last():
                break
            else:
                nxt = nxt.reference
        if look_backwards:
            bck = self.parent
            while True:
                li.insert(0, bck)
                if bck.is_first():
                    break
                else:
                    bck = bck.parent
                    
        return li
                
    def is_first(self):
        if not self.parent:
            return True
        else:
            return False
                
    def is_last(self):
        if not self.reference:
            return True
        else:
            return False

    def find_last(self):
        if self.is_last():
            return self
        else:
            return self.reference.find_last()


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
    print("Parent of Second:", n.reference.parent)
    print("Is the first last?", n.is_last())
    n.reference.reference.rem()
    print("Is the third last?", n.reference.reference.is_last())
    n_last = n.find_last()
    print("The last is:", n_last)
    node_list = n.get_list()
    print("Printing Node List:")
    for node in node_list:
        print(node)
    print("Print parent of second node:", node_list[1].parent)
    node_list = n.reference.get_list()
    print("Printing Node List, Collected from Second:")
    for node in node_list:
        # Ought to be the same as previous list
        print(node)
        
