#include "../include/Snake.h"
#include "../include/Lion.h"
#include "../include/Penguin.h"
#include <cstring>
#include <iostream>

Lion::Lion(const std::string &name, int age, const char *desc)
    : Animal(name, age)
{
    description = new char[strlen(desc) + 1];
    std::strcpy(description, desc);
}

Lion::~Lion()
{
    delete[] description;
}

void Lion::feed()
{
    std::cout << name << " the Lion is eating" << std::endl;
    ;
}

void Lion::makeSound()
{
    std::cout << name << " the Lion make sound" << std::endl;
    ;
}

void Lion::move()
{
    std::cout << name << " the Lion is moving" << std::endl;
    ;
}

void Lion::interact(Animal &other)
{
    if (auto *lion = dynamic_cast<Lion *>(&other))
    {
        std::cout << name << " the Lion interact with" << lion->getName() << " the Lion." << std::endl;
    }
    else if (auto *penguin = dynamic_cast<Penguin *>(&other))
    {
        std::cout << name << " the Lion interact with" << penguin->getName() << " the Penguin." << std::endl;
    }
    else if (auto *snake = dynamic_cast<Snake *>(&other))
    {
        std::cout << name << " the Lion interact with" << snake->getName() << " the Snake." << std::endl;
    }
    else
    {
        std::cout << name << " the Lion interact with" << other.getName() << std::endl;
    }
}

void Lion::saveToFile(std::ostream &os) const
{
    os << "Lion," << name << "," << age << "," << description << "\n";
}