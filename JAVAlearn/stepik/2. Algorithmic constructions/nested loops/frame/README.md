# Task Description

The user inputs the height and width of a frame separated by a space. The program should draw a frame made of asterisks.

If invalid data is entered (height or width <= 0), the program should output "ERROR".

**Note:** The empty spaces inside the frame should be filled with spaces.

## Test Data

```markdown
| Test Number | Input Data | Output Data    |
|-------------|------------|----------------|
| 1           | 5 4        | ****           |
|             |            | *  *           |
|             |            | *  *           |
|             |            | *  *           |
|             |            | ****           |
| 2           | -1 5       | ERROR          |
| 3           | 1 2        | **             |
| 4           | 5 7        | *******        |
|             |            | *     *        |
|             |            | *     *        |
|             |            | *     *        |
|             |            | *******        |
| 5           | 3 2        | **             |
|             |            | **             |
|             |            | **             |
| 6           | 5 -1       | ERROR          |
| 7           | 1 1        | *              |