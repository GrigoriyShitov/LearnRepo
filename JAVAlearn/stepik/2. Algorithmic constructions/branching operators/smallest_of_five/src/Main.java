import java.util.Scanner;

public class Main {
    public static void main(String[] args) {

        Scanner scanner = new Scanner(System.in);
        int x = scanner.nextInt();
        for (int i = 0; i < 4; i++) {
            int y = scanner.nextInt();
            if (x>y){
                x=y;
            }
        }
        System.out.println(x);
    }
}