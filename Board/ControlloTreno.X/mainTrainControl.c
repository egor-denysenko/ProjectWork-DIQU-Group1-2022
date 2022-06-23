/*
 * File:   mainTrainControl.c
 * Author: Federico Foggiato
 *
 * Created on 20 maggio 2022, 9.30
 */

#pragma config FOSC = HS        // Oscillator Selection bits (HS oscillator)
#pragma config WDTE = OFF       // Watchdog Timer Enable bit (WDT disabled)
#pragma config PWRTE = ON       // Power-up Timer Enable bit (PWRT enabled)
#pragma config BOREN = ON       // Brown-out Reset Enable bit (BOR enabled)
#pragma config LVP = ON         // Low-Voltage (Single-Supply) In-Circuit Serial Programming Enable bit (RB3/PGM pin has PGM function; low-voltage programming enabled)
#pragma config CPD = OFF        // Data EEPROM Memory Code Protection bit (Data EEPROM code protection off)
#pragma config WRT = OFF        // Flash Program Memory Write Enable bits (Write protection off; all program memory may be written to by EECON control)
#pragma config CP = OFF         // Flash Program Memory Code Protection bit (Code protection off)

#define _XTAL_FREQ 8000000

#include <xc.h>

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

//Define vari
#define TMAX 30
#define TMIN 5

//Define maschera seriale
#define SER 0x01

void LCDInit();
void LCDData(char, char);
void initSerial(unsigned long int);
void sendSerial(char);
void sendString(char);
void sendLCDString(char *);
char Send595(char);
void init_ADC(void);
int read_ADC(char channel);

#define cessless 
//#define cessyess 
//#define chadVagon 

char pos = 0;
char data = 'gigi';
char dataLCD [12] =  "T: 22.3C";
char received;
char c = 0;

char codiceSeriale0 = 0xFF;
char codiceSeriale1 = 0xF4;
char codiceSeriale2 = 0x76;


void main(void) {
    LCDInit();
    initSerial(9600);
    init_ADC();
    while(1)
    {
        #ifdef cessless
        
        PORTA = 0x01;
        //Write LCD
        if (c==0){
            sendLCDString(dataLCD);
            c = 1;
        }        

        /*
        //SerialComm da vedere
        Send595(codiceSeriale0);
        */
        //ADC funzionamento
        TRISD=0x00;
        while(1){
            PORTD = read_ADC(0) >> 2;             //era usato per fare un test se funzionava read_ADC()
        }
        

        #endif

        #ifdef cessyess
        PORTB = 0x01;
        data = 'B';
        LCDData(data, sendData);
        __delay_ms(500);
        #endif

        #ifdef chadVagon
        PORTC = 0x01;
        data = 'C';
        LCDData(data, sendData);
        __delay_ms(500);
        #endif
    }
    return;
}



void init(void)
{
    TRISA |= 0x01;
    TRISB = 0x00;
    TRISC = 0x00;
    TRISD = 0x00;
    TRISE = 0x00;
    PORTA = 0x20;
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
	
	SPBRG = (char) (_XTAL_FREQ/(long) (64UL*baudRate)) - 1;
}

void sendSerial(char data)
{
	while(!(PIR1 & 0x10));
	PIR1 &= ~0x10;
	TXREG = data;
}

/*
 * char Send595(char code)
{
    __delay_ms(1000);
	for(char i=0; i<8; i++)
	{
		PORTB = (code >> i) & SER;
        if((code >> i) & SER)
        {
            PORTB |= SER;
        }
        else
        {
            PORTB &= ~SER;
        }
		PORTB |= 0x02;
		__delay_ms(1);
        PORTB &= ~(0x02);
        __delay_ms(1);
	}
	PORTB |= 0x04;
}
*/

/*
void sendString(char str)
{
	int i = 0;
	while(str[i] != '\0')
	{
		sendSerial(str[i++]);
	}
}
*/
/*
 * void __interrupt() ISR()
{
    if(RCIF)
	{
		data = RCREG;
		received = 1;
		RCIF = 0;
	}
}
 */

void sendLCDString(char *string)
{
    char i;
    for(i=0; string[i]!='\0'; i++){
        LCDData(string[i], sendData);
    }
}

void init_ADC(void)
{
    ADCON0 = 0x01;
    ADCON1 = 0x8E;
    __delay_us(20);
}

int read_ADC(char channel)
{
    ADCON0 &= 0xC7;                      //devo prima resettare i tre chs altrimenti nel caso di chs = 111 quando faccio OR mi resterebbero a 1
    ADCON0 |= (channel & 0x07) << 3;     //dopo aver resettato solo i 3 chs posso inserire il "channel" shiftato di 3
    __delay_us(10);
    ADCON0 |= 0x04;                      //inizio la conversione settando il GO a 1
    while(ADCON0 & 0x04){
    }
    return ADRESL | (ADRESH <<8 );
}