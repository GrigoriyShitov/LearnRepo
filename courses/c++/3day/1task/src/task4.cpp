#include "../include/task4.h"
namespace task4
{

    void Do()
    {
        Library lib;
        Book newBook;
        newBook.author = "testAuthor";
        newBook.id = 0;
        newBook.isAvailable = true;
        newBook.title = "testTitle";
        newBook.year = 2025;
        lib.addBook(newBook);
        lib.addBook(newBook);
        lib.addBook(newBook);
        lib.printBookList();
        lib.searchByAuthor("testAutor");
        lib.searchByTitle("testTitle");
        lib.takeBook("testTitle");
        lib.returnBook("testTitle");
    };

}
