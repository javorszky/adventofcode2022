# Day 13:  Distress Signal

[back to index](https://github.com/javorszky/adventofcode2022/)

## Part 1

You climb the hill and again try contacting the Elves. However, you instead receive a signal you weren't expecting: a distress signal.

Your handheld device must still not be working properly; the packets from the distress signal got decoded out of order. You'll need to re-order the list of received packets (your puzzle input) to decode the message.

Your list consists of pairs of packets; pairs are separated by a blank line. You need to identify how many pairs of packets are in the right order.

For example:
```
[1,1,3,1,1]
[1,1,5,1,1]

[[1],[2,3,4]]
[[1],4]

[9]
[[8,7,6]]

[[4,4],4,4]
[[4,4],4,4,4]

[7,7,7,7]
[7,7,7]

[]
[3]

[[[]]]
[[]]

[1,[2,[3,[4,[5,6,7]]]],8,9]
[1,[2,[3,[4,[5,6,0]]]],8,9]
```
Packet data consists of lists and integers. Each list starts with [, ends with ], and contains zero or more comma-separated values (either integers or other lists). Each packet is always a list and appears on its own line.

When comparing two values, the first value is called left and the second value is called right. Then:

* If both values are integers, the lower integer should come first. If the left integer is lower than the right integer, the inputs are in the right order. If the left integer is higher than the right integer, the inputs are not in the right order. Otherwise, the inputs are the same integer; continue checking the next part of the input.
* If both values are lists, compare the first value of each list, then the second value, and so on. If the left list runs out of items first, the inputs are in the right order. If the right list runs out of items first, the inputs are not in the right order. If the lists are the same length and no comparison makes a decision about the order, continue checking the next part of the input.
* If exactly one value is an integer, convert the integer to a list which contains that integer as its only value, then retry the comparison. For example, if comparing [0,0,0] and 2, convert the right value to [2] (a list containing 2); the result is then found by instead comparing [0,0,0] and [2].

Using these rules, you can determine which of the pairs in the example are in the right order:
```
== Pair 1 ==
- Compare [1,1,3,1,1] vs [1,1,5,1,1]
    - Compare 1 vs 1
    - Compare 1 vs 1
    - Compare 3 vs 5
        - Left side is smaller, so inputs are in the right order

== Pair 2 ==
- Compare [[1],[2,3,4]] vs [[1],4]
    - Compare [1] vs [1]
        - Compare 1 vs 1
    - Compare [2,3,4] vs 4
        - Mixed types; convert right to [4] and retry comparison
        - Compare [2,3,4] vs [4]
            - Compare 2 vs 4
                - Left side is smaller, so inputs are in the right order

== Pair 3 ==
- Compare [9] vs [[8,7,6]]
    - Compare 9 vs [8,7,6]
        - Mixed types; convert left to [9] and retry comparison
        - Compare [9] vs [8,7,6]
            - Compare 9 vs 8
                - Right side is smaller, so inputs are not in the right order

== Pair 4 ==
- Compare [[4,4],4,4] vs [[4,4],4,4,4]
    - Compare [4,4] vs [4,4]
        - Compare 4 vs 4
        - Compare 4 vs 4
    - Compare 4 vs 4
    - Compare 4 vs 4
    - Left side ran out of items, so inputs are in the right order

== Pair 5 ==
- Compare [7,7,7,7] vs [7,7,7]
    - Compare 7 vs 7
    - Compare 7 vs 7
    - Compare 7 vs 7
    - Right side ran out of items, so inputs are not in the right order

== Pair 6 ==
- Compare [] vs [3]
    - Left side ran out of items, so inputs are in the right order

== Pair 7 ==
- Compare [[[]]] vs [[]]
    - Compare [[]] vs []
        - Right side ran out of items, so inputs are not in the right order

== Pair 8 ==
- Compare [1,[2,[3,[4,[5,6,7]]]],8,9] vs [1,[2,[3,[4,[5,6,0]]]],8,9]
    - Compare 1 vs 1
    - Compare [2,[3,[4,[5,6,7]]]] vs [2,[3,[4,[5,6,0]]]]
        - Compare 2 vs 2
        - Compare [3,[4,[5,6,7]]] vs [3,[4,[5,6,0]]]
            - Compare 3 vs 3
            - Compare [4,[5,6,7]] vs [4,[5,6,0]]
                - Compare 4 vs 4
                - Compare [5,6,7] vs [5,6,0]
                    - Compare 5 vs 5
                    - Compare 6 vs 6
                    - Compare 7 vs 0
                        - Right side is smaller, so inputs are not in the right order
                          What are the indices of the pairs that are already in the right order? (The first pair has index 1, the second pair has index 2, and so on.) In the above example, the pairs in the right order are 1, 2, 4, and 6; the sum of these indices is 13.
```
Determine which pairs of packets are already in the right order. What is the sum of the indices of those pairs?

### Solution

Hardest part here was coming up with a data store that allowed me to represent infinitely recursive data types.

In the end I created an `item` interface that implemented three methods:
* `Day13()`, which does nothing, but limits the implementations to this interface
* `String() string` to return a string representation of the data we're looking at
* `Type() int` which returns the concrete type of the implementation.

There are also two iota constants to denote the type: `typeList` and `typeInteger`.

The two concrete types are called `integer`, which is based on the built in `int`, and `list`, which is `[]item`. Note that here I'm using the `item` interface type, so a list can be made up of a bunch of `integers`, some `list`s, etc... This way I can nest it however I want.

Next up was parsing the individual lines. The parser function returns a list, ie something that's wrapped in a `[]`, as every packet is a list.

Then I go character by character:
* is this a `[`? grab the rest of the string, including that char, and send it to the parser function, recursively, and grab the resulting `list` that comes out, and add it to the currently open `list`
* is this a `]`? neat, we have a complete list, return the currently open list from the parser function. Also check whether we've been collecting numbers, and if we have, parse that, add that to the list, and then return
* is this a `,`? cool, check if we've been collecting numbers, and if so, add that to the list. We've already added the preceding sub-list to the currently open list
* is this an anything else? presumably a number. Add it to the `strings.Builder`

That gives me a neat little data type that looks like this:
```
// [2,[[8],[9,0]],[1]]
list{
    integer(2),
    list{
        list{
            integer(8)
        },
        list{
            integer(9),
            integer(0),
        },
    },
    list{
        integer(1),
    },
}
```

Next step is to compare the two. The comparison function is also a recursive one that compares two `list`s.

* start by figuring out the lengths of the two lists and creating a fallback return, in case the contents of the list do not make a decision on whether it's in the correct order
  * same length: fallback is "continue"
  * left is shorter: fallback is "correct order"
  * right is shorter: fallback is "incorrect order"
* then start a for loop that goes from 0 to the length of the shorter of the two, and for each index
  * checks the two types. (remember the `.Type() int` method on the interface?) There are four possibilities
    * element from left is integer, elem from right is integer => compare ints
      * left is smaller, return correct
      * right is smaller, return incorrect
      * they're the same, `continue` with the for loop
    * elem from left is int, right is list
      * convert left int into a list with a single int element
      * send both the left as list, and right list to the compare function (recursion)
      * check return value, if the value is continue, continue, otherwise return decision
    * elem from left is list, right is list
      * send both to the compare function (recursion)
      * check value, if continue, continue, otherwise return decision
    * elem from left is list, right is int
      * convert right int to list
      * send both lists to compare function (recursion)
      * check value, if continue, continue, otherwise return decision

At the end each group needs to return either correct / incorrect. If they returned continue once the full strings were parsed, that's an error and stop.

Add the indices (starting from 1) to an accumulator where the decision was "correct order", and that's the solution.
## Part 2

Now, you just need to put all of the packets in the right order. Disregard the blank lines in your list of received packets.

The distress signal protocol also requires that you include two additional divider packets:
```
[[2]]
[[6]]
```
Using the same rules as before, organize all packets - the ones in your list of received packets as well as the two divider packets - into the correct order.

For the example above, the result of putting the packets in the correct order is:
```
[]
[[]]
[[[]]]
[1,1,3,1,1]
[1,1,5,1,1]
[[1],[2,3,4]]
[1,[2,[3,[4,[5,6,0]]]],8,9]
[1,[2,[3,[4,[5,6,7]]]],8,9]
[[1],4]
[[2]]
[3]
[[4,4],4,4]
[[4,4],4,4,4]
[[6]]
[7,7,7]
[7,7,7,7]
[[8,7,6]]
[9]
```
Afterward, locate the divider packets. To find the decoder key for this distress signal, you need to determine the indices of the two divider packets and multiply them together. (The first packet is at index 1, the second packet is at index 2, and so on.) In this example, the divider packets are 10th and 14th, and so the decoder key is 140.

Organize all of the packets into the correct order. What is the decoder key for the distress signal?

### Solution

* parse and comparison function are already done
create a slice to hold all the parsed lists, parse all of them, and add them to the slice in whatever order
* parse the 2 divider and 6 divider strings into lists, and add them to the slice with the others
* do a `sort.Slice` passing in a `less` function that uses the comparison function between two `list`s to establish less / more relationship that the sorter can use
* create an accumulator value with 1
* iterate over the now sorted list, if the string representation of the item matches the 2 divider or 6 divider strings, multiply the accumulator by whatever the 1-based index of the current item is
* that accumulator is the solution
