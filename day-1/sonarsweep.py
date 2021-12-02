textFile = open("input", "r")
lines = textFile.readlines()
textFile.close()

measurPoints = [int(line) for line in lines]

dephtIncrease=0
for index, value in enumerate(measurPoints):
    if index > 0 and measurPoints[index-1] < value:
        dephtIncrease+=1

print(dephtIncrease)