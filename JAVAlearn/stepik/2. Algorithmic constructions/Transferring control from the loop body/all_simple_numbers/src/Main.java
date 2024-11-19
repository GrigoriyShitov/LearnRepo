import java.util.Scanner;

public class Main {
    public static void main(String[] args) {
        Scanner scanner = new Scanner(System.in);
        int a = scanner.nextInt();
        int b = scanner.nextInt();
        int count=0;
        for(int i = a; i <= b; i++){
            for(int j= 1; j <= i; j++){
                if(i%j==0){
                    count++;
                }
                if (count>2){
                    break;
                }
            }
            if (count<=2 && i!= 1){
                System.out.printf("%d ",i);
            }
            count=0;
        }
    }
}