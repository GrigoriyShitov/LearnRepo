import java.util.Scanner;

class Main {
    public static void main(String[] args) {
        Scanner scan = new Scanner(System.in);
        int a = scan.nextInt();
        int b = scan.nextInt();
        int max = maxNumberDivider(a, b);
        System.out.println(max);
    }

    // put your code here
    public static int maxNumberDivider(int a, int b) {
        if (b < a) {
            int tmp = a;
            a = b;
            b = tmp;
        }
        int max = a, cur = 0;
        for (int i = a; i <= b; i++) {
            if (cur < dividersCount(i)) {
                cur = dividersCount(i);
                max = i;
            }
        }
        return max;
    }

    public static int dividersCount(int num) {
        int count = 0;
        for (int i = 1; i <= num; i++) {
            if (num % i == 0) {
                count++;
            }
        }
        return count;
    }
}