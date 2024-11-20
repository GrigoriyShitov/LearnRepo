import java.util.Random;
import java.util.Scanner;

public class Main {
    public static void main(String[] args) {
        Scanner scanner = new Scanner(System.in);
        int n = scanner.nextInt();
        Random rand  = new Random(scanner.nextInt());
        int[] arr = new int[n];
        int imin=0,imax=0;
        for (int i = 0; i < arr.length; i++) {
            arr[i] = rand.nextInt(-5,16);
            if (arr[i] <0) {
                imin = i;

            }
            if (arr[i] > arr[imax]) {
                imax = i;

            }
            System.out.printf("%d ", arr[i]);
        }
        System.out.println();
        if (arr[imin] <0){
            int temp = arr[imax];
            arr[imax] = arr[imin];
            arr[imin] = temp;
        }
        for (int i:arr){
            System.out.print(i+" ");
        }
    }
}