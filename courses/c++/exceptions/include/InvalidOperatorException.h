#ifndef inv_op
#define inv_op
#include <bits/stdc++.h>
class InvalidOperatorException : std::exception
{
public:
    InvalidOperatorException(const char *msg) : msg(msg) {};
    const char *what();

private:
    const char *msg;
};

#endif