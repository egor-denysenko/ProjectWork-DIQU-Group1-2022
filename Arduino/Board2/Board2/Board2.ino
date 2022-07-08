#include <SPI.h>
#include <RFID.h>

//D10:pin of card reader SDA. D9:pin of card reader RST
RFID rfid(10, 9);
unsigned char status;
unsigned char str[MAX_LEN]; //MAX_LEN is 16: size of the array

char openedPortLed = 8;

int latchPin = 6; // Pin connected to ST_CP of 74HC595（Pin12）
int clockPin = 7; // Pin connected to SH_CP of 74HC595（Pin11）
int dataPin = 5; // Pin connected to DS of 74HC595（Pin14）

// Define the encoding of characters 0-F of the common-anode 7-segment Display
byte num[] = {0xc0, 0xf9, 0xa4, 0xb0, 0x99, 0x92, 0x82, 0xf8, 0x80, 0x90, 0x88, 0x83,
              0xc6, 0xa1, 0x86, 0x8e
             };



void setup()
{
  Serial.begin(9600);
  SPI.begin();
  rfid.init(); //initialization
  Serial.println("Please put the card to the induction area...");

  pinMode(latchPin, OUTPUT);
  pinMode(clockPin, OUTPUT);
  pinMode(dataPin, OUTPUT);

}

void loop()
{
  //Search card, return card types
  if (rfid.findCard(PICC_REQIDL, str) == MI_OK) {
    Serial.println("Carta trovata");
    Serial.println("La porta rimane aperta per 5 secondi ");
    digitalWrite(openedPortLed, HIGH);
    // Cycling display 0-F
    for (int i = 0; i <= 5; i++) {
      // Output low level to latchPin
      digitalWrite(latchPin, LOW);
      // Send serial data to 74HC595
      shiftOut(dataPin, clockPin, MSBFIRST, num[i]);
      // Output high level to latchPin, and 74HC595 will update the data to the parallel output port.
      digitalWrite(latchPin, HIGH);
      delay(1000);
    }
    rfid.selectTag(str);
  }
  rfid.halt(); // command the card to enter sleeping state
  digitalWrite(latchPin, LOW);
  shiftOut(dataPin, clockPin, MSBFIRST, 0xc0);
  digitalWrite(latchPin, HIGH);
  digitalWrite(openedPortLed, LOW);
}
