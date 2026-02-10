#include "../include/InvalidOperatorException.h"

const char *InvalidOperatorException::what()
{
    return this->msg;
}
