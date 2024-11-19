import java.util.Scanner;

class Main {
    public static void main(String[] args) {
        Scanner sc = new Scanner(System.in);
        int op = sc.nextInt();
        float x = sc.nextFloat();
        switch (op) {
            case 1:
                System.out.printf("%.2f", square(x));
                break;
            case 2:
                float y = sc.nextFloat();
                System.out.printf("%.2f", square(x, y));
                break;
            default:
                break;
        }
    }

    public static float square(float x) {
        return x * x;
    }

    public static float square(float x, float y) {
        return x * y;
    }
}