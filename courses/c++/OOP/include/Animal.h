#ifndef task1_h
#define task1_h
#include <bits/stdc++.h>

class IFeedable
{
public:
    virtual void feed() = 0;
};

class Animal : public IFeedable
{
protected:
    std::string name;
    int age;

public:
    Animal(const std::string &name, int age);
    virtual ~Animal() = default;

    virtual void makeSound() = 0;
    virtual void move() = 0;
    virtual void interact(Animal &other) = 0;
    virtual void saveToFile(std::ostream &os) const = 0;

    std::string getName() const { return name; }
    int getAge() const { return age; }
};

#endif