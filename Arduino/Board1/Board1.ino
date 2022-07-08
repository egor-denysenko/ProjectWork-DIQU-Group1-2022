#include <dht.h>

int buttonOpenPort = 12; // the number of the push button pin
int ledOpenPort = 11; // the number of the LED pin

int buttonClosePort = 10; // the number of the push button pin
int ledClosePort = 9; // the number of the LED pin

dht DHT; // create dht object
int dhtPin = 6; // the number of the DHT11 sensor pin

#include <Servo.h>
Servo myservo; // create servo object to control a servo
int pos = 0; // variable to store the servo position
int servoPin = 6; // define the pin of servo signal line
Servo myservo1;
int servoPin1 = 13; // define the pin of servo signal line


#include <LiquidCrystal.h>
// initialize the library with the numbers of the interface pins
LiquidCrystal lcd(7, 8, 2, 3, 4, 5);

char receivedData = 0;
String stringData = "";


void setup() {

  Serial.begin(9600); // Initialize the serial port and set the baud rate to 9600

  pinMode(buttonOpenPort, INPUT); // set push button pin into input mode
  pinMode(ledOpenPort, OUTPUT); // set LED pin into output mode

  pinMode(buttonClosePort, INPUT); // set push button pin into input mode
  pinMode(ledClosePort, OUTPUT); // set LED pin into output mode

  // set up the LCD's number of columns and rows:
  lcd.begin(16, 2);
  // Print a message to the LCD
  //lcd.print("hello, world!");

  myservo.attach(servoPin);
  myservo1.attach(servoPin1);
}


void loop() {
  //int chk = DHT.read11(dhtPin);

  // Convert analog value of A0 port into digital value
  int adcVal = analogRead(A0);
  // Calculate voltage
  float v = adcVal * 5.0 / 1024;
  // Calculate resistance value of thermistor
  float Rt = 10 * v / (5 - v);
  // Calculate temperature (Kelvin)
  float tempK = 1 / (log(Rt / 10) / 3950 + 1 / (273.15 + 25));
  // Calculate temperature (Celsius)
  float tempC = tempK - 273.15;
  // Send the result to computer through serial port
  Serial.print("Current temperature is: ");
  Serial.print(tempK);
  Serial.print(" K, " + 0x01);
  Serial.print(tempC);
  Serial.println(" C" + 0x01);

  lcd.setCursor(0, 0);
  lcd.print("Temp:");
  lcd.print(tempC);
  lcd.print("C");

  if (digitalRead(buttonOpenPort) == HIGH) // if the button is not pressed
    digitalWrite(ledOpenPort, LOW); // switch off LED
  else { // if the button is pressed
    digitalWrite(ledOpenPort, HIGH); // switch on LED
    lcd.setCursor(0, 1);
    lcd.print("Apertura Porte");
    for (pos = pos; pos >= 0; pos--) {
      myservo.write(pos);
      myservo1.write(pos);
      delay(10);
    }
    delay(250);
    lcd.setCursor(0, 1);
    for (int i = 0; i < 16; i++)
      lcd.print(" ");
    lcd.setCursor(0, 1);
    lcd.print("Porte Aperte");
  }

  if (digitalRead(buttonClosePort) == HIGH) // if the button is not pressed
    digitalWrite(ledClosePort, LOW); // switch off LED
  else { // if the button is pressed
    digitalWrite(ledClosePort, HIGH); // switch on LED
    lcd.setCursor(0, 1);
    lcd.print("Chiusura Porte");
    for (pos = pos; pos <= 90; pos++) {
      myservo.write(pos);
      myservo1.write(pos);
      delay(10);
    }
    delay(250);
    lcd.setCursor(0, 1);
    for (int i = 0; i < 16; i++)
      lcd.print(" ");
    lcd.setCursor(0, 1);
    lcd.print("Porte Chiuse");
  }


  if (Serial.available()) { // judge whether data has been received
    receivedData = Serial.read(); // read one character
    if (receivedData != '\n') {
      stringData += receivedData;
    }
    else {
      lcd.clear();
      lcd.setCursor(0, 0);
      lcd.print(stringData); // print the received character
    }
    delay(1000);
    lcd.clear();
  }

  /*
    switch (chk)
    {
      case DHTLIB_OK: // When read data successfully, print temperature and humidity data
        Serial.print("Humidity: ");
        Serial.print(DHT.humidity);
        Serial.print("%, Temperature: ");
        Serial.print(DHT.temperature);
        Serial.println("C");
        break;
      case DHTLIB_ERROR_CHECKSUM: // Checksum error
        Serial.println("Checksum error");
        break;
      case DHTLIB_ERROR_TIMEOUT: // Time out error
        Serial.println("Time out error");
        break;
      default: // Unknown error
        Serial.println("Unknown error");
        break;
    }
  */
  delay(250);

}
