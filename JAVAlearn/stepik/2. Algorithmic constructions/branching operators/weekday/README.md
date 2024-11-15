# Task Description

Write a translation program. The user inputs the name of the day of the week in Russian, and the program outputs this name in English. The input can be given in either uppercase or lowercase letters. However, the name of the day of the week in English is a proper noun, so it should be capitalized.

If the name of the day of the week is entered incorrectly, the program should output "ERROR".

Use a switch statement.

## Test Data

| Test Number | Input Data      | Output Data |
|-------------|------------------|-------------|
| 1           | Понедельник      | Monday      |
| 2           | понедельник      | Monday      |
| 3           | Воскресенье      | Sunday      |
| 4           | воскресенье      | Sunday      |
| 5           | март             | ERROR       |
| 6           | Четверг          | Thursday    |
| 7           | среда            | Wednesday   |
| 8           | Вторник          | Tuesday     |
| 9           | пятница          | Friday      |
| 10          | суббота          | Saturday    |
| 11          | Суббота          | Saturday    |
| 12          | Saturday         | ERROR       |

