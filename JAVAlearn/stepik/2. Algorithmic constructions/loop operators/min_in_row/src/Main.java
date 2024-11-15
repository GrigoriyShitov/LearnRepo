import java.util.Scanner;

public class Main {
    public static void main(String[] args) {

        Scanner scanner = new Scanner(System.in);
        int n = scanner.nextInt();
        int min =Integer.MAX_VALUE;
        for (int i = 0; i < n; i++) {
            int a = scanner.nextInt();
            if (a < min) {
                min = a;
            }
        }
        System.out.println(min);
    }
}