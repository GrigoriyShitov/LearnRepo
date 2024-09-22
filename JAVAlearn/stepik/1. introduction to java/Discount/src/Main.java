import java.util.Scanner;

public class Main {
    public static void main(String[] args) {
        Scanner scan = new Scanner(System.in);
        int sum = scan.nextInt() * 100;
        sum+= scan.nextInt();
        int discount = scan.nextInt();

        int finalsum = sum- sum*discount/100;
        System.out.printf("%02d руб. %02d коп.", finalsum/100, finalsum%100);

    }
}
