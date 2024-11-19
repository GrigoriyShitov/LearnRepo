
# Task Description

Write a static method `printTriangle()` that prints a triangle to the console, as shown in the test examples. The parameters of the method are the width of the base and the character used to draw the triangle.

There should be no invisible spaces at the end of each line; the cursor should move immediately after the last character.

Also, write a `main()` method where the width of the triangle and the character are read from the console, and then call the `printTriangle()` method.

**Tip:** The `Scanner` class does not have a method for reading a single character. You can read the input as a string and then take the first character from that string:

```java
String str = scan.next();
char sim = str.charAt(0);
```

## Test Data

| Test Number | Input Data | Output Data                          |
|-------------|------------|--------------------------------------|
| 1           | 7 #        | #<br/>###<br/>#####<br/>#######      |
| 2           | 4 $        | $$<br/>$$$$                          |
| 3           | 3 *        | * <br/>***                           |
| 4           | 8 ?        | ??<br/>????<br/>??????<br/>????????? |

