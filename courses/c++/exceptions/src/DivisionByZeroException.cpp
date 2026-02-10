#include "../include/DivisionByZeroException.h"


const char *DivisionByZeroException::what()
{
    return this->msg;
}
