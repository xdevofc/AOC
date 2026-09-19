import hashlib 

input = 'yzbqklnj' 
example = 'abcdef'
flag = True 
number = 282749
test = ""



while ( flag ): 

    test = input + str(number)
    res = hashlib.md5(test.encode("UTF-8"))
    first_digits = res.hexdigest()[:6]
    if ( first_digits != '000000' ):
        print("digits: ", first_digits)
        number += 1 
    else: 
        print ("hash: ", res.hexdigest())
        print( "number is: ", number)
        flag = False
        break




#609043 -> too high

