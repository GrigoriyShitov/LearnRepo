#ifndef task3_h
#define task3_h
#include <bits/stdc++.h>
#include "NetworkException.h"
namespace task3
{

    void fetchData();
    void Do();


    class Data
    {
    private:
        void cleanup() noexcept;

    public:
        int *resource;
        Data();
        ~Data() noexcept
        {
            cleanup();
        };
    };

    
}

#endif