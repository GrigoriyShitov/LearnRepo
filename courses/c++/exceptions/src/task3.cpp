#include "../include/task3.h"

namespace task3
{
    void fetchData()
    {

        int i = rand() % 4 + 1;
        // i = 4;
        std::cout<< i;
        switch (i)
        {
        case 1:
            throw NetworkException("network");
            break;
        case 2:
            throw TimeoutException("timeout");
            break;
        case 3:
            throw("std::exception");
            break;
        case 4:
            throw ProtocolException("protocol");
            break;
        default:
            break;
        }
    };

    void Data::cleanup() noexcept
    {
        try
        {
            
            int i = rand() % 5 + 1;
            // i=1;
            if (i == 1)
            {
                throw;
            }
            delete this->resource;
        }
        catch (...)
        {
            std::cout << "catch exception in cleanup" << std::endl;
            delete this->resource;
        }
    }
    Data::Data()
    {
        resource = new int;
        try
        {
            fetchData();
        }
        catch (TimeoutException &t)
        {
            std::cout << "TimeoutException: " << t.what() << std::endl;
            cleanup();
            throw t;
        }
        catch (NetworkException &n)
        {
            std::cout << "NetworkException: " << n.what() << std::endl;
            cleanup();
            throw n;
        }
        catch (std::exception &e)
        {
            std::cout << "std::exception: " << e.what() << std::endl;
            cleanup();
            throw e;
        }
    }

    void Do()
    {
        {
            srand(time(0));
            try
            {
                Data test1;
            }catch(...){
                
            }
        }
        {
            srand(1);
            try
            {
                Data test2;
            }catch(...){
                
            }
        }
    };
}