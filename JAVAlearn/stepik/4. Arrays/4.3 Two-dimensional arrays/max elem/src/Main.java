import java.util.Random;
import java.util.Scanner;

public class Main {
    public static void main(String[] args) {
        Scanner scanner = new Scanner(System.in);
        int[][] arr = new int[scanner.nextInt()][scanner.nextInt()];
        Random rand = new Random(scanner.nextInt());
        int max = Integer.MIN_VALUE, kol = 0;
        for (int i = 0; i < arr.length; i++) {
            for (int j = 0; j < arr[i].length; j++) {
                arr[i][j] = rand.nextInt(-5, 5);
                System.out.print(arr[i][j] + "\t");
                if (arr[i][j] == max) {
                    kol++;
                }
                if (arr[i][j] > max) {
                    max = arr[i][j];
                    kol = 1;
                }
            }
            System.out.println();
        }
        System.out.printf("%d %d", max, kol);
    }

}