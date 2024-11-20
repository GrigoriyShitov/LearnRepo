import java.util.Random;
import java.util.Scanner;

public class Main {
    public static void main(String[] args) {


        Scanner scanner = new Scanner(System.in);
        int n = scanner.nextInt();
        Random rand = new Random(scanner.nextInt());
        int min = -5, max = 5, sum = 0;
        double mult = 1;
        int[] arr = new int[n];
        for (int i = 0; i < n; i++) {
            arr[i] = rand.nextInt(min, max + 1);
            System.out.printf("%d ", arr[i]);
            if (arr[i] >= 0) {
                sum += arr[i];
            } else {
                mult *= arr[i];
            }
        }
        System.out.printf("\n%d %.1f ", sum, mult);
    }
}