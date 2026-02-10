#include "../include/task1.h"

namespace task1
{

    void Do()
    {
        Point p1 = {0, 0};
        Point p2 = {3, 4};
        std::cout << task1::distance(p1, p2) << std::endl;
    };

    int distance(Point p1, Point p2)
    {
        return std::sqrt(std::pow(p2.x - p1.x, 2) + std::pow(p2.y - p1.y, 2));
    
    }
}