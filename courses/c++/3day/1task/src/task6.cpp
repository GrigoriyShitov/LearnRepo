#include "../include/task6.h"

namespace task6
{
    void Do()
    {
        int num;

        do
        {
            std::cout << "Pls input odd nubmer of members: ";
            std::cin >> num;

        } while (num % 2 != 0);
        Student members[num];

        float avg = generateData(members, num);
        Student weakBasket[num / 2];
        Student strongBasket[num / 2];

        sortToBaskets(weakBasket, strongBasket, members, num, avg);
        std::cout<<"\n Strong Basket:\n";
        for(int i=0; i<num/2;i++){
            std::cout<< strongBasket[i].name<<" "<<strongBasket[i].grade<< std::endl;
        }
    }
    void sortToBaskets(Student *weakBasket, Student *strongBasket, Student *members, int size, float avg)
    {
        Student* ptrMember = members;
        Student* ptrWeak= weakBasket;
        Student* ptrStrong= strongBasket;
        
        for(int i = 0; i < size; i++)
        {
            if(ptrMember->grade<avg){
                ptrWeak->grade = ptrMember->grade;
                ptrWeak->name= ptrMember->name;
                ptrWeak++;
            }else{
                ptrStrong->grade = ptrMember->grade;
                ptrStrong->name= ptrMember->name;
                ptrStrong++;
            }
            ptrMember++;
        }
    }
    float generateData(Student *members, int size)
    {
        srand(time(NULL));
        Student* ptr= members;
        float avg;
        int sum;
        for (int i = 0; i < size; i++)
        {
            ptr->grade = rand() % 10 + 1;
            sum += ptr->grade;
            ptr->name = getRandomName();
            std::cout << ptr->name << " " << ptr->grade << std::endl;
            ptr++;
        }
        avg = sum / size;
        return avg;
    }

    // Определяем enum с большим количеством имен
    enum class Name
    {
        Alice,
        Bob,
        Charlie,
        David,
        Eve,
        Frank,
        Grace,
        Henry,
        Isabella,
        Jack,
        Kelly,
        Liam,
        Mia,
        Noah,
        Olivia,
        Peter,
        Quinn,
        Rose,
        Sam,
        Tessa,
        Uma,
        Victor,
        Wendy,
        Xavier,
        Yvonne,
        Zach,
        Amy,
        Ben,
        Clara,
        Daniel,
        // Добавим еще имена для демонстрации
        Emma,
        Finn,
        Hannah,
        Ian,
        Julia,
        Kevin,
        Lily,
        Mason,
        Natalie,
        Oliver,
        Penelope,
        Quentin,
        Rachel,
        Steven,
        Tracy,
        Ulysses,
        Violet,
        William,
        Xena,
        Yara,
        // Количество элементов в enum
        _COUNT // Специальный элемент для подсчета
    };

    // Функция для преобразования Name в string
    std::string nameToString(Name name)
    {
        switch (name)
        {
        case Name::Alice:
            return "Alice";
        case Name::Bob:
            return "Bob";
        case Name::Charlie:
            return "Charlie";
        case Name::David:
            return "David";
        case Name::Eve:
            return "Eve";
        case Name::Frank:
            return "Frank";
        case Name::Grace:
            return "Grace";
        case Name::Henry:
            return "Henry";
        case Name::Isabella:
            return "Isabella";
        case Name::Jack:
            return "Jack";
        case Name::Kelly:
            return "Kelly";
        case Name::Liam:
            return "Liam";
        case Name::Mia:
            return "Mia";
        case Name::Noah:
            return "Noah";
        case Name::Olivia:
            return "Olivia";
        case Name::Peter:
            return "Peter";
        case Name::Quinn:
            return "Quinn";
        case Name::Rose:
            return "Rose";
        case Name::Sam:
            return "Sam";
        case Name::Tessa:
            return "Tessa";
        case Name::Uma:
            return "Uma";
        case Name::Victor:
            return "Victor";
        case Name::Wendy:
            return "Wendy";
        case Name::Xavier:
            return "Xavier";
        case Name::Yvonne:
            return "Yvonne";
        case Name::Zach:
            return "Zach";
        case Name::Amy:
            return "Amy";
        case Name::Ben:
            return "Ben";
        case Name::Clara:
            return "Clara";
        case Name::Daniel:
            return "Daniel";
        case Name::Emma:
            return "Emma";
        case Name::Finn:
            return "Finn";
        case Name::Hannah:
            return "Hannah";
        case Name::Ian:
            return "Ian";
        case Name::Julia:
            return "Julia";
        case Name::Kevin:
            return "Kevin";
        case Name::Lily:
            return "Lily";
        case Name::Mason:
            return "Mason";
        case Name::Natalie:
            return "Natalie";
        case Name::Oliver:
            return "Oliver";
        case Name::Penelope:
            return "Penelope";
        case Name::Quentin:
            return "Quentin";
        case Name::Rachel:
            return "Rachel";
        case Name::Steven:
            return "Steven";
        case Name::Tracy:
            return "Tracy";
        case Name::Ulysses:
            return "Ulysses";
        case Name::Violet:
            return "Violet";
        case Name::William:
            return "William";
        case Name::Xena:
            return "Xena";
        case Name::Yara:
            return "Yara";
        default:
            throw std::invalid_argument("Unknown name");
        }
    }

    // Функция для получения случайного имени
    std::string getRandomName()
    {
        // Инициализация генератора случайных чисел
        static std::random_device rd;
        static std::mt19937 gen(rd());
        // Диапазон от 0 до Name::_COUNT - 1
        std::uniform_int_distribution<> dis(0, static_cast<int>(Name::_COUNT) - 1);

        // Получаем случайное значение enum
        Name randomName = static_cast<Name>(dis(gen));
        // Преобразуем в строку
        return nameToString(randomName);
    }
}
