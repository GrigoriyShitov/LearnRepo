#include "../include/task3.h"

namespace task3
{

    void task3()
    {
        std::vector<std::string> words = {"apple", "banana", "cherry", "date", "fig"};
        std::cout << *std::max_element(words.begin(), words.end(), [](std::string &a, std::string &b)
                                       { return a.length() < b.length(); })
                  << std::endl;

        std::sort(words.begin(), words.end());
        for (auto w : words)
        {
            std::cout << w << " ";
        }

        std::cout << std::count_if(words.begin(), words.end(), [](std::string &a)
                                   { return (a.compare(0, 1, "a") == 0); })
                  << std::endl;

        std::random_device rd;
        std::mt19937 g(rd());
        std::shuffle(words.begin(), words.end(), g);
        for (auto w : words)
        {
            std::cout << w << " ";
        }
    }

}