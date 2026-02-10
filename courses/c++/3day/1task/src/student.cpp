#include "../include/student.h"

bool Student::isExcellent()
{
    if (grade >= 9)
    {
        return true;
    }
    return false;
}

void Student::printStudent()
{
    std::cout << "Student name: " << name << "\nGrage: " << grade << std::endl;
    ;
}
