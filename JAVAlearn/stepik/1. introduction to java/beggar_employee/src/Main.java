import java.util.Scanner;

public class Main {
    public static void main(String[] args) {
        Scanner scan = new Scanner(System.in);
        int Salary_1 = scan.nextInt();
        int Salary_2 = scan.nextInt();
        int Salary_3 = scan.nextInt();

        int temp = Salary_1>Salary_2? Salary_2: Salary_1;
        int min = temp>Salary_3? Salary_3: temp;
        temp = Salary_1>Salary_2? Salary_1:Salary_2;
        int max = temp > Salary_3? temp:Salary_3;
        System.out.println(max-min);
    }
}
