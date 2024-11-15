import java.util.Scanner;

import static java.lang.Math.sqrt;

public class Main {
    public static void main(String[] args) {
        Scanner scan = new Scanner(System.in);
        double x = scan.nextDouble();
        double y = scan.nextDouble();
        double res1 = sqrt(x*x + y*y);
        x = scan.nextDouble();
        y = scan.nextDouble();
        double res2 = sqrt(x*x + y*y);
        if (res1>res2) {
            System.out.println("Вторая точка ближе");
        }else if (res1<res2) {
            System.out.println("Первая точка ближе");
        }
        else{
            System.out.println("Точки на равных расстояниях");
        }
    }
}