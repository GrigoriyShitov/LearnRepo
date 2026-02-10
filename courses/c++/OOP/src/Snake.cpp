#include "../include/Snake.h"
#include "../include/Lion.h"
#include "../include/Penguin.h"
#include <cstring>
#include <iostream>

Snake::Snake(const std::string &name, int age, const char *desc)
    : Animal(name, age)
{
    description = new char[strlen(desc) + 1];
    std::strcpy(description, desc);
}

Snake::~Snake()
{
    delete[] description;
}

void Snake::feed()
{
    std::cout << name << " the Snake is eating" << std::endl;
}

void Snake::makeSound()
{
    std::cout << name << " the Snake makes sound" << std::endl;
    ;
}

void Snake::move()
{
    std::cout << name << " the Snake is moving" << std::endl;
    ;
}

void Snake::interact(Animal &other)
{

    if (auto *lion = dynamic_cast<Lion *>(&other))
    {
        std::cout << name << " the Snake interact with " << lion->getName() << " the Lion" << std::endl;
    }
    else if (auto *penguin = dynamic_cast<Penguin *>(&other))
    {
        std::cout << name << " the Snake interact with " << penguin->getName() << " the Penguin" << std::endl;
    }
    else if (auto *snake = dynamic_cast<Snake *>(&other))
    {
        std::cout << name << " the Snake interact with " << snake->getName() << " the Snake" << std::endl;
    }
    else
    {
        std::cout << name << " the Snake interact with " << other.getName() << std::endl;
    }
}

void Snake::saveToFile(std::ostream &os) const
{
    os << "Snake," << name << "," << age << "," << description << std::endl;
    ;
}