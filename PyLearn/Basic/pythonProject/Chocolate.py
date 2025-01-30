class Chocolate:
    def __init__(self, cost, topping=""):
        if type(cost) != float:
            raise TypeError("Неверный тип")
        if cost < 0:
            raise ValueError("Цена меньше либо равна нулю")
        self.cost = cost
        self.topping = topping

    def show_my_chocolate(self):
        if self.topping != "":
            print(f'Шоколадный батончик, добавка: {self.topping}')
        else:
            print(f'Обычный шоколадный батончик')
        return
    def __str__(self):
        return f'Шоколадный батончик, {self.cost} рублей'


choco1 = Chocolate(10.0)

choco2 = Chocolate(10.0,"milk")

choco1.show_my_chocolate()

choco2.show_my_chocolate()

print(choco1)

# chocoExeption= Chocolate(10, "")