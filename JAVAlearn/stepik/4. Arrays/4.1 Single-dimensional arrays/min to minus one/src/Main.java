import java.util.Scanner;

public class Main {
    public static void main(String[] args) {
        Scanner scanner = new Scanner(System.in);
        int n = scanner.nextInt();
        int[] arr = new int[n];
        int imin = 0;
        for (int i = 0; i < n; i++) {
            arr[i] = scanner.nextInt();
            if (arr[i] <= arr[imin]) {
                imin = i;
            }
        }
        arr[imin] = -1;
        for (int i = 0; i < n; i++) {
            System.out.print(arr[i] + " ");
        }
    }
}