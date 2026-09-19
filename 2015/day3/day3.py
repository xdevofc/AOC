# strategy 
'''
we will use a dictionary to store the coordinates
map to extract the unique 
'''

def check(key,dic):
    if key in dic:
       dic[key] += 1 
    else: 
        dic[key] = 1
 
x= 0
y= 0
dic = {
    f"{x},{y}": 1
}
# reading the file 
input = open("input.txt")
print(input.read())

# read each character of the file 
with open("input.txt") as file:
    for line in file:
        for char in line: 
            print("this is the index: ", index)
            if char == "^":
                y += 1
                key = f"{x},{y}"
                check(key, dic)
            elif char == "v":
                y -= 1
                key = f"{x},{y}"
                check(key, dic)
            elif char == ">":
                x += 1 
                key = f"{x},{y}"
                check(key, dic)
            elif char == "<":
                x -= 1
                key = f"{x},{y}"
                check(key, dic)
            print("this is char : ", char)
#closing the file
input.close()

cant = sum( 1 for value in dic.values() if value >= 1 )
print("at least one gift: ", cant)

# one : 2572


