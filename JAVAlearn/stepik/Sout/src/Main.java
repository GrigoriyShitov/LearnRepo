//TIP To <b>Run</b> code, press <shortcut actionId="Run"/> or
// click the <icon src="AllIcons.Actions.Execute"/> icon in the gutter.
public class Main {
    public static void main(String[] args) {
        String name = "Андрей";
        int age = 23;
        double discount = 15.5;
        System.out.printf("С Днем Рождения, %s!\n" +
                "Сегодня Вам %d!\n" +
                "По этому поводу Вам скидка %.1f%%\n" +
                "на весь ассортимент нашего сайта!",name,age,discount);
    }
}