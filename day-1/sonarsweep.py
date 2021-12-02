textFile = open("input", "r")
lines = textFile.readlines()
textFile.close()

measurPoints = [int(line) for line in lines]

def getDephtIncrease(entries):
    dephtIncrease=0
    for index, value in enumerate(entries):
        if index > 0 and entries[index-1] < value:
            dephtIncrease+=1
    return dephtIncrease

def consolidateMeasure(entries, chunkSize, overlap):
    result = []
    for index in range(0, len(entries), chunkSize-overlap):
        result.append(sum(entries[index:index + chunkSize]))
    return result

print("Depthincrease single measures:", getDephtIncrease(measurPoints))

flattendMeasurePoints=consolidateMeasure(measurPoints, 3, 2)

print("Flattend points dephincrease:", getDephtIncrease(flattendMeasurePoints))
