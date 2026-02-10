#include "../include/task4.h"

namespace Task2
{
    void DoTask4()
    {
        int size;
        std::cout << "Enter size of array: ";
        std::cin >> size;
        int arr[size];
        fillArray(arr, size);
        printArray(arr, size);

        sortArray(arr, size);
        printArray(arr, size);
    }

    void sortArray(int arr[], int size)
    {
        int temp = 0;
        for (int i = 0; i < size - 1; i++)
        {
            for (int j = 0; j < size - i - 1; j++)
            {
                if (arr[j] > arr[j + 1])
                {
                    temp = arr[j];
                    arr[j] = arr[j + 1];
                    arr[j + 1] = temp;
                }
            }
        }
    }
}