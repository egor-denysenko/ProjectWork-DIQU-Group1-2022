/*
 * File:   ADCMain.c
 * Author: Federico Foggiato
 *
 * Created on 15 giugno 2022, 15.58
 */

#pragma config FOSC = HS        // Oscillator Selection bits (HS oscillator)
#pragma config WDTE = OFF       // Watchdog Timer Enable bit (WDT disabled)
#pragma config PWRTE = ON       // Power-up Timer Enable bit (PWRT enabled)
#pragma config BOREN = ON       // Brown-out Reset Enable bit (BOR enabled)
#pragma config LVP = ON         // Low-Voltage (Single-Supply) In-Circuit Serial Programming Enable bit (RB3/PGM pin has PGM function; low-voltage programming enabled)
#pragma config CPD = OFF        // Data EEPROM Memory Code Protection bit (Data EEPROM code protection off)
#pragma config WRT = OFF        // Flash Program Memory Write Enable bits (Write protection off; all program memory may be written to by EECON control)
#pragma config CP = OFF         // Flash Program Memory Code Protection bit (Code protection off)


//Define necessari per i comandi dell'LCD
#define L_ON    0X0F
#define L_OFF   0X08
#define L_CLR   0X01
#define L_L1    0X80
#define L_L2    0XC0
#define L_CR    0X0F
#define L_NCR   0X0C
#define L_CFG   0X38    //function set, lavoro con un display a 2 linee e 8 bit
#define L_CUR   0X0E

//Define agevolazione LCD
#define RS 0x04
#define E 0x02
#define LCDC PORTE
#define LCDV PORTD
#define sendData 1
#define sendCommand 0

#define _XTAL_FREQ 20000000

#include <xc.h>

void delay(int);
void ADC_Init();
int ADC_Read(int);
void IntToString(int);
void LCDInit();
void LCDData(char, char);
void LCDPosition(char, char);
void IntToString(int);
void sendStringLCD(char *, char, char);
int map(int, int, int, int, int);


char convInt[5];
char stringTemp[6] = {'T', 'm', 'p', ':', '\0'};
char stringWC[4] = {'W','C',':','\0'};
char stringFreeWC[5] = {'F','r','e','e','\0'};
char stringBuisyWC[5] = {'B','u','i','s','y','\0'};
char stringClosed[12] = {'C','h','i','u','s','u','r','a',' ','P', '\0'};
char stringOpened[12] = {'A','p','e','r','t','u','r','a',' ','P', '\0'};
char stringClear[17] = {' ',' ',' ',' ',' ',' ',' ',' ',' ',' ',' ',' ',' ',' ',' ',' ', '\0'};
char setTemp = 35;
char currentTemp = 0;

void main(void) 
{
    
    int adcValue=0;
    // Configure PORTB and PORTD as output to display the ADC values on LEDs
    TRISD = 0x00;
    TRISB = 0x38;
    PORTD = 0x00;
    TRISC = 0x00;
    

    ADC_Init();             //Initialize the ADC module
    LCDInit();
    LCDData(0x0C, sendCommand);
    while(1)
    {
        //PORTC |= 0x20; 
        adcValue = ADC_Read(2);
        PORTD = adcValue;
        currentTemp = map(adcValue, 56, 158, 28, 78);
        if (currentTemp < setTemp)
        {
            PORTC |= 0x20;
            PORTC &= 0x20; 
        }
        else
        {
            PORTC |= 0x04;
            PORTC &= 0x04;
        }
        
        IntToString(currentTemp);
        sendStringLCD(stringTemp, 1, 0);
        sendStringLCD(convInt, 1, 4);
        sendStringLCD(stringWC, 1, 8);
        
        //Apertura e chiusura porte
        if(!(PORTB &= 0x10)){
            sendStringLCD(stringClosed, 2, 0);
        }
        else
        if(!(PORTB &= 0x08)){
            sendStringLCD(stringOpened, 2, 0);
        }else{sendStringLCD(stringClear, 2, 0);}
        
        //Occupazione toilette
        if(!(PORTB &= 0x20)){
            sendStringLCD(stringBuisyWC, 1, 11);
        }
        else
        {
            sendStringLCD(stringFreeWC, 1, 11);
        }
    }
    return;
}

int map(int x, int in_min, int in_max, int out_min, int out_max) {
  return (x - in_min) * (out_max - out_min) / (in_max - in_min) + out_min;
}

void ADC_Init()
 {
    TRISA = 0x04;
    ADCON0 = 0x91;
    ADCON1 = 0x89;//seleziono l'ingresso analogico da attivare
    __delay_ms(20);  
}


int ADC_Read(int channel)
 {  
    TRISA = 0x04;
	ADCON0 = (ADCON0 & 0xC7) | (unsigned char) (channel<<3); //shift canale di conversione
	__delay_ms(20);
	//delay per carica condensatore
	
	ADCON0 = ADCON0 | 0x04; //avvio la conversione protando adgo a 1
	
	while(!(ADCON0 & ~0x04));
	
	return ADRESL + (unsigned int) (ADRESH << 8);
 }

//conversione valore intero (fino a 9999) in stringa
void IntToString(int value)
{
    convInt[0] = '0' + ((value % 100) / 10);
    convInt[1] = '0' + (value % 10);
    convInt[2] = 'C';    
    convInt[3] = '\0';
    convInt[4] = '\0';
}

//invia stringa con set riga e colonna a lcd
void sendStringLCD(char *str, char r, char c)
{
    char i;
    for(i=0; str[i] != '\0'; i++)
    {
        LCDPosition(r, (c + i));
        LCDData(str[i], sendData);
    }
}

void LCDInit()
{
   TRISE = 0;
   TRISD = 0;
   LCDC = ~0X06;
   __delay_ms(20);
   LCDC = E;
   LCDData(L_CFG, sendCommand);
   __delay_ms(5);
   LCDData(L_CFG, sendCommand);
   __delay_ms(1);
   LCDData(L_CFG, sendCommand);
   LCDData(L_OFF, sendCommand);
   LCDData(L_ON,  sendCommand);
   LCDData(L_CLR, sendCommand);
   LCDData(L_CUR, sendCommand);
   LCDData(L_L1,  sendCommand);
}

//funzione per l'invio di comandi/dati al display LCD
void LCDData(char data, char mode)
{
        LCDC |= E;
        LCDV = data;
        if(mode) 
        {
            LCDC |= RS;
        }
        else 
        {
            LCDC &= ~RS;
        }
        __delay_ms(3);
        LCDC &= ~E;
        __delay_ms(3);
        LCDC |= E;   
}

//funzione per impostare la posizione di visualizzazione del display
void LCDPosition(char r, char c)
{
    if (r == 1){
        LCDData(L_L1 + c, sendCommand);
    }
    else if (r == 2){
        LCDData(L_L2 + c, sendCommand);    
    }
}



