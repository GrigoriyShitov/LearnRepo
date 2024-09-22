import java.util.Scanner;

public class Main {
    public static void main(String[] args) {
        Scanner scan = new Scanner(System.in);
        String name = scan.nextLine();
        int age = scan.nextInt();
        float discount = scan.nextFloat();
        System.out.printf("С днем рождения, %s! Сегодня Вам %d!\n" +
                "По этому поводу дарим Вам скидку %.1f%%\n" +
                "на весь ассортимент нашего сайта!", name,age,discount);
    }
}
