#include "../include/task1.h"

namespace task1
{
    double getSquare()
    {
        double digit;
        std::cin >> digit;
        if (digit < 0)
        {
            throw std::invalid_argument("less than zero");
        }
        else if (digit > 100)
        {
            throw std::out_of_range("out of range");
        }else if (digit ==0){
            throw std::logic_error("exit from task1");
        }
        return std::pow(digit,2);
    };

    void Do() {
        while(true){
            try{
                std::cout<< getSquare()<<std::endl;

            }catch(std::invalid_argument& msg){
                std::cout<<"Your input in less than zero. Error: "<< msg.what() <<std::endl;
            }
            catch(std::out_of_range& msg){
                std::cout<<"Your input is more than max accessed. Error: "<< msg.what() <<std::endl;
            }catch(...){
                break;
            }
        };
    };
}