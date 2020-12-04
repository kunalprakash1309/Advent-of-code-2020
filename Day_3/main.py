def Part1(course):
    treecounter = 0
    xcounter = 3
    for length in range(1, len(course)):
        print("----------------------------------  start  ---------------------------")
        print(course[length])

        print("X value", xcounter)
        print("Y value", length)

        if(course[length][xcounter] == "#"):
            treecounter = treecounter + 1
        if (xcounter + 3) > len(course[length]) - 1:
            xcounter = (xcounter + 3) - len(course[length])
        else:
            xcounter = xcounter + 3

        print("------------------sum:= ", treecounter)

        print("----------------------------------  end  ---------------------------")

    return(treecounter)

def Part2(xslope, yslope, course):
    treecounter = 0
    xcounter = xslope

    for length in range(yslope, len(course), yslope):
        if(course[length][xcounter] == "#"):
            treecounter = treecounter + 1

        if (xcounter + xslope) > len(course[length]) - 1:
            xcounter = (xcounter + xslope) - len(course[length])
        else:
            xcounter = xcounter + xslope

    return(treecounter)


with open("input.txt") as file:
  course = [line.strip() for line in file]

print('{count} is the number of trees in Part1'.format(count=Part1(course)))

print('{count} is the number of trees in Part2'.format(count=(Part2(1, 1, course) * Part2(3, 1, course) * Part2(5, 1, course) * Part2(7, 1, course) * Part2(1, 2, course))))