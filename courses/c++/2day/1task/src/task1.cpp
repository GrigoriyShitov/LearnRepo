#include <iostream>
namespace task1
{
    void Do()
    {
        int a, b;
        char *op;
        std::cout << "Enter 2 digits\n";
        std::cin >> a >> b;
        std::cout << "Enter operation\n";
        std::cin >> op;
        double result;
        switch (*op)
        {
        case '+':
            result = a + b;
            break;

        case '-':
            result = a - b;
            break;
        case '*':
            result = a * b;
            break;
        case '/':
            if (b == 0)
            {
                std::cout << "Error zero division";
                return;
            }
            result = a / b;
            break;
        default:
            break;
        }
        std::cout << "Result: " << result;
    }
}