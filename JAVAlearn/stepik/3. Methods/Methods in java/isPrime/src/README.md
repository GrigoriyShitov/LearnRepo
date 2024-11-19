# Task Description

Write a static method `isPrime()` that returns `true` if the argument is a prime number, and `false` otherwise.

A prime number is defined as a positive integer that has no divisors other than 1 and itself.

Negative numbers cannot be prime (therefore, the function should return `false`). The numbers 0 and 1 are also not considered prime.

The usage of the `isPrime()` method is demonstrated in the `main` method. This code should not be modified!

## Test Data

| Test Number | Input Data         | Output Data                  |
|-------------|---------------------|------------------------------|
| 1           | 1 13 23 67 10       | false true true true false    |
| 2           | -5 6 9 29 5         | false false false true true    |
| 3           | 127 69 -3 0 7       | true false false false true    |
| 4           | 1 0 2 3 17          | false false true true true     |
