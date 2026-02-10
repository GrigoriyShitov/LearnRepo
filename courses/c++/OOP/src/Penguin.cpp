#include "../include/Snake.h"
#include "../include/Lion.h"
#include "../include/Penguin.h"
#include <cstring>
#include <iostream>

Penguin::Penguin(const std::string &name, int age, const char *desc)
    : Animal(name, age)
{
    description = new char[strlen(desc) + 1];
    std::strcpy(description, desc);
}

Penguin::~Penguin()
{
    delete[] description;
}

void Penguin::feed()
{
    std::cout << name << " the Penguin is eating" << std::endl;
}

void Penguin::makeSound()
{
    std::cout << name << " the Penguin make sound" << std::endl;
}

void Penguin::move()
{
    std::cout << name << " the Penguin move" << std::endl;
}

void Penguin::interact(Animal &other)
{
    if (auto *lion = dynamic_cast<Lion *>(&other))
    {
        std::cout << name << " the Penguin interact with " << lion->getName() << " the Lion" << std::endl;
    }
    else if (auto *penguin = dynamic_cast<Penguin *>(&other))
    {
        std::cout << name << " the Penguin interact with " << penguin->getName() << " the Penguin" << std::endl;
    }
    else if (auto *snake = dynamic_cast<Snake *>(&other))
    {
        std::cout << name << " the Penguin interact with " << snake->getName() << " the Snake" << std::endl;
    }
    else
    {
        std::cout << name << " the Penguin interact with " << other.getName() << std::endl;
    }
}

void Penguin::saveToFile(std::ostream &os) const
{
    os << "Penguin," << name << "," << age << "," << description << std::endl;
}