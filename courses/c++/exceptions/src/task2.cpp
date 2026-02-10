#include "../include/task2.h"

namespace task2
{
    double calculate(int a, int b, char *op)
    {
        if (op == "+")
        {
            return a + b;
        }
        else if (op == "-")
        {
            return a - b;
        }
        else if (op == "*")
        {
            return a * b;
        }
        else if (op == "/")
        {
            if (b == 0)
            {
                throw DivisionByZeroException("zero div");
            }
            return a / b;
        }
        else
        {
            throw InvalidOperatorException("invalid operation");
        }
    };

    void Do()
    {
        for (int i = 0; i < 5; i++)
        {

            try
            {
                switch (i)
                {
                case 0:
                    std::cout << calculate(5, 6, "+") << std::endl;
                    break;

                case 1:
                    std::cout << calculate(5, 6, "*") << std::endl;
                    break;

                case 2:
                    std::cout << calculate(5, 0, "/") << std::endl;
                    break;

                case 3:
                    std::cout << calculate(5, 0, "++") << std::endl;
                    break;
                default:
                    break;
                }
            }
            catch (DivisionByZeroException &z)
            {
                std::cout << "Zero div error occured while math: " << z.what() << std::endl;
            }
            catch (InvalidOperatorException &i)
            {
                std::cout << "Invalid op error occured while math: " << i.what() << std::endl;
            }
        }
    };
}