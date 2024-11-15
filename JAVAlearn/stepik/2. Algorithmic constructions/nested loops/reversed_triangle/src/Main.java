import java.util.Scanner;

public class Main {
    public static void main(String[] args) {
        Scanner scanner = new Scanner(System.in);

        int n = scanner.nextInt();
        int kol = (n+1)/2;
        if (n<=0){
            System.out.println("ERROR");
            return;
        }
        for (int i = 0; i < kol; i++) {
            for (int j = 0; j < n; j++) {
                if (j>= i && j < n-i) {
                    System.out.print("*");
                }else if (j<i){
                    System.out.print(" ");
                }
            }
            System.out.println();
        }

    }
}