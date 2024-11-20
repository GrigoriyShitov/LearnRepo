import java.util.Arrays;
import java.util.Random;
import java.util.Scanner;

public class Main {
    public static void main(String[] args) {
        Scanner scanner = new Scanner(System.in);
        int n = scanner.nextInt();
        double[] arr = new double[n];
        Random rand = new Random(scanner.nextInt());
        int imin = 0, imax = 0;
        for (int i = 0; i < arr.length; i++) {
            arr[i] = rand.nextDouble(0, 2);
            if (arr[i] < arr[imin]) {
                imin = i;
            }
            if (arr[i] > arr[imax]) {
                imax = i;
            }
        }
        System.out.println(Arrays.toString(arr));
        System.out.println(arr[imin] + arr[imax]);
    }
}