import java.util.Scanner;

public class Main {
    public static void main(String[] args) {
        Scanner scan = new Scanner(System.in);
        int Rubles = scan.nextInt();
        int Kopeyki = scan.nextInt();

        if (Kopeyki > 100){
            Rubles += Kopeyki/100;
            Kopeyki %= 100;
        }
        System.out.printf("%d р. %d к.", Rubles,Kopeyki);
    }
}
