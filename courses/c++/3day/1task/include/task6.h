#ifndef TASK6_H
#define TASK6_H
#include "student.h"
#include "library.h"

namespace task6
{
    void Do();
    float generateData(Student *members, int size);
    std::string getRandomName();

    void sortToBaskets(Student *weakBasket, Student *strongBusket, Student *members, int size, float avg);
}

#endif