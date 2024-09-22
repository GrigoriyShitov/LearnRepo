import java.util.Scanner;

public class Main {
    public static void main(String[] args) {
        Scanner scan = new Scanner(System.in);
        double x = scan.nextDouble();

        double f=Math.log(x+7*Math.sqrt(Math.pow(x,4) + Math.PI));
        System.out.printf("%.4f",f);
    }
}
