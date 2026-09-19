x = 0
y = 0
a = 0
b = 0 
santa = "" 
robot = ""
santa_dic = {
    f"{x},{y}": 1
}

robot_dic = {
    f"{a},{b}": 1
}


#remove santa steps 
#remove robo-santa steps
# repeat same process
 
def check(key,santa_dic):
    if key in santa_dic:
       santa_dic[key] += 1 
    else: 
        santa_dic[key] = 1

input = open("input.txt")

#separate steps 
with open("input.txt") as file:
    text = file.read()
    for i, c in enumerate(text): 
        if i % 2 == 0: 
            santa += c
        else: 
            robot += c

print("santa steps: ", len(santa))
print("robot steps: ", len(robot))
print("sum of steps: ", len(santa) + len(robot))
print("total steps: ", len(text))
input.close()

for char in santa: 
    if char == "^":
        y += 1
        key = f"{x},{y}"
        check(key, santa_dic)
    elif char == "v":
        y -= 1
        key = f"{x},{y}"
        check(key, santa_dic)
    elif char == ">":
        x += 1 
        key = f"{x},{y}"
        check(key, santa_dic)
    elif char == "<":
        x -= 1
        key = f"{x},{y}"
        check(key, santa_dic)

for char in robot: 
    if char == "^":
        b += 1
        key = f"{a},{b}"
        check(key, robot_dic)
    elif char == "v":
        b -= 1
        key = f"{a},{b}"
        check(key, robot_dic)
    elif char == ">":
        a += 1 
        key = f"{a},{b}"
        check(key, robot_dic )
    elif char == "<":
        a -= 1
        key = f"{a},{b}"
        check(key, robot_dic)


# check for only houses
houses =  set( santa_dic.keys() | robot_dic.keys())
print("total houses: ", len(houses))


# 2492 -> too low 
# 2631 -> 2631
