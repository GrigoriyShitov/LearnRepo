# Task Description

The user inputs the size of the array and the initial value for the random number generator.

Create an array of integers with the specified size and fill it with random numbers ranging from 10 to 20 (inclusive). Print the original array using the `toString()` method, in square brackets and separated by commas.

Then, the user inputs two indices (the indices in the tests are valid). Sort the portion of the array between these two indices in ascending order (including both indices).

On a new line, output the transformed array to the console using the `toString()` method.

## Test Data

| Test Number | Input Data | Output Data                              |
|-------------|------------|------------------------------------------|
| 1           | 10 85      | [20, 14, 10, 19, 18, 18, 20, 20, 16, 11] |
|             |            | [20, 10, 14, 18, 19, 18, 20, 20, 16, 11] |
| 2           | 8 32       | [12, 17, 17, 18, 12, 13, 19, 13]         |
|             |            | [12, 17, 17, 12, 13, 18, 19, 13]         |
| 3           | 10 125     | [15, 18, 10, 16, 12, 14, 18, 18, 11, 14] |
|             |            | [10, 12, 14, 15, 16, 18, 18, 18, 11, 14] |
