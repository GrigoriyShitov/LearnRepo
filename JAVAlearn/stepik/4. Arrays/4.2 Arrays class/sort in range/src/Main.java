import java.util.Arrays;
import java.util.Random;
import java.util.Scanner;

public class Main {
    public static void main(String[] args) {
        Scanner scanner = new Scanner(System.in);
        int n = scanner.nextInt();
        Random rand = new Random(scanner.nextInt());
        int i1 = scanner.nextInt(), i2 = scanner.nextInt();
        int[] arr = new int[n];
        for (int i = 0; i < arr.length; i++) {
            arr[i] = rand.nextInt(10, 21);
        }
        System.out.println(Arrays.toString(arr));

        Arrays.sort(arr, i1, i2 + 1);
        System.out.println(Arrays.toString(arr));
    }
}