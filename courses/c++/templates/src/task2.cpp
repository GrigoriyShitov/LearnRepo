#include "../include/task2.h"

namespace task2
{

    void task2()
    {
        std::vector<int> numbers = {};
        std::list<int> list={};
        randFill(numbers);
        for (auto n : numbers)
        {
            std::cout<< n << " ";
        }
        std::cout<<std::endl<<std::endl;
        copy(numbers,list);
        // for (auto l: list){
        //     std::cout<< l << " ";
        // };
        list.sort();
        list.reverse();
        std::cout<<std::endl<<std::endl;
        for (auto l: list){
            std::cout<< l << " ";
        }
    }
    void randFill(std::vector<int> &numbers)
    {
        srand(time(NULL));
        for (int i = 0; i < 100; i++)
        {
            numbers.push_back(rand() % 100 + 1);
        }
    }
    void copy(std::vector<int> &numbers, std::list<int> &list)
    {
        std::set<int> set;
        for (auto num : numbers)
        {
            auto result = set.insert(num);
            if (result.second){
                list.push_back(num);
            }
        }
    }

}