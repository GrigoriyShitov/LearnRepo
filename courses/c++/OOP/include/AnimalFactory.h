#ifndef ANIMAL_FACTORY_H
#define ANIMAL_FACTORY_H

#include <memory>
#include <string>
#include "Animal.h"

class AnimalFactory {
public:
    static std::unique_ptr<Animal> createAnimal(
        const std::string& type,
        const std::string& name,
        int age,
        const std::string& description
    );
};

#endif