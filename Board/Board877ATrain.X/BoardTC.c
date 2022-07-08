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
char recievedData;

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
char reciveSerial();


void main(void) {
    init();
    initSerial(9600);
    
    PORTE = 0x00;
    received = 0;
    int i = 0;
    
    char mittente = 0x01;
    char destinatario = 0xfe;
    char codice = 0x45;
    char temperatura = 0x35;
    char sendDataLCD[32] = {'C', 'a', 'm', 'b', 'i', 'o', ' ', 't', 'e', 'm', 'p', 'e', 'r', 'a', 't', 'u', 'r', 'a', '\0'};
    
    while (1)
    {
        
        //recievedData = reciveSerial();// ricezione dati da seriale        

        
        PORTB = 0x01;
        __delay_ms(500);
        
        if(received)
        {
            received = 0;
            PORTD = recievedData;
        } 
        
        PORTB = 0x00;
        __delay_ms(500);
        
        /*
        sendSerial(mittente);//invio dati tramite seriale
        __delay_ms(1);
        sendSerial(destinatario);
        __delay_ms(1);
        sendSerial(codice);
        __delay_ms(1);
        sendSerial(temperatura);
        __delay_ms(1);
        */
        
        //PORTD = 0x01 << i; //led portd che si illuminano in sequenza
        //__delay_ms(1000);
    }
    return;
}

void __interrupt() ISR()
{
    //PORTD = 2;
    if(PIR1 & 0x20)
    {
        recievedData = RCREG;
        PORTD = 1;
        received = 1;
        //RCIF = 0;
    }
    if(TXIF)
    {
        TXIF = 0;
    }
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
    PORTD = 0;
    INTCON = 0xC0;
    PIE1 = 0x60;
}

//init specifico porta seriale
void initSerial(unsigned long int baudRate)
{
	TRISC &= ~0x40;
	TRISC |= 0x80;
	TXSTA = 0x20; //provare a impostare a 24
	
	RCSTA = 0x90;
	
	INTCON |= 0x80;
	INTCON |= 0x40;
	PIE1 |= 0x20; //abilito rcie
	
	received = 0;
	
	SPBRG =(_XTAL_FREQ / (long) (64UL*baudRate)) - 1;
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
    if (r == 1){
        LCDData(L_L1 + c, sendCommand);
    }
    else if (r == 2){
        LCDData(L_L2 + c, sendCommand);    
    }
}

//funzione invio dati tramite seriale
void sendSerial(char data)
{
    PORTE |= 0x01;
	while(!(PIR1 & 0x10));
	PIR1 &= ~0x10;
	TXREG = data;
    while(!(PIR1 & 0x10));
    PORTE &= ~0x01;
}

//funzione ricezione dati tramite seriale

char reciveSerial(){
    while(RCIF==0);    // Wait till the data is received 
    RCIF=0;            // Clear receiver flag
    return(RCREG);  
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
    convInt[0] = '0' + (value % 10000) / 1000;
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
        LCDPosition(r, (c + i));
        LCDData(str[i], sendData);
    }
}


//funzionamento servo motore
void servoRotate0() //0 Degree
{
  unsigned int i;
  for(i=0;i<50;i++)
  {
    PORTB &= 0x01;
    __delay_us(800);
    PORTB &= 0x00;
    __delay_us(19200);
  }
}

void servoRotate90() //90 Degree
{
  unsigned int i;
  for(i=0;i<50;i++)
  {
    PORTB &= 0x01;
    __delay_us(1500);
    PORTB &= 0x00;
    __delay_us(18500);
  }
}

void servoRotate180() //180 Degree
{
  unsigned int i;
  for(i=0;i<50;i++)
  {
    PORTB &= 0x01;
    __delay_us(2200);
    PORTB &= 0x00;
    __delay_us(17800);
  }
}
//fine funzionamento servo motore