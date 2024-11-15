import java.awt.*;
import java.util.Scanner;

public class Main {
    public static void main(String[] args) {
        Scanner scanner = new Scanner(System.in);

        String day = scanner.nextLine();

        day = day.toLowerCase();
        String result;
        switch (day) {
            case "понедельник":
                result = "Monday";
                break;
            case "вторник":
                result = "Tuesday";
                break;
            case "среда":
                result = "Wednesday";
                break;
            case "четверг":
                result = "Thursday";
                break;
            case "пятница":
                result = "Friday";
                break;
            case "суббота":
                result = "Saturday";
                break;
            case "воскресенье":
                result = "Sunday";
                break;
            default:
                result = "ERROR";
                break;
        }
        System.out.println(result);

    }
}