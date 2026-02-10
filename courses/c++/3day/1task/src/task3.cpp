#include "../include/task3.h"
namespace task3
{

    Student arr[5] = { {"Alice", 8}, {"Bob", 9}, {"Charlie", 7}, {"David", 10}, {"Eve", 6} };

    void Do()
    {
        std::cout << "The best student is " << findBestStudent().name << ", with grage:" << findBestStudent().grade << std::endl;
        printStudent();
        std::cout << std::endl;
        printStudent(true);
    }

    Student findBestStudent()
    {
        Student max = {"", 0};
        for (Student stud : arr)
        {
            if (stud.grade > max.grade)
            {
                max.grade = stud.grade;
                max.name = stud.name;
            }
        }
        return max;
    }
    void printStudent(bool bEndl)
    {
        for (Student stud : arr)
        {
            std::cout << stud.name << " : " << stud.grade << "  ";
            bEndl ? std::cout << "\n" : std::cout << "";
        }
    }
    void printStudent()
    {
        printStudent(false);
    }
}
