#ifndef PENGUIN_H
#define PENGUIN_H

#include "Animal.h"

class Penguin : public Animal {
private:
    char* description;

public:
    Penguin(const std::string& name, int age, const char* desc);
    ~Penguin() override;

    void feed() override;
    void makeSound() override;
    void move() override;
    void interact(Animal& other) override;
    void saveToFile(std::ostream& os) const override;
};

#endif