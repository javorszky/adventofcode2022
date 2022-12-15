# Day 15: Beacon Exclusion Zone

[back to index](https://github.com/javorszky/adventofcode2022/)

## Part 1

You feel the ground rumble again as the distress signal leads you to a large network of subterranean tunnels. You don't have time to search them all, but you don't need to: your pack contains a set of deployable sensors that you imagine were originally built to locate lost Elves.

The sensors aren't very powerful, but that's okay; your handheld device indicates that you're close enough to the source of the distress signal to use them. You pull the emergency sensor system out of your pack, hit the big button on top, and the sensors zoom off down the tunnels.

Once a sensor finds a spot it thinks will give it a good reading, it attaches itself to a hard surface and begins monitoring for the nearest signal source beacon. Sensors and beacons always exist at integer coordinates. Each sensor knows its own position and can determine the position of a beacon precisely; however, sensors can only lock on to the one beacon closest to the sensor as measured by the Manhattan distance. (There is never a tie where two beacons are the same distance to a sensor.)

It doesn't take long for the sensors to report back their positions and closest beacons (your puzzle input). For example:

```
Sensor at x=2, y=18: closest beacon is at x=-2, y=15
Sensor at x=9, y=16: closest beacon is at x=10, y=16
Sensor at x=13, y=2: closest beacon is at x=15, y=3
Sensor at x=12, y=14: closest beacon is at x=10, y=16
Sensor at x=10, y=20: closest beacon is at x=10, y=16
Sensor at x=14, y=17: closest beacon is at x=10, y=16
Sensor at x=8, y=7: closest beacon is at x=2, y=10
Sensor at x=2, y=0: closest beacon is at x=2, y=10
Sensor at x=0, y=11: closest beacon is at x=2, y=10
Sensor at x=20, y=14: closest beacon is at x=25, y=17
Sensor at x=17, y=20: closest beacon is at x=21, y=22
Sensor at x=16, y=7: closest beacon is at x=15, y=3
Sensor at x=14, y=3: closest beacon is at x=15, y=3
Sensor at x=20, y=1: closest beacon is at x=15, y=3
```
So, consider the sensor at `2,18`; the closest beacon to it is at `-2,15`. For the sensor at `9,16`, the closest beacon to it is at `10,16`.

Drawing sensors as S and beacons as B, the above arrangement of sensors and beacons looks like this:
```
               1    1    2    2
     0    5    0    5    0    5
 0 ....S.......................
 1 ......................S.....
 2 ...............S............
 3 ................SB..........
 4 ............................
 5 ............................
 6 ............................
 7 ..........S.......S.........
 8 ............................
 9 ............................
10 ....B.......................
11 ..S.........................
12 ............................
13 ............................
14 ..............S.......S.....
15 B...........................
16 ...........SB...............
17 ................S..........B
18 ....S.......................
19 ............................
20 ............S......S........
21 ............................
22 .......................B....
```
This isn't necessarily a comprehensive map of all beacons in the area, though. Because each sensor only identifies its closest beacon, if a sensor detects a beacon, you know there are no other beacons that close or closer to that sensor. There could still be beacons that just happen to not be the closest beacon to any sensor. Consider the sensor at `8,7`:
```
               1    1    2    2
     0    5    0    5    0    5
-2 ..........#.................
-1 .........###................
 0 ....S...#####...............
 1 .......#######........S.....
 2 ......#########S............
 3 .....###########SB..........
 4 ....#############...........
 5 ...###############..........
 6 ..#################.........
 7 .#########S#######S#........
 8 ..#################.........
 9 ...###############..........
10 ....B############...........
11 ..S..###########............
12 ......#########.............
13 .......#######..............
14 ........#####.S.......S.....
15 B........###................
16 ..........#SB...............
17 ................S..........B
18 ....S.......................
19 ............................
20 ............S......S........
21 ............................
22 .......................B....
```
This sensor's closest beacon is at `2,10`, and so you know there are no beacons that close or closer (in any positions marked #).

None of the detected beacons seem to be producing the distress signal, so you'll need to work out where the distress beacon is by working out where it isn't. For now, keep things simple by counting the positions where a beacon cannot possibly be along just a single row.

So, suppose you have an arrangement of beacons and sensors like in the example above and, just in the row where `y=10`, you'd like to count the number of positions a beacon cannot possibly exist. The coverage from all sensors near that row looks like this:
```
                 1    1    2    2
       0    5    0    5    0    5
9 ...#########################...
10 ..####B######################..
11 .###S#############.###########.
```
In this example, in the row where `y=10`, there are `26` positions where a beacon cannot be present.

Consult the report from the sensors you just deployed. In the row where `y=2000000`, how many positions cannot contain a beacon?

### Solution

#### 1. create representations of different things
* coordinate:
  * it's a [2]int that holds x (left-to-right) and y (up-and-down) coordinates
* line
  * two coordinates for endpoints
  * a flag to denote whether it's vertical or horizontal
  * the plane it's on: which row or column it is
  * a constructor that takes in two coordinates, and
    * checks that they're in the same row / column
    * sets the horizontal/vertical flag appropriately
    * orders the two points, so the smaller (more to the left, more to the top) are always first
* lines
  * is a slice of line which we can attach methods to
* sensor
  * coordinate for itself
  * coordinate for the nearest beacon
  * manhattan distance between it and the beacon
  * constructor that takes to coordinates, one for itself, one for the beacon
* the grid
  * slice to hold the sensors present

#### 2. add helper functions

* `absDiff` to help calculating the manhattan distance. It takes two ints, no matter the order passed in, it will be a positive distance between them
* `manhattanDistance`: using absdiff, takes two coordinates, spits out an int
* `uniqueCoordinates`: given a list of coordinates, remove duplicates
* `pluck` generic function, takes in a slice of anything and an index, returns the thing at index, the slice without the thing in it, and optionally an error if something went wrong. For example: [0,1,2,3], you pluck that with index 2, you get 2 (the number at index 2), and [0,1,3] as result
* `mergeLines`: takes two lines, tries to merge them. If successful, returns the single merged line, if not, returns error
* `reduceLines`: takes a slice of lines, and merges them together as much as it can, until no more merges can happen
#### 3. add methods to things
* grid
  * addSensor: takes a sensor, adds it to the list
  * sensorExcludingRow: essentially a filter function, returns a list of sensors that have an exclusion zone that touches the given zone
  * sensorsBeaconsOnRow: filter function, returns a list of sensors that have either the sensor itself, or its beacon on that specific row, ignoring its exclusion zone. Just the device itself
* line, singular
  * Len(): returns the length of the line from start to end, including
  * isCoordInLine: returns whether a given coordinate falls on the line
* lines
  * implements the `sort.Interface` interface, so we can use `sort.Sort(lines)` on it. Sort order:
    * horizontals go before verticals
    * within equal values of horizontal, starts on left goes first
    * within equal values of vertical, starts on top goes first
* sensor
  * `rowInExclusion`: whether the given row touches the exclusion zone of the sensor
  * `lineForRow`: the horizontal slice of the exclusion zone on the given row

#### 4. the methodology

1. parse input into sensor/beacon pairs
2. populate grid
3. grab all sensors that are on the row we're looking at
4. grab all the lines from all the sensors for the row we're looking at
5. reduce those lines (merge them together) so there's no overlap, to guard from double counting
6. grab all the sensors and beacons for the row
7. filter the coordinates of sensors and beacons to be unique
8. filter the coordinates of sensors and beacons by checking them against all the lines, if they aren't on any of the lines, discard them
9. sum up the line lengths
10. subtract the number of beacons / sensors from the sum
11. that is the solution

## Part 2

Your handheld device indicates that the distress signal is coming from a beacon nearby. The distress beacon is not detected by any sensor, but the distress beacon must have x and y coordinates each no lower than 0 and no larger than 4000000.

To isolate the distress beacon's signal, you need to determine its tuning frequency, which can be found by multiplying its x coordinate by 4000000 and then adding its y coordinate.

In the example above, the search space is smaller: instead, the x and y coordinates can each be at most 20. With this reduced search area, there is only a single position that could have a beacon: x=14, y=11. The tuning frequency for this distress beacon is 56000011.

Find the only possible position for the distress beacon. What is its tuning frequency?

### Solution

Same as above, except we're introducing a `clip`: that is used to have an upper and lower bound of the lines being returned for the slices of exclusion zones.

#### Methodology

* for each row from 0 to 4,000,000:
  * grab all sensors that touch that row
  * grab the clipped lines from the sensors
  * merge those together
  * ignore sensors / beacon coordinates 
  * check the length, compare to the full width
  * if the width is less than the full width of the 4,000,000
  * print out the merged lines
  * manually count the product we're looking for
  * that's the solution
