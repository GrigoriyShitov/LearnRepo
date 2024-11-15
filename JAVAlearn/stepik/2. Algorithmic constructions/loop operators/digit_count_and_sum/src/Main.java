import java.util.Scanner;

public class Main {
    public static void main(String[] args) {
        Scanner scanner = new Scanner(System.in);
        int x = scanner.nextInt();
        int count=0;
        int sum=0;
        int digit;
        do{
            digit =x%10;
            count++;
            sum+=digit;
            x=x/10;
        }while(x != 0);
        System.out.printf("%d %d",count, Math.abs(sum));
    }
}