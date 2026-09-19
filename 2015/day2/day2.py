total_paper = 0

input = open("input.txt")
print(input.read())

# read each character of the file 
# formula: 2*l*w + 2*w*h + 2*h*l.
# read each side, separate for x
# (length l, width w, and height h)
with open("input.txt") as file:
    for dimention in file:
        #3x12x32
        list = dimention.strip("\n").split("x",-1)
        nlist = [int(num) for num in list]
        lowest = min([(nlist[0]*nlist[1]), (nlist[1]*nlist[2]), (nlist[0]*nlist[2])])
        paper = (2*nlist[0]*nlist[1]) + (2*nlist[1]*nlist[2]) + (2*nlist[0]*nlist[2]) + lowest 
        total_paper += paper
        print("this is the dimentions separated", list)
        print("this is the dimention", dimention) 
        print("this is the lowest", lowest)

#closing the file
input.close()

print("this is the total", total_paper) 

# 1448967 -> too low
# 1456998 --> too low 
# 1600273 -> too high
# 1586300 -> tentative (correct one)
