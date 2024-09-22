import java.util.Scanner;

public class Main {
    public static void main(String[] args) {
        Scanner scan = new Scanner(System.in);
        int time= scan.nextInt();
        int hours = time / 3600;
        time %=3600;
        int minutes = time /60;
        time %=60;
        int seconds =time;
        System.out.printf("%d часов %d минут %d секунд", hours, minutes,seconds);
    }
}
