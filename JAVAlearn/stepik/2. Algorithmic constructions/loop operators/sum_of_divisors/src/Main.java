import java.util.Scanner;

public class Main {
    public static void main(String[] args) {
        Scanner scanner = new Scanner(System.in);
        int x = scanner.nextInt();
        int sum=0;
        for (int i=1 ;i<=x;i++){
            if (x%i ==0){
                sum+=i;
            }
        }
        System.out.println(sum);
    }
}