import java.util.Scanner;

public class Main {
    public static void main(String[] args) {
        Scanner scanner = new Scanner(System.in);
        int a = scanner.nextInt();
        int b = scanner.nextInt();
        double result=1;
        int i;
        if (a>b){
            i = b;
            while (i<=a){
                result*=i++;

            }
            System.out.println(result);
        }else{
            i = a;
            while (i<=b){
                result*=i++;
            }
            System.out.println(result);
        }
    }
}