/*
 * File:   BoardTC.c
 * Author: Federico Foggiato
 *
 * Created on 22 giugno 2022, 15.53
 */

// CONFIG
#pragma config FOSC = HS        // Oscillator Selection bits (HS oscillator)
#pragma config WDTE = OFF       // Watchdog Timer Enable bit (WDT disabled)
#pragma config PWRTE = ON       // Power-up Timer Enable bit (PWRT enabled)
#pragma config BOREN = ON       // Brown-out Reset Enable bit (BOR enabled)
#pragma config LVP = ON         // Low-Voltage (Single-Supply) In-Circuit Serial Programming Enable bit (RB3/PGM pin has PGM function; low-voltage programming enabled)
#pragma config CPD = OFF        // Data EEPROM Memory Code Protection bit (Data EEPROM code protection off)
#pragma config WRT = OFF        // Flash Program Memory Write Enable bits (Write protection off; all program memory may be written to by EECON control)
#pragma config CP = OFF         // Flash Program Memory Code Protection bit (Code protection off)


#define _XTAL_FREQ 8000000


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

#include <xc.h>


//acknowledge comunicazione seriale
char received;

//buffer per conversione da intero a stringa
char convInt[5];

void initSerial(unsigned long int /*baud rate*/);
void sendSerial(char /*carattere da inviare*/);
void init();
void UART_TxString(const char*);
void LCDInit();
void LCDData(char /*data*/, char/*sendData / sendCommand*/);
void LCDPosition(char /*riga*/, char /*colonna*/);
void sendStringLCD(char * /*stringa da inviare*/, char /*riga*/, char /*colonna*/);
void IntToString(int /*intero da convertire*/);


void main(void) {
    init();
    initSerial(9600);
    
    PORTEbits.RE0 = 1;
    int i = 0;
    char c = 0x44;
    char a = 0x00;
    char b = 0xfe;
    char sendDataLCD = 'giovanni';
    while (1)
    {
        sendSerial(a);
        __delay_ms(1);
        sendSerial(b);
        __delay_ms(1);
        sendSerial(c);

        PORTD = 0x01 << i;
        __delay_ms(1000);
        //per decidere dove scrivere, impostare come data la riga + la colonna
        LCDPosition(L_L1, 3);
        LCDData(sendDataLCD, sendData);
        i++;
        //c++;
        //if (c > 126) c = 32;
        //if (c == 'c') c = 'i';
        //else c = 'c';
        if(i > 7) {
            i = 0;            
        }       
    }
    return;
}

//init generale
void init(void)
{
    TRISA |= 0x01;
    TRISB = 0x00;
    TRISC = 0x00;
    TRISD = 0x00;
    TRISE = 0x00;
    PORTA = 0x20;
}

//init specifico porta seriale
void initSerial(unsigned long int baudRate)
{
	TRISC &= ~0x40;
	TRISC |= 0x80;
	
	TXSTA = 0x20;
	
	RCSTA = 0x90;
	
	INTCON |= 0x80;
	INTCON |= 0x40;
	PIE1 |= 0x20; //abilito rcie
	
	received = 0;
	
	SPBRG = (char) (_XTAL_FREQ / (long) (64UL*baudRate)) - 1;
}

//init LCD
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
    LCDData(r + c, sendCommand);
}

//funzione invio dati tramite seriale
void sendSerial(char data)
{
	while(!(PIR1 & 0x10));
	PIR1 &= ~0x10;
	TXREG = data;
}

//invio stringa seriale
void UART_TxString(const char* str)
{
    unsigned char i = 0;

    while (str[i] != 0)
    {
        sendSerial(str[i]);
        i++;
    }
}

//conversione valore intero (fino a 9999) in stringa
void IntToString(int value)
{
    convInt[0] = '0' + (value / 1000);
    convInt[1] = '0' + ((value % 1000) / 100);
    convInt[2] = '0' + ((value % 100) / 10);
    convInt[3] = '0' + (value % 10);
    convInt[4] = '\0';
}

//invia stringa con set riga e colonna a lcd
void sendStringLCD(char *str, char r, char c)
{
    char i;
    for(i=0; str[i] != '\0'; i++)
    {
        LCDPosition(r + (c + i), sendCommand);
        LCDData(str[i], sendData);
    }
}