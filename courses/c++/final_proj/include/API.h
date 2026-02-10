#ifndef API_H
#define API_H
#include <bits/stdc++.h>
#include "json.hpp"
#include "HTTP_client.h"
using json = nlohmann::json;

class API
{
private:
    json data;
    http_client http;
public:
    API() {};
    API(http_client http): http(http) {

    };
    std::vector<std::string> getAllPlanets();
    ~API(){};
};


std::vector<std::string> API::getAllPlanets()
{
    char* url = "https://swapi.info/api/planets";
    auto resp =  http.GET(url);
    json data = json::parse(resp);
    std::vector<std::string> outPlantes;
    outPlantes.reserve(data.size());

    for (auto planet : data){
        outPlantes.emplace_back(planet["name"]);
    } 
    
    return outPlantes;
}


#endif