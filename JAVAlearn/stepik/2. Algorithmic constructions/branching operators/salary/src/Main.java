import java.util.Scanner;

public class Main {
    public static void main(String[] args) {

        Scanner scanner = new Scanner(System.in);
        int hours = scanner.nextInt();
        float tarif = scanner.nextFloat();
        float sum = 0;
        if (hours < 0 || tarif < 0) {
            System.out.println("ERROR");
            return;
        }
        if (hours <= 20) {
            sum = hours * tarif;
        } else if (hours <= 40) {
            sum += 20 * tarif;
            sum += (hours - 20) * tarif * 1.5;
        } else {
            sum += 20 * tarif;
            sum += 20 * tarif * 1.5;
            hours -= 20;
            sum += (hours - 20) * tarif * 2;

        }

        System.out.printf("%.2f",sum);
    }
}