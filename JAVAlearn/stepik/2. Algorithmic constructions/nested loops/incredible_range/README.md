# Task Description

The user inputs two integers (the boundaries of a segment on the number line). The boundaries may be entered incorrectly (the first number is greater than the second). In this case, the boundaries should be swapped, ensuring that the iteration always goes from the smaller to the larger number.

The program should find the first number within this segment that has the maximum sum of its digits. For negative numbers, the sign should not be considered when calculating the sum of digits. For example, the sum of digits for -324 is 9.

## Test Data

| Test Number | Input Data | Output Data |
|-------------|------------|-------------|
| 1           | 15 22      | 19          |
| 2           | -16 5      | -9          |
| 3           | 58 47      | 49          |
| 4           | 6 18       | 9           |
| 5           | -58 -47    | -58         |
| 6           | 567 582    | 579         |
| 7           | 11235 11250| 11249       |