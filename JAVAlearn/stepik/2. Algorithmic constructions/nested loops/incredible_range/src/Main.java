import java.util.Scanner;

public class Main {
    public static void main(String[] args) {
        Scanner scanner = new Scanner(System.in);
        int a = scanner.nextInt();
        int b = scanner.nextInt();
        int cur,out=0,temp,prevsum=Integer.MIN_VALUE,sum=0;
        if (a>b){
            temp =a;
            a=b;
            b=temp;
        }
        for(int i=a; i<b;i++){
            for(cur=Math.abs(i); cur >0;cur/=10){
                sum+=cur%10;
            }
            if (prevsum<sum){
                prevsum=Math.abs(sum);
                out=i;
            }
            sum=0;
        }
        System.out.println(out);
    }
}