import java.util.Scanner;

class Main {
    public static void main(String[] args) {
        Scanner scan = new Scanner(System.in);
        for (int i = 0; i < 5; i++) {
            int number = scan.nextInt();
            System.out.print(isPrime(number) + " ");
        }
    }

    public static boolean isPrime(int number) {
        if (number == 1 || number <= 0) {
            return false;
        }
        int cnt = 0;
        for (int i = 2; i <= number; i++) {
            if (number % i == 0) {
                cnt++;
                if (cnt == 2) {
                    return false;
                }
            }
        }
        return true;
    }
}