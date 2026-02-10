#include "../include/Environment.h"
#include "../include/AnimalFactory.h"
#include <iostream>

int main() {
    Environment zoo;

    // Добавляем животных через фабрику
    zoo.addAnimal(AnimalFactory::createAnimal("Lion", "Leo", 5, "Majestic king of the jungle"));
    zoo.addAnimal(AnimalFactory::createAnimal("Penguin", "Pingu", 3, "Cute waddling bird"));
    zoo.addAnimal(AnimalFactory::createAnimal("Snake", "Slyther", 2, "Venomous slithering reptile"));

    // Выводим информацию о всех животных
    std::cout << "=== Zoo Information ===\n";
    zoo.printAll();

    // Симулируем взаимодействие
    std::cout << "\n=== Animal Interactions ===\n";
    zoo.simulateInteraction();

    // Сохраняем в файл
    std::cout << "\n=== Saving to file ===\n";
    zoo.saveToFile("build/zoo.txt");

    // Очищаем зоопарк и загружаем из файла
    std::cout << "\n=== Loading from file ===\n";
    zoo.loadFromFile("build/zoo.txt");
    zoo.printAll();

    return 0;
}