#ifndef CLI_H
#define CLI_H
#include <bits/stdc++.h>
#include "curl/curl.h"
#include "HTTP_client.h"
#include "API.h"

class CLI
{
public:
    void StartCLI();
    void Menu();
    int ServeInput();
    void ShowPlanets(std::vector<std::string> &planets);
};

void CLI::StartCLI()
{
    while (true)
    {
        Menu();
        if (this->ServeInput() == 0)
        {
            break;
        }
    };
}

inline void CLI::ShowPlanets(std::vector<std::string> &planets)
{
    std::cout<<"Список планет:"<< std::endl;
    for (int i=0;i<planets.size();i++){
        std::cout<<i+1<<" "<< planets[i]<< std::endl;
    }
}

inline void CLI::Menu()
{
    std::cout << "Введите команду" << std::endl
              << "1. Вывести список планет" << std::endl
              << "2. Вывести информацию о планете" << std::endl
              << "0. Выйти" << std::endl;
}

int CLI::ServeInput()
{
    int choise;
    std::cin >> choise;
    CURL *curl = curl_easy_init();
    http_client http_client(curl);
    API swapi(http_client);
    switch (choise)
    {
    case 1:
        {
        auto planets = swapi.getAllPlanets();
        ShowPlanets(planets);
        break;
        }
    case 2:

        break;
    case 0:

        break;
    default:
        break;
    }
    return choise;
}

#endif