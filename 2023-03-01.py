from datetime import datetime


def gettime():
    return datetime.now().strftime("%H:%M:%S")


class X:
    def __init__(self):
        self.logs = []
        self._v = 0

    @property
    def v(self):
        return self._v

    @v.setter
    def v(self, val):
        self.logs.append((val, gettime()))
        self._v = val


x = X()
for i in range(5):
    x.v = i
print(x.v, x.logs)
# 4 [(0, '17:12:32'), (1, '17:12:32'), (2, '17:12:32'), (3, '17:12:32'), (4, '17:12:32')]


for item in inventory.items:
    $iname = item.name + "_s.png"
    imagebutton auto iname action Call("ShowDescription", item)

label ShowDescription(item=None):
    t "[item.description]"
    $ renpy.pause(hard=True)
