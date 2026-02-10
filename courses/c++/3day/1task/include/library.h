#ifndef LIBRARY_H
#define LIBRARY_H
#include <iostream>
#include <cstring>
#include <random>
#include <string>
struct Book{
    int id=0;
    char* author= "";
    char* title="";
    int year;
    bool isAvailable;
    Book() : id(0), author(nullptr), title(nullptr), year(0), isAvailable(false) {}
};

struct Library
{
    Book arrLib[1000];
    bool addBook(Book book);
    Book searchByTitle(char* title);
    bool takeBook(char* title);
    void returnBook(char* title);
    void printBookList();
    Book searchByAuthor(char* author);
    int getLastNonEmptyIndex();
};


#endif