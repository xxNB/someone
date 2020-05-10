import random

MAX_DEPTH = 5


class ListNode:
    def __init__(self, key=None, value=None):
        self._key = key
        self._value = value
        self._forwards = []


class SkipList:

    _MAX_LEVEL = 4

    def __init__(self):
        self._level_count = 1
        self._head = ListNode()
        self._head._forwards = [None] * self._MAX_LEVEL

    def find(self, value):
        """
        查找一个元素，返回一个 ListNode 对象
        """
        pass

    def find_range(self, begin_value, end_value):
        """
        查找一个元素，返回一组有序 ListNode 对象
        """
        pass

    def insert(self, value):
        """
        插入一个元素，返回 True 或 False
        """
        level = self._random_level()
        if self._level_count < level:
            self._level_count = level
        new_node = ListNode(value)
        new_node._forwards = [None] * level
        update = [self._head] * level   
        # update数组为层级索引，用以存储新节点所有层数上，各自的前一个节点的信息
        # forward[]数组，用以存储该节点所有层的下一个节点的信息
        p = self._head
        for i in range(level - 1, -1, -1):
            while p._forwards[i] and p._forwards[i]._data < value:
                p = p._forwards[i]
            if p._forwards[i] and p._forwards[i]._data == value:
                return False
            update[i] = p
        for i in range(level):
            new_node._forwards[i] = update[i]._forwards[i]
            update[i]._forwards[i] = new_node
        return True

    def delete(self, value):
        """
        删除一个元素，返回 True 或 False
        """
        update = [None] * self._level_count
        p = self._head
        for i in range(self._level_count - 1, -1, -1):
            while p._forwards[i] and p._forwards[i]._data < value:
                p = p._forwards[i]
            update[i] = p

        if p._forwards[0] and p._forwards[0]._data == value:
            for i in range(self._level_count - 1, -1, -1):
                if update[i]._forwards[i] and update[i]._forwards[i]._data == value:
                    update[i]._forwards[i] = (
                        update[i]._forwards[i]._forwards[i]
                    )  # Similar to prev.next = prev.next.next
            return True
        else:
            return False

    def _random_level(self, p=0.5):
        """
        返回随机层数
        """
        level = 1
        while random.random() < p and level < self._MAX_LEVEL:
            level += 1
        return level

    def pprint(self):
        """
        打印跳表
        """
        pass
