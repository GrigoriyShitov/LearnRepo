#include "../include/AnimalFactory.h"
#include "../include/Lion.h"
#include "../include/Penguin.h"
#include "../include/Snake.h"

std::unique_ptr<Animal> AnimalFactory::createAnimal(
    const std::string& type,
    const std::string& name,
    int age,
    const std::string& description
) {
    if (type == "Lion") {
        return std::make_unique<Lion>(name, age, description.c_str());
    } else if (type == "Penguin") {
        return std::make_unique<Penguin>(name, age, description.c_str());
    } else if (type == "Snake") {
        return std::make_unique<Snake>(name, age, description.c_str());
    }
    return nullptr; 
}