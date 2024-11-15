import java.util.Scanner;

public class Main {
    public static void main(String[] args) {
        Scanner scanner = new Scanner(System.in);
        int i=0;
        int count= 0;
        while (i <10) {
            int grade =scanner.nextInt();
            if (grade < 4) {
                count++;
            }
            i++;
        }
        System.out.println(count);
    }
}