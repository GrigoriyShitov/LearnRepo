#include "../include/task1.h"

namespace task1
{
    struct Test
    {
        int value;
        Test(int v) : value(v)
        {
            std::cout << "Test created: " << value << "\n";
        }
        ~Test()
        {
            std::cout << "Test destroyed: " << value << "\n";
        }
    };
    void task1()
    {
        int a = 5;
        int b = 0;
        std::cout << a << " " << b << std::endl;
        swap_values(a, b);
        std::cout << a << " " << b << std::endl;

        std::string c = "test";
        std::string d = "proverkra";
        std::cout << c << " " << d << std::endl;
        swap_values(c, d);
        std::cout << c << " " << d << std::endl;

        // Доп задание
        // Создание SharedPtr из сырого указателя
        SharedPtr<Test> ptr1(new Test(42));
        std::cout << "ptr1 use_count: " << ptr1.use_count() << "\n"; // 1
        std::cout << "ptr1 value: " << ptr1->value << "\n";          // 42

        // Копирование
        {
            SharedPtr<Test> ptr2 = ptr1;
            std::cout << "ptr2 use_count: " << ptr2.use_count() << "\n";     // 2
            std::cout << "ptr1 use_count: " << ptr1.use_count() << "\n";     // 2
            *ptr2 = Test(100);                                               // Изменяем объект
            std::cout << "ptr1 value after change: " << ptr1->value << "\n"; // 100
        }

        std::cout << "After ptr2 destroyed, ptr1 use_count: " << ptr1.use_count() << "\n"; // 1

        // Присваивание
        SharedPtr<Test> ptr3;
        ptr3 = ptr1;
        std::cout << "ptr3 use_count: " << ptr3.use_count() << "\n"; // 2

        // Пустой указатель
        SharedPtr<Test> ptr4;
        std::cout << "ptr4 use_count: " << ptr4.use_count() << "\n"; // 0

        return; // Все указатели уничтожаются, объект освобождается
    }

}