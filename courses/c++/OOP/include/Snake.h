#ifndef SNAKE_H
#define SNAKE_H

#include "Animal.h"

class Snake : public Animal {
private:
    char* description;

public:
    Snake(const std::string& name, int age, const char* desc);
    ~Snake() override;

    void feed() override;
    void makeSound() override;
    void move() override;
    void interact(Animal& other) override;
    void saveToFile(std::ostream& os) const override;
};

#endif