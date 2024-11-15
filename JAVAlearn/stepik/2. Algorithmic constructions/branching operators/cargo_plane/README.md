A cargo airplane has to fly with cargo from point A to point C via point B. The capacity of the aircraft's fuel tank is 300 liters. At point A, the tank is fully refueled. The fuel consumption per 1 km depending on the weight of the airplane's cargo is as follows:

          - up to 500 kg (inclusive) :    1 liter / km

          - up to 1000 kg (inclusive):   4 liters / km

          - Up to 1500 kg (inclusive):   7 liters / km

          - up to 2000 kg (inclusive):   9 liters / km.

          - more than 2000 kg - the aircraft does not lift.

The user enters the distance between points A and B, the distance between points B and C, and the weight of the cargo. The program must calculate what is the minimum amount of fuel needed to refuel the aircraft at point B in order to fly from point A to point C. In case of impossibility to cover any of the distances - the program should output ERROR.

The numerical result of the program is rounded to two decimal places after the decimal point.

| Test Number | Input Data   | Output Data |
|-------------|--------------|-------------|
| 1           | 200 300      | 100.00      |
| 2           | 136 268      | ERROR       |
| 3           | 136 268      | 104.00      |
| 4           | 136 150      | 0.00        |
| 5           | 150 120      | ERROR       |
| 6           | 150 400      | ERROR       |
| 7           | 400 150      | ERROR       |
| 8           | 200 150      | 50.00       |
| 9           | 20 60        | 20.00       |
| 10          | 80 60        | ERROR       |

