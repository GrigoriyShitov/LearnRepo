import java.util.Scanner;

class Main {
    public static void main(String[] args) {
        Scanner sc = new Scanner(System.in);
        int size = sc.nextInt();
        char sign = sc.next().charAt(0);
        printTriangle(size, sign);
    }

    public static void printTriangle(int size, char sign) {

        int height = size / 2 + size % 2;
        int probel, kol;
        if (size % 2 == 0) {
            probel = size / 2 - 1;
            kol = 2;
        } else {
            probel = size / 2;
            kol = 1;
        }

        for (int i = 0; i < height; i++) {
            for (int j = 1; j <= probel; j++) { //выводим пробелы
                System.out.print(' ');
            }
            for (int j = 1; j <= kol; j++) { //выводим символы
                System.out.print(sign);
            }
            System.out.println();
            probel--;//готовимся к следующей строке
            kol += 2;
        }
    }
}