import java.util.Scanner;

public class Main {
    public static void main(String[] args) {
        Scanner scanner = new Scanner(System.in);

        int k = scanner.nextInt();
        for (int i=k; i>0; i--) {
            for (int j=0; j<i; j++) {
                System.out.print(i);
            }
            System.out.println();
        }
    }
}