#ifndef div_zero
#define div_zero
#include <bits/stdc++.h>
class DivisionByZeroException : std::exception
{
public:
    DivisionByZeroException(const char *msg) : msg(msg) {};
    const char *what();

private:
    const char *msg;
};

#endif