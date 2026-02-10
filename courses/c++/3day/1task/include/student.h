#ifndef STUDENT_H
#define STUDENT_H
#include <string>
#include <iostream>
struct Student
{
    std::string name;
    int grade;
    bool isExcellent();
    void printStudent();
};
#endif