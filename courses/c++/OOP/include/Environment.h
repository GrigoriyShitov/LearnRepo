#ifndef ENVIRONMENT_H
#define ENVIRONMENT_H

#include <vector>
#include <memory>
#include <string>
#include "Animal.h"

class Environment {
private:
    std::vector<std::unique_ptr<Animal>> animals;

public:
    void addAnimal(std::unique_ptr<Animal> animal);
    void printAll() const;
    void simulateInteraction();
    void saveToFile(const std::string& filename) const;
    void loadFromFile(const std::string& filename);
};

#endif