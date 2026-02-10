#include "../include/task3.h"

namespace Task3
{
    void Do()
    {
        int size;
        std::cout << "Enter size of array: ";
        std::cin >> size;
        int arr[size];
        Task2::fillArray(arr, size);
        Task2::printArray(arr, size);

        int x;
        std::cout << "Enter x: ";
        std::cin >> x;

        countAndSum(arr, size, x);
    }
    void countAndSum(int arr[], int size, int x)
    {
        int count = 0, sum = 0;
        for (int i = 0; i < size; i++)
        {
            if (arr[i] < x)
            {
                count++;
                sum += arr[i];
            }
        }
        std::cout << "Count: " << count << "\nSum: " << sum;
    }
}
