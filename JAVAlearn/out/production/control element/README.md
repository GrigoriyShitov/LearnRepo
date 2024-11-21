# Task Description

The user inputs the size of the array and the initial value for the random number generator.

Then, the user inputs a control element (an integer).

Create an array of the specified size and fill it with random numbers ranging from 2 to 15 (inclusive). Sort the array in ascending order. Print the sorted array using the `toString()` method, in square brackets and separated by commas.

Find the index of the control element in the sorted array and remove all elements after it (create a new array of the required size).

If the control element is not present in the array, output "ERROR". Otherwise, print the resulting array using the `toString()` method.

## Test Data

| Test Number | Input Data | Output Data                         |
|-------------|------------|-------------------------------------|
| 1           | 8 11       | [5, 6, 7, 10, 11, 13, 14, 15]       |
|             | 10         | [5, 6, 7, 10]                       |
| 2           | 10 25      | [3, 3, 6, 8, 8, 10, 10, 11, 13, 15] |
|             | 13         | [3, 3, 6, 8, 8, 10, 10, 11, 13]     |
| 3           | 7 16       | [4, 7, 7, 9, 12, 14, 15]            |
|             | 22         | ERROR                               |
| 4           | 5 33       | [5, 6, 7, 13, 15]                   |
|             | 5          | [5]                                 |
