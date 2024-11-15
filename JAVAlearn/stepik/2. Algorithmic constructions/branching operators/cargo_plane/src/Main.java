import java.util.Scanner;

public class Main {
    public static void main(String[] args) {

        Scanner scanner = new Scanner(System.in);
        int AB = scanner.nextInt();
        int BC = scanner.nextInt();
        int mass = scanner.nextInt();
        float airplaneFuel = 300;
        int consumption=0;
        if (mass > 2000){
            System.out.println("ERROR");
            return;
        } else if (mass <=500) {
            consumption=1;
        } else if (mass <=1000) {
            consumption=4;
        } else if (mass <=1500) {
            consumption=7;
        } else if (mass <=2000) {
            consumption=9;
        }
        float fuelAB = consumption*AB;
        if (fuelAB>300){
            System.out.println("ERROR");
            return;
        }
        airplaneFuel-=fuelAB;
        float fuelBC = consumption*BC;
        if (fuelBC>300){
            System.out.println("ERROR");
            return;
        }
        System.out.printf("%.2f",(fuelBC-airplaneFuel)>0 ? fuelBC-airplaneFuel: 0);
    }
}