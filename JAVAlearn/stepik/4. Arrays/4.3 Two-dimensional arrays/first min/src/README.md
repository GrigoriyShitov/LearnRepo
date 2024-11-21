# Task Description

The user inputs the number of rows and columns for a two-dimensional array, followed by the initial value for the random number generator.

Create an integer array of the specified dimensions and initialize it with random numbers ranging from -10 to 10 (inclusive). Print the array to the console in a tabular format, with elements separated by tab characters. A tab character should also be included at the end of each row.

For each row, find the first negative element and output the index of the corresponding column or "NO" if it is not present.

## Test Data

| Test Number | Input Data | Output Data   |
|-------------|------------|---------------|
| 1           | 4 3 100    | 6	0	-6        |
|             |            | 2	-9	-4       |
|             |            | 7	7	3         |
|             |            | 0	-9	-8       |
|             |            |               |
|             |            | 2             |
|             |            | 1             |
|             |            | NO            |
|             |            | 1             |
| 2           | 4 5 82     | 9	2	-6	-7	-2  |
|             |            | -8	-8	-8	-7	7 |
|             |            | -9	10	-1	2	2  |
|             |            | 8	2	3	8	-3    |
|             |            |               |
|             |            | 2             |
|             |            | 0             |
|             |            | 0             |
|             |            | 4             |
| 3           | 3 3 2 170  | 0	-10         |
|             |            | 9	5           |
|             |            | 5	9           |
|             |            |               |
|             |            | 1             |
|             |            | NO            |
|             |            | NO            |
