#include "../include/Environment.h"
#include "../include/AnimalFactory.h"
#include <fstream>
#include <sstream>
#include <iostream>

void Environment::addAnimal(std::unique_ptr<Animal> animal) {
    if (animal) {
        animals.push_back(std::move(animal));
    }
}

void Environment::printAll() const {
    for (const auto& animal : animals) {
        std::cout << "Animal: " << animal->getName() << "\n";
        animal->makeSound();
        animal->move();
        animal->feed();
        std::cout << "------------------------\n";
    }
}

void Environment::simulateInteraction() {
    for (size_t i = 0; i < animals.size(); ++i) {
        for (size_t j = i + 1; j < animals.size(); ++j) {
            animals[i]->interact(*animals[j]);
            animals[j]->interact(*animals[i]);
            std::cout << "------------------------\n";
        }
    }
}

void Environment::saveToFile(const std::string& filename) const {
    std::ofstream ofs(filename);
    if (!ofs.is_open()) {
        std::cerr << "Failed to open file " << filename << " for writing.\n";
        return;
    }
    for (const auto& animal : animals) {
        animal->saveToFile(ofs);
    }
    ofs.close();
}

void Environment::loadFromFile(const std::string& filename) {
    std::ifstream ifs(filename);
    if (!ifs.is_open()) {
        std::cerr << "Failed to open file " << filename << " for reading.\n";
        return;
    }

    animals.clear(); 
    std::string line;
    while (std::getline(ifs, line)) {
        std::stringstream ss(line);
        std::string type, name, description;
        int age;
        std::getline(ss, type, ',');
        std::getline(ss, name, ',');
        ss >> age;
        ss.ignore(1); 
        std::getline(ss, description);

        auto animal = AnimalFactory::createAnimal(type, name, age, description);
        if (animal) {
            animals.push_back(std::move(animal));
        }
    }
    ifs.close();
}