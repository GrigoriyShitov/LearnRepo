#ifndef LION_H
#define LION_H
#include "Animal.h"

class Lion : public Animal {
private:
    char* description;

public:
    Lion(const std::string& name, int age, const char* desc);
    ~Lion() override;

    void feed() override;
    void makeSound() override;
    void move() override;
    void interact(Animal& other) override;
    void saveToFile(std::ostream& os) const override;
};

#endif