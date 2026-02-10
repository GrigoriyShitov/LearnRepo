#ifndef endofpipeline_h
#define endofpipeline_h

#include <bits/stdc++.h>
class EndOfPipeline : std::exception
{
public:
    EndOfPipeline(const char *msg) : msg(msg) {};
    const char *what();

private:
    const char *msg;
};
#endif