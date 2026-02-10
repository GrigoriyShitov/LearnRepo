#ifndef task2_h
#define task2_h
#include <iostream>
#include <vector>
#include <list>
#include <set>
namespace task2
{
    extern std::vector<int> numbers;
    extern std::list<int> list;
    void randFill(std::vector<int>& numbers);
    void copy(std::vector<int>& numbers, std::list<int>& list);
    void task2();
}

#endif