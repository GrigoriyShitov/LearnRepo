#include "../include/task5.h"

namespace task5
{

    void task5()
    {
        std::map<std::string, int> Studs;
        Studs["Alice"] = 85;
        Studs["Bob"] = 92;
        Studs["Charlie"] = 78;
        Studs["Diana"] = 95;
        Studs["Eve"] = 88;
        auto a = std::max_element(Studs.begin(), Studs.end(), [](std::pair<std::string, int> a, std::pair<std::string, int> b)
                                  { return a.second < b.second; });
        std::cout << a->second << std::endl;
        double avg = static_cast<double>(std::accumulate(Studs.begin(), Studs.end(), 0, [](int sum, const auto &pair)
                                                             {
                                                                 return sum + pair.second; // Суммируем баллы
                                                             })) /
                         Studs.size();
        std::cout << avg << std::endl;

        for(auto s: Studs){
            if (s.second > avg){
                std::cout<< s.second<<" ";
            }
        }
    }

}