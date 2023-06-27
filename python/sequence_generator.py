from random import randint

def generate(size = 10000000, duplicates = False):
    '''Generates a random number sequence of specified size,
    avoids creating duplicate numbers in the sequence by default'''
    randomSequence = []

    for i in range(size):
        if duplicates:
            randomSequence.append(randint(0, size - 1))
        else:
            randomSequence.insert(randint(0, len(randomSequence)), i)
    
    return randomSequence