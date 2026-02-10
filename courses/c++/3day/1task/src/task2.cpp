#include "../include/task2.h"

namespace task2
{
    void Do()
    {
        Student student;
        student.grade = 10;
        student.name = "Grigoriy";
        printf("Is excelent: %s\n", student.isExcellent() ? "true" : "false");
        student.printStudent();
    }
}
