#include "../include/library.h"

bool Library::addBook(Book book)
{

    int last = getLastNonEmptyIndex();
    if (last == -1)
    {
        return false;
    }
    arrLib[last] = book;
    return true;
}
Book Library::searchByTitle(char *title)
{
    for (Book book : arrLib)
    {
        if (book.title == title)
        {
            return book;
        }
    }
    return Book();
}
bool Library::takeBook(char *title)
{
    for (Book book : arrLib)
    {
        if (book.title == title)
        {
            if (book.isAvailable)
            {
                book.isAvailable = false;
                std::cout << "You taked the book with title: " << book.title << std::endl;
                return true;
            }
            return false;
        }
    }
    return false;
}

void Library::returnBook(char *title)
{
    for (Book book : arrLib)
    {
        if (book.title == title)
        {
            book.isAvailable = true;
            std::cout << "You returned the book with title: " << book.title << std::endl;
            return;
        }
    }
    return;
}

void Library::printBookList()
{
    std::cout << "Here the list of the books: " << std::endl;

    for (Book book : arrLib)
    {
        if (book.author== nullptr){
            return;
        }
        std::cout << book.id << " " << book.author << " " << book.title << " " << book.year << " " << (book.isAvailable ? "available\n" : "not available\n");
    }
    std::cout<<"a che takoe";
    return;
}

Book Library::searchByAuthor(char *author)
{
    for (Book book : arrLib)
    {
        if (book.author == author)
        {
            std::cout << book.id << " " << book.author << " " << book.title << " " << book.year << " " << book.isAvailable << std::endl;
        }
    }
    return Book();
}

int Library::getLastNonEmptyIndex()
{
    for (int i = 0; i < 1000; i++)
    {
        if (arrLib[i].author == nullptr)
        {
            return i;
        }
    }
    return -1; // если все элементы пустые
}