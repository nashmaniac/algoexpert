


def isValidPart(ipString) -> bool:
    stringAsInt = int(ipString)
    if stringAsInt > 255:
        return False
    return len(str(stringAsInt)) == len(ipString)    
    

def validIPAddresses(ipString):
    # Write your code here.
    totalIPAddressList = []
    for i in range(1, min(len(ipString), 4)):
        currentIPAddress = ['', '', '', '']

        firstPart = ipString[:i]
        if not isValidPart(firstPart):
            continue
        
        currentIPAddress[0] = firstPart
        for j in range(i+1, i+min(len(ipString)-i, 4)):
            secondPart = ipString[i:j]
            if not isValidPart(secondPart):
                continue
            currentIPAddress[1] = secondPart
            
            for k in range(j+1, j+min(len(ipString)-j, 4)):
                thirdPart = ipString[j:k]
                fourthPart = ipString[k:]
                
                if isValidPart(thirdPart) and isValidPart(fourthPart):
                    currentIPAddress[2] = thirdPart
                    currentIPAddress[3] = fourthPart
                    
                    totalIPAddressList.append('.'.join(currentIPAddress))
    return totalIPAddressList


if __name__ == "__main__": 
    validIPAddresses(ipString="1921680")