# floors 
floor = 0
index = 0
# reading the file 
input = open("input.txt")
print(input.read())

# read each character of the file 
# if char = ( -> up one floor 
# if char = ) -> down one floor 
with open("input.txt") as file:
    for line in file:
        for char in line: 
            #apply the logic 
                if char == "(":
                    floor += 1
                    index += 1
                if char == ")":
                    floor -= 1
                    index += 1
                if floor == -1: 
                    print("this is the first time he enters: ", index)
                    break 
#closing the file
input.close()


print("this is where he need to go: ", floor )

