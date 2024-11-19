import java.util.Scanner;

class Main {
    public static void main(String[] args) {
        Scanner scan = new Scanner(System.in);
        int a = scan.nextInt();
        int b = scan.nextInt();
        int kol = simpleInRange(a, b);
        System.out.println(kol);
    }

    public static int simpleInRange(int a, int b) {
        if (b < a) {
            int temp = a;
            a = b;
            b = temp;
        }
        int primeCnt = 0;
        int divCnt = 0;
        for (int i = a; i <= b; i++) {
            if (i == 1) {
                continue;
            }
            for (int j = 1; j <= i; j++) {
                if (i % j == 0) {
                    divCnt++;
                }
            }
            if (divCnt == 2) {
//                System.out.println(i);
                primeCnt++;
            }
            divCnt = 0;
        }
        return primeCnt;
    }
}