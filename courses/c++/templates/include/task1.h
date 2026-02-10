#ifndef task1_h
#define task1_h

#include <iostream>
namespace task1
{
    template <typename T>
    void swap_values(T &a, T &b)
    {
        T temp = a;
        a = b;
        b = temp;
    };

    void task1();

    template <typename T>
    class SharedPtr
    {
    private:
        // Вспомогательная структура для хранения указателя и счётчика ссылок
        struct ControlBlock
        {
            T *ptr;       // Указатель на объект
            size_t count; // Счётчик ссылок

            ControlBlock(T *p) : ptr(p), count(1) {}
            ~ControlBlock()
            {
                if (ptr)
                {
                    delete ptr;
                }
            }
        };

        ControlBlock *control; // Указатель на блок управления

        // Вспомогательная функция для уменьшения счётчика и освобождения памяти
        void decrement()
        {
            if (control && --control->count == 0)
            {
                delete control;
                control = nullptr;
            }
        }

    public:
        // Конструктор из сырого указателя
        explicit SharedPtr(T *ptr = nullptr) : control(ptr ? new ControlBlock(ptr) : nullptr) {}

        // Конструктор копирования
        SharedPtr(const SharedPtr &other) : control(other.control)
        {
            if (control)
            {
                ++control->count;
            }
        }

        // Оператор присваивания
        SharedPtr &operator=(const SharedPtr &other)
        {
            if (this != &other)
            {
                decrement(); // Уменьшаем счётчик текущего объекта
                control = other.control;
                if (control)
                {
                    ++control->count;
                }
            }
            return *this;
        }

        // Деструктор
        ~SharedPtr()
        {
            decrement();
        }

        // Оператор разыменования (*)
        T &operator*() const
        {
            if (!control || !control->ptr)
            {
                throw std::runtime_error("Null pointer dereference");
            }
            return *(control->ptr);
        }

        // Оператор доступа к членам (->)
        T *operator->() const
        {
            if (!control || !control->ptr)
            {
                throw std::runtime_error("Null pointer access");
            }
            return control->ptr;
        }

        // Получить сырой указатель
        T *get() const
        {
            return control ? control->ptr : nullptr;
        }

        // Получить текущий счётчик ссылок
        size_t use_count() const
        {
            return control ? control->count : 0;
        }
    };

}

#endif