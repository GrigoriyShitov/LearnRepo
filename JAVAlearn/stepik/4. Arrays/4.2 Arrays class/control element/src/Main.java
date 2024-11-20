import java.util.Arrays;
import java.util.Random;
import java.util.Scanner;

public class Main {
    public static void main(String[] args) {
        Scanner scanner = new Scanner(System.in);
        int n = scanner.nextInt();
        Random rand = new Random(scanner.nextInt());
        int control = scanner.nextInt();

        int[] arr = new int[n];
        for (int i = 0; i < arr.length; i++) {
            arr[i]=rand.nextInt(2,16);
        }
        Arrays.sort(arr);
        System.out.println(Arrays.toString(arr));
        int[] result = new int[0];
        for (int i=0;i<arr.length;i++)  {
            if (arr[i] == control) {
                result = Arrays.copyOfRange(arr,0,i+1);
                System.out.println(Arrays.toString(result));
                return;
            }
        }
        System.out.println("ERROR");
    }
}