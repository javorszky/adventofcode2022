# Day 12:  Hill Climbing Algorithm

[back to index](https://github.com/javorszky/adventofcode2022/)

## Part 1

You try contacting the Elves using your handheld device, but the river you're following must be too low to get a decent signal.

You ask the device for a heightmap of the surrounding area (your puzzle input). The heightmap shows the local area from above broken into a grid; the elevation of each square of the grid is given by a single lowercase letter, where `a` is the lowest elevation, `b` is the next-lowest, and so on up to the highest elevation, `z`.

Also included on the heightmap are marks for your current position (S) and the location that should get the best signal (E). Your current position (S) has elevation `a`, and the location that should get the best signal (E) has elevation `z`.

You'd like to reach `E`, but to save energy, you should do it in as few steps as possible. During each step, you can move exactly one square up, down, left, or right. To avoid needing to get out your climbing gear, the elevation of the destination square can be at most one higher than the elevation of your current square; that is, if your current elevation is `m`, you could step to elevation `n`, but not to elevation `o`. (This also means that the elevation of the destination square can be much lower than the elevation of your current square.)

For example:
```
Sabqponm
abcryxxl
accszExk
acctuvwj
abdefghi
```
Here, you start in the top-left corner; your goal is near the middle. You could start by moving down or right, but eventually you'll need to head toward the e at the bottom. From there, you can spiral around to the goal:
```
v..v<<<<
>v.vv<<^
.>vv>E^^
..v>>>^^
..>>>>>^
```
In the above diagram, the symbols indicate whether the path exits each square moving up (^), down (v), left (<), or right (>). The location that should get the best signal is still E, and . marks unvisited squares.

This path reaches the goal in 31 steps, the fewest possible.

What is the fewest steps required to move from your current position to the location that should get the best signal?

### Solution

* parse the input into a grid. The grid is essentially a `map[int]int32`. Key is the coordinate in binary ( row << 8 | col ), the value is the character as codepoint where `a = 97`, `z = 122`, `S = 83` and `E = 69` (nice).
* write a recursive function that we will use to visit nodes and do things with it
* write a bunch of helper functions that translate between row, col -> int, and int -> row, col, functions to get the next int coordinate if we were to move left, right, up, down
* then kick off the recursive function with the start point, set the success criteria to reaching the coordinate of the end
* for each point
  * mark it as visited and record the depth (how many steps it took to get there)
  * add the coordinate to the current route
  * have we reached the goal? Excellent, return the entire route back the call stack
  * then grab the coordinates of the neighbours, and then for each
    * check if we're about to move back to the tile we came from. Let's not check that one
    * check if we can move there by
      * does this tile even exist, ie are we going off the bounds of the map? if so, you can't visit
      * is this tile too tall, ie it's 2+ higher than the current? if so, you can't visit
      * have we been to that tile, AND have we been there in fewer steps already? if so, you can't visit
    * if we can visit, let's call the recursive function and pass
      * the coordinate of the new tile, to be recorded in the copy of the route, and it's used to get the elevation of that tile
      * the coordinate of the current one (which will become the previous in the "are we moving back?" check)
      * the current elevation, which will become the previous elevation when comparing tile elevations
      * depth+1 for the steps
      * a copy of the current route
  * collect the returned routes and merge them together and sort them by length
* grab the length of the first route in the slice, and that's the solution -1.

## Part 2

As you walk up the hill, you suspect that the Elves will want to turn this into a hiking trail. The beginning isn't very scenic, though; perhaps you can find a better starting point.

To maximize exercise while hiking, the trail should start as low as possible: elevation a. The goal is still the square marked E. However, the trail should still be direct, taking the fewest steps to reach its goal. So, you'll need to find the shortest path from any square at elevation a to the square marked E.

Again consider the example from above:

```
Sabqponm
abcryxxl
accszExk
acctuvwj
abdefghi
```
Now, there are six choices for starting position (five marked a, plus the square marked S that counts as being at elevation a). If you start at the bottom-left square, you can reach the goal most quickly:

```
...v<<<<
...vv<<^
...v>E^^
.>v>>>^^
>^>>>>>^
```
This path reaches the goal in only 29 steps, the fewest possible.

What is the fewest steps required to move starting from any square with elevation a to the location that should get the best signal?

### Solution

Same as above, except
* start at the end goal
* modify the success criteria to be "elevation == 97" instead of "coordinate == that of the end goal"
* modify the can climb function to be reverse
