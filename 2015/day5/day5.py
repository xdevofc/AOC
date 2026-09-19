
vowels = 0
repeat = 0
nice = 0
wrong = 0
vowels_list = ['a','e','i','o','u']
last = ""
const = ['ab','cd','pq','xy']



input = open("input.txt")
print(input.read())

# read each character of the file 
with open("input.txt") as file:
    for line in file:

        for x in const: 
            if x in line : 
              wrong += 1 

        if wrong >= 1:
            wrong = 0
            continue 
        else:
            wrong = 0

        for char in line: 
            if char in vowels_list :
               vowels += 1

            if last == char:
                repeat += 1
            else: 
                last = char
            

        if ( vowels >= 3 and repeat >= 1):
            nice  += 1
            vowels = 0
            repeat = 0
            last = ""
            print("is a nice word:", line)
        else: 
            vowels = 0
            repeat = 0
            last = ""
            print("not a nice word:", line)
        

input.close()
print("total nice word: ", nice)

#359 too hihg
#226 too low


