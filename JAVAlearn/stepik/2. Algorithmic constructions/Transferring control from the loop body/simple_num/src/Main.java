import java.util.Scanner;

public class Main {
    public static void main(String[] args) {

        Scanner scanner = new Scanner(System.in);
        int a = scanner.nextInt();
        if (a==1){
            System.out.println("NO");
            return;
        }
        int count =0;
        for (int i=1; i<=a; i++) {
            if (a % i ==0){
                count++;
            }
            if (count>2){
                System.out.println("NO");
                return;
            }
        }
        System.out.println("YES");
    }
}