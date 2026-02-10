#ifndef HTTP_http_client
#define HTTP_http_client
#include <bits/stdc++.h>
#include "curl/curl.h"

class http_client
{
private:
    CURL *curl;
    static size_t WriteCallback(void *data, size_t size, size_t nmemb, std::string *response)
    {
        response->append((char *)data, size * nmemb);
        return size * nmemb;
    }

public:
    http_client() {};
    http_client(CURL *curl_client) : curl(curl_client)
    {
        curl_easy_init();
        curl_easy_setopt(this->curl, CURLOPT_SSL_VERIFYPEER, 0L);
        curl_easy_setopt(this->curl, CURLOPT_SSL_VERIFYHOST, 0L);
    };
    std::string GET(char *path);
    ~http_client();
};

static size_t WriteCallback(void *contents, size_t size, size_t nmemb, std::string *s)
{
    size_t newLength = size * nmemb;
    s->append((char *)contents, newLength);
    return newLength;
}
std::string http_client::GET(char *path)
{
    size_t size;
    std::string response;
    curl_easy_setopt(curl, CURLOPT_URL, path);
    curl_easy_setopt(curl, CURLOPT_WRITEDATA, &response);
    std::cout<<response <<std::endl;
    auto res = curl_easy_perform(curl);
    if (res != CURLE_OK)
    {
        // throw std::runtime_error("Failed to method GET path: " + std::string(path));
    }
    return response;
}

http_client::~http_client()
{
    curl_easy_cleanup(curl);
}

#endif