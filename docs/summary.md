# Summary Examples

Here you'll find examples that is made up of a combination of the basic data structures and algorithms:

* [BIFID Cipher](#bifid-cipher)
* [Roman numerals](#roman-numerals)

## BIFID Cipher

The BIFID cipher uses a grid and was invented by Felix Delastelle in 1901. In its simplest form it creates a grid and maps the letters into numeric values.

For illustrations, let's attempt to scramble the plaintext "maryland" using BIFID.

### Requirements

The steps to implement this are as follows:

<u>Step 1: Encoding</u>

The first letter "m" corresponds to "4" and "5". We place "4" in the first row and "5" in the second row. Continuing with this process we have:

|  |  |  |  |  |  |  |  |
| -- | -- | -- | -- | -- | -- | -- | -- |
| m | a | r | y | l | a | n | d |
| 4 | 3 | 5 | 5 | 5 | 3 | 2 | 2 |
| 5 | 3 | 5 | 3 | 3 | 3 | 3 | 4 |

Next we read along the rows and merge, to give:

43 55 43 22 53 53 33 34

And we convert them back to the letters from the grid, we get:

L R L P Y Y A X

<u>Step 2: Decode</u>

Let's try to revise this cipher DXETE, and we get:

24 34 35 51 35

We can then arrange these into:

|  |  |  |  |  |
| -- | -- | -- | -- | -- |
| 2 | 4 | 3 | 4 | 3 |
| 5 | 5 | 1 | 3 | 5 |

### Working examples

Please refer to:

* [Executable code](../cmd/bifid-cypher/main.go)
* [Unit tests](../cmd/bifid-cypher/main_test.go)

### Reference

Cryptogrpahy. William J. Buchanan, OBE. Edinburgh Napier University, UK.

## Roman numerals

Roman numerals have *seven* symbols:

| Symbol | Value |
| -- | -- |
|  I | 1    |
|  V | 5    |
|  X | 10   |
|  L | 50   |
|  C | 100  |
|  D | 500  |
|  M | 1000 |

### Requirements

The value of a sequence of numerals is calculated by adding or subtracting numeric value as you read the characters from left to right. When you have read a numeral that is smaller or equal to the next one you add to the accumulated sum. When the next numeral is smaller than the previous numeral, you deduct from the accumulated sum.

### Working examples

* [Executable codes](../cmd/roman-numerals/main.go)
* [Unit tests](../cmd/roman-numerals/main_test.go)