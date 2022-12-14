# Day 14: Regolith Reservoir

[back to index](https://github.com/javorszky/adventofcode2022/)

## Part 1

The distress signal leads you to a giant waterfall! Actually, hang on - the signal seems like it's coming from the waterfall itself, and that doesn't make any sense. However, you do notice a little path that leads behind the waterfall.

Correction: the distress signal leads you behind a giant waterfall! There seems to be a large cave system here, and the signal definitely leads further inside.

As you begin to make your way deeper underground, you feel the ground rumble for a moment. Sand begins pouring into the cave! If you don't quickly figure out where the sand is going, you could quickly become trapped!

Fortunately, your familiarity with analyzing the path of falling material will come in handy here. You scan a two-dimensional vertical slice of the cave above you (your puzzle input) and discover that it is mostly air with structures made of rock.

Your scan traces the path of each solid rock structure and reports the x,y coordinates that form the shape of the path, where x represents distance to the right and y represents distance down. Each path appears as a single line of text in your scan. After the first point of each path, each point indicates the end of a straight horizontal or vertical line to be drawn from the previous point. For example:

498,4 -> 498,6 -> 496,6
503,4 -> 502,4 -> 502,9 -> 494,9
This scan means that there are two paths of rock; the first path consists of two straight lines, and the second path consists of three straight lines. (Specifically, the first path consists of a line of rock from 498,4 through 498,6 and another line of rock from 498,6 through 496,6.)

The sand is pouring into the cave from point 500,0.

Drawing rock as #, air as ., and the source of the sand as +, this becomes:

```
  4     5  5
  9     0  0
  4     0  3
0 ......+...
1 ..........
2 ..........
3 ..........
4 ....#...##
5 ....#...#.
6 ..###...#.
7 ........#.
8 ........#.
9 #########.
```
Sand is produced one unit at a time, and the next unit of sand is not produced until the previous unit of sand comes to rest. A unit of sand is large enough to fill one tile of air in your scan.

A unit of sand always falls down one step if possible. If the tile immediately below is blocked (by rock or sand), the unit of sand attempts to instead move diagonally one step down and to the left. If that tile is blocked, the unit of sand attempts to instead move diagonally one step down and to the right. Sand keeps moving as long as it is able to do so, at each step trying to move down, then down-left, then down-right. If all three possible destinations are blocked, the unit of sand comes to rest and no longer moves, at which point the next unit of sand is created back at the source.

So, drawing sand that has come to rest as o, the first unit of sand simply falls straight down and then stops:
```
......+...
..........
..........
..........
....#...##
....#...#.
..###...#.
........#.
......o.#.
#########.
```
The second unit of sand then falls straight down, lands on the first one, and then comes to rest to its left:
```
......+...
..........
..........
..........
....#...##
....#...#.
..###...#.
........#.
.....oo.#.
#########.
```
After a total of five units of sand have come to rest, they form this pattern:
```
......+...
..........
..........
..........
....#...##
....#...#.
..###...#.
......o.#.
....oooo#.
#########.
```
After a total of 22 units of sand:
```
......+...
..........
......o...
.....ooo..
....#ooo##
....#ooo#.
..###ooo#.
....oooo#.
...ooooo#.
#########.
```
Finally, only two more units of sand can possibly come to rest:
```
......+...
..........
......o...
.....ooo..
....#ooo##
...o#ooo#.
..###ooo#.
....oooo#.
.o.ooooo#.
#########.
```
Once all 24 units of sand shown above have come to rest, all further sand flows out the bottom, falling into the endless void. Just for fun, the path any new sand takes before falling forever is shown here with ~:

```
.......+...
.......~...
......~o...
.....~ooo..
....~#ooo##
...~o#ooo#.
..~###ooo#.
..~..oooo#.
.~o.ooooo#.
~#########.
~..........
~..........
~..........
```
Using your scan, simulate the falling sand. How many units of sand come to rest before sand starts flowing into the abyss below?

### Solution

#### Parse the input into discreet coordinates
* for each line, break them up by ` -> `
* then for each resulting coordinate, grab the x, y and turn them into integers
* loop through the list of coordinates (now integer pair) until the penultimate one, and generate the missing points between the current, and the next element
  * for example if the element is [2,6] and the next one is [2,9], the entire generated line is going to be [2,6], [2,7], [2,8], [2, 9]
* store those into a global slice. This is where the rocks are
* while doing that, also keep track of the min max of row and col
  * row's lowest value is 0, this is the height the sand's entry point is, max is whatever max y value happens to be
  * col's lowest value is 1<<32 (a ridiculously large number), highest is 0 to start with, and then for each point in the input adjust this
* once we have our min max of row / col coordinates, and our list of rocks, let's build the map in three steps
  * step 1: for each row:
    * for each column in that row:
      * turn the coordinate into a single number using the helper function
      * mark that coordinate as air, `.` on the map
    * this should give us a rectangle of all air
  * step 2: for each item in the list of rocks:
    * turn their coordinate into the single int using the helper function
    * mark it as rock `#`
  * step 3: mark 0, 500 as the entry point for sand with `+`

#### Deal with sand

* create a new struct that represents a unit of sand
* on initialisation it receives the current snapshot of the grid, air, rocks, entry point, and existing sands at rest
* set the starting coordinate to 0, 500 (entry point)
* have helper functions that give us the next coordinate for down-left, down, down-right
* start looping. On each iteration:
  * check whether it can move down, ie the thing at the `down` coordinate relative to current coordinate is `air`
    * if it is, set the current coordinate to the down coordinate, and iterate
    * if it is not, it doesn't matter what other thing it is, continue checking
  * check down-left:
    * if air, move there
    * if not, continue checking
  * check down-right:
    * if air, move there
    * if not, we've come to rest
  * if during any point the next coordinate doesn't exist on the world (not in the map), then it flows to the abyss, and we bubble this error back up
  * if we've come to rest, return the resting coordinate, and an "at rest" error

#### Putting the two together

* generate the world from the inputs
* create a counter with 0 as starting value
* start an infinite for loop. In each iteration:
  * create a new sand from the current state of the grid
  * loop until it finds a resting position
  * check for return error:
    * if it's an abyss error, we're done, break
    * if it's an at rest error:
      * grab returned coordinate, add that to the grid
      * increment counter
      * loop again
* the only time we could break out from the loop is if one of the sands was yeeted into the abyss
* that means we have a counter, which is the solution

## Part 2

You realize you misread the scan. There isn't an endless void at the bottom of the scan - there's floor, and you're standing on it!

You don't have time to scan the floor, so assume the floor is an infinite horizontal line with a y coordinate equal to two plus the highest y coordinate of any point in your scan.

In the example above, the highest y coordinate of any point is 9, and so the floor is at y=11. (This is as if your scan contained one extra rock path like -infinity,11 -> infinity,11.) With the added floor, the example above now looks like this:

```
        ...........+........
        ....................
        ....................
        ....................
        .........#...##.....
        .........#...#......
        .......###...#......
        .............#......
        .............#......
        .....#########......
        ....................
<-- etc #################### etc -->
```

To find somewhere safe to stand, you'll need to simulate falling sand until a unit of sand comes to rest at 500,0, blocking the source entirely and stopping the flow of sand into the cave. In the example above, the situation finally looks like this after 93 units of sand come to rest:

```
............o............
...........ooo...........
..........ooooo..........
.........ooooooo.........
........oo#ooo##o........
.......ooo#ooo#ooo.......
......oo###ooo#oooo......
.....oooo.oooo#ooooo.....
....oooooooooo#oooooo....
...ooo#########ooooooo...
..ooooo.......ooooooooo..
#########################
```

Using your scan, simulate the falling sand until the source of the sand becomes blocked. How many units of sand come to rest?

### Solution

Reusing the above.

#### Modification to the inputs and the grid

* once all lines of the input have been parsed, we'll add the floor:
* grab the `rowMax` value and add 2 to it to get the height where our rock floor is going to sit
* check if 500 - rowMax - 2 (2 for padding) is less than current colMin, if it is, set colMin to that
* check if 500 + rowMax + 2 is more than current colMax, and if it is, set colMax to that
* once we have our updated row / col min / max values, draw air everywhere
* draw rocks from the input
* generate a new line of rocks between [rowMax, 500-rowMax-2] and [rowMax, 500+rowMax+2] for the infinite floor
* add those as rocks as well

#### Sand
Same as above, no change needed

#### Altogether

Only difference is the success criteria. Instead of checking for an abyss error in the infinite for loop, we're checking for the rest coordinate being the same as the entry point.
