import java.util.Random;
import java.util.Scanner;

public class Main {
    public static void main(String[] args) {
        Scanner scanner = new Scanner(System.in);
        int n = scanner.nextInt();
        int k = scanner.nextInt();
        int[][] arr = new int[n][k];
        Random rand = new Random(scanner.nextInt());
        for (int i = 0; i < arr.length; i++) {
            for (int j = 0; j < arr[i].length; j++) {
                arr[i][j] = rand.nextInt(-10, 11);
                System.out.printf("%d\t", arr[i][j]);
            }
            System.out.println();
        }
        System.out.println();
        int sum = 0;

        for (int j = 0; j < k; j++) {
            for (int i = 0; i < n; i++) {
                if (arr[i][j] > 0) {
                    sum += arr[i][j];
                }
            }
            System.out.printf("%d ", sum);
            sum = 0;
        }
    }
}