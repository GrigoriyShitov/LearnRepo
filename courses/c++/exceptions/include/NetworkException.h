#ifndef NetworkException_h
#define NetworkException_h
#include <bits/stdc++.h>

class NetworkException : public std::exception
{
public:
    NetworkException(const char *msg) : msg(msg) {};
    const char *what(){
        return this->msg;
    };
private:
    const char *msg;
};

class TimeoutException : public NetworkException
{
public:
    TimeoutException(const char *msg) : NetworkException(msg) {};


};

class ProtocolException : public TimeoutException
{
public:
    ProtocolException(const char *msg) : TimeoutException(msg) {};

};

#endif
