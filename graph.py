import os
import matplotlib.pyplot as plt
import json
import sys

def calcCollatz(num):
    collatz = os.popen("./ccg {}".format(num)).read()
    collatz = collatz.replace(" ", ", ")
    return json.loads(collatz)

def plot(nums):
    for n in nums:
        plt.plot(n, label="{}".format(n[0]))
    plt.ylabel("Calculated Value")
    plt.xlabel("Steps")
    #plt.legend()
    plt.show()


out = []
nums = sys.argv[1:]
for n in nums:
    out.append(calcCollatz(n))
plot(out)

