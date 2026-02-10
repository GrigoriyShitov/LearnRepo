#include "../include/task2.hpp"

namespace Task2
{
    void fillArray(int arr[], int size)
    {
        srand(time(0));
        for (int i = 0; i < size; i++)
        {
            arr[i] = rand() % 100;
        }
        return;
    }
    void printArray(int arr[], int size)
    {
        std::cout << "Array: ";
        for (int i = 0; i < size; i++)
        {
            std::cout << arr[i] << " ";
        }
        std::cout << "\n";
        return;
    }

    void findStats(int arr[], int size, int &min, int &max, double &avg)
    {

        int sum = 0;
        for (int i = 0; i < size; i++)
        {
            if (arr[i] > max)
            {
                max = arr[i];
            }
            if (arr[i] < min)
            {
                min = arr[i];
            }
            sum += arr[i];
        }
        avg = sum / size;
    }

    void Do()
    {
        int size;
        std::cout << "Enter size (less than 20)";
        std::cin >> size;
        if (size >= 20)
        {
            std::cout << "Wrong size";
            return;
        }

        int arr[size];
        fillArray(arr, size);
        printArray(arr, size);

        int max = INT_MIN;
        int min = INT_MAX;
        double avg;
        findStats(arr, size, min, max, avg);
        std::cout << "Stats:\nMin: " << min << " Max: " << max << " Avg: " << avg;
        return;
    }
}