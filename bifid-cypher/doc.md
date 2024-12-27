# BIFID Cipher

The BIFID cipher uses a grid and was invented by Felix Delastelle in 1901. In its simplest form it creates a grid and maps the letters into numeric values.

For illustrations, let's attempt to scramble the plaintext "maryland" using BIFID.

## Encode

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

## Decode

Let's try to revise this cipher DXETE, and we get:

24 34 35 51 35

We can then arrange these into:

|  |  |  |  |  |
| -- | -- | -- | -- | -- |
| 2 | 4 | 3 | 4 | 3 |
| 5 | 5 | 1 | 3 | 5 |

## Working examples

* [main.go](./main.go)

## Reference

Cryptogrpahy. William J. Buchanan, OBE. Edinburgh Napier University, UK.
