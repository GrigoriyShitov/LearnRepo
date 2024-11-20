import java.util.Random;
import java.util.Scanner;

public class Main {
    public static void main(String[] args) {


        Scanner scanner = new Scanner(System.in);
        int n = scanner.nextInt();
        Random rand = new Random(scanner.nextInt());
        int min = 0, max = 5;
        double avg = 0;
        double[] arr = new double[n];
        for (int i = 0; i < n; i++) {
            arr[i] = rand.nextDouble(min, max);
            System.out.printf("%.2f ", arr[i]);
            avg += arr[i];
        }
        avg /= n;
        System.out.printf("\n%.2f\n", avg);
        for (int i = 0; i < n; i++) {
            if (arr[i] > avg) {
                arr[i] = avg;
            }
            System.out.printf("%.2f ", arr[i]);
        }
    }
}