import java.util.Scanner;

public class Main {
    public static void main(String[] args) {
        Scanner scanner = new Scanner(System.in);
        int a = scanner.nextInt();
        int b = scanner.nextInt();
        if (a <= 0 || b <= 0) {
            System.out.println("ERROR");
            return;
        }
        for (int i = 0; i < a; i++) {
            for (int j = 0; j < b; j++) {
                if (i != 0 && j != 0 && i != a - 1 && j != b - 1) {
                    System.out.print(" ");
                    continue;
                }
                System.out.print("*");
            }
            System.out.println();
        }
    }
}