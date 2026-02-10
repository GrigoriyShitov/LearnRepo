#ifndef task4_h
#define task4_h
#include <iostream>
#include <vector>
#include <list>
#include <set>
#include <algorithm>
#include <string>
#include <random>
namespace task4
{
    template <typename T>
    class Stack
    {
    private:
        std::vector<T> content;

    public:
        void push(const T &value)
        {
            content.push_back(value);
        }
        T pop()
        {
            if (content.empty())
            {
                throw std::runtime_error("Stack is empty");
            }
            T topElement = content.back();
            content.pop_back();
            return topElement;
        }
        const T &top() const
        {
            if (content.empty())
            {
                throw std::runtime_error("Stack is empty");
            }
            return content.back();
        }
        size_t size() const
        {
            return content.size();
        }

    };

    void task4() {
        Stack<int> intStack;
        try {
            intStack.push(10);
            intStack.push(20);
            std::cout << "Top element: " << intStack.top() << std::endl; // Вывод: 20
            std::cout << "Popped: " << intStack.pop() << std::endl;      // Вывод: 20
            std::cout << "Size: " << intStack.size() << std::endl;       // Вывод: 1
        } catch (const std::runtime_error& e) {
            std::cout << "Error: " << e.what() << std::endl;
        }
    }
}

#endif