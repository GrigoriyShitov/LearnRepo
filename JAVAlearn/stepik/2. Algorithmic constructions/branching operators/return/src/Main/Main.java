package Main;

import java.util.Scanner;

public class Main {
    public static void main(String[] args) {
        Scanner scan = new Scanner(System.in);
        int x = scan.nextInt();
        int sum=0;
        if (Math.abs(x) <100 || Math.abs(x)>=1000) {
            System.out.println("ERROR");
            return;
        }
        else {
            while (x!=0){
                if ((x%10)%2!=0){
                sum+=Math.abs(x%10);
                }
                x/=10;
            }
        }
        if (sum ==0){
            System.out.println("NO");
        }
        else {
            System.out.println(sum);
        }
    }
}