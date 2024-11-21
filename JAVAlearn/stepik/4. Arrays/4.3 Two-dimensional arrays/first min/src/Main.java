import java.util.Random;
import java.util.Scanner;

public class Main {
    public static void main(String[] args) {
        Scanner scanner = new Scanner(System.in);
        int[][] arr = new int[scanner.nextInt()][scanner.nextInt()];
        Random rand = new Random(scanner.nextInt());
        for (int i = 0; i < arr.length; i++) {
            for (int j = 0; j < arr[i].length; j++) {
                arr[i][j] = rand.nextInt(-10, 11);
                System.out.printf("%d\t", arr[i][j]);
            }
            System.out.println();
        }
        boolean flag = false;
        for (int i = 0; i < arr.length; i++) {
            for (int j = 0; j < arr[i].length; j++) {
                if (arr[i][j] < 0) {
                    System.out.println(j);
                    flag = true;
                    break;
                }
            }
            if (!flag) {
                System.out.println("NO");
            }
            flag = false;
        }
    }
}