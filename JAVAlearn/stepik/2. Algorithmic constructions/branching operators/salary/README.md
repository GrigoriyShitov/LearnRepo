### Task Description

The program should determine the salary of a company employee. The user inputs the number of hours worked (integer) and the hourly rate (float). The company pays the regular rate for the first 20 hours (inclusive), a time-and-a-half rate for hours worked from 20 (exclusive) to 40 hours (inclusive), and double the rate for hours worked over 40 hours.

For example, if an employee worked 42 hours and the hourly rate is 10.5 rubles, the calculation would be as follows:

- For the first 20 hours: $$20 \times 10.5 = 210$$ rubles.
- For the next 20 hours: $$20 \times 10.5 \times 1.5 = 315$$ rubles.
- For the last 2 hours: $$2 \times 10.5 \times 2 = 42$$ rubles.

Total salary: $$210 + 315 + 42 = 567$$ rubles.

The program should output the employee's salary rounded to two decimal places. If a negative value is entered for either the number of hours or the hourly rate, it should output "ERROR".

### Test Data

| Test Number | Input Data     | Output Data |
|-------------|----------------|-------------|
| 1           | 42 10.5        | 567.00      |
| 2           | -12 34.2       | ERROR       |
| 3           | 12 -3          | ERROR       |
| 4           | 15 20.5        | 307.50      |
| 5           | 25 14.3        | 393.25      |
| 6           | 60 13.4        | 1206.00     |
