import java.util.Scanner;

public class Main {
    public static void main(String[] args) {
        Scanner scanner = new Scanner(System.in);
        int control = scanner.nextInt();
        int x;
        int count =0;
        do{
            x = scanner.nextInt();
            if (control %x==0 && x>0){
                count++;
            }
        } while(x>=0);
        System.out.println(count);
    }
}