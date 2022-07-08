/*
 * File:   soluzioneMain.c
 * Author: Federico Foggiato
 *
 * Created on 6 luglio 2022, 9.13
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
#define BUFMAX 3
#define RXADDRESS 0
#define RXDECINE 1
#define RXUNITA 2
#define MYADDRESS '7'
#define BROADADDRESS '*'


#include <xc.h>

void initPic(void);
void setTimer0();

void UART_init(long int);
void clearBuffer(char *, unsigned char, unsigned char *);//serve a ripulire il buffer, in modo da evitare eventuali problemi

void decode(void);
void stampa (unsigned char);

const char disp[17] = {0x3F, 0x06, 0x5B, 0x4F, 0x66, 0x6D, 0x7D, 0x07,
                        0x7F, 0x6F, 0x77,0x7C, 0x39, 0x5E, 0x79, 0x71, 0x00};

unsigned char timeCount;
char dataReceived[BUFMAX];
unsigned char indexRC;
unsigned char isReceived;
unsigned char valore;

void main(void) {
    
    initPic();
    
    while(1)
    {
        if (isReceived)//tratto il dato
        {
            decode();
            clearBuffer(dataReceived, BUFMAX, &indexRC); //pulizia del buffer, la & serve per accedere all'indirizzo di memoria
            isReceived = 0;//acknowledge dato ricevuto
        }
        
        //stampa del risultato
        stampa(valore);
    }
    return;
}

//inizializzazione pic
void initPic(void)
{
    TRISA = 0x00;
    TRISB = 0x00;
    TRISC = 0x00;
    TRISD = 0x00;
    TRISE = 0x00;
    timeCount = 0;
    indexRC = 0;
    isReceived = 0;
    valore = 0;
    setTimer0();
    UART_init(9600);
}

void setTimer0()
{
    INTCON = 0xA0;
    OPTION_REG = 0x07;
    TMR0 = 6;
}

void UART_init(long int baudRate)
{
    TRISC &= ~0x40;
    TRISC |= 0x80;
    TXSTA |= 0x24;
    RCSTA |= 0x90; //abilita la ricezione continua dei dati
    SPBRG = (char) (_XTAL_FREQ / (unsigned long) (64UL * (unsigned long) baudRate)) - 1;
    INTCON = 0xC0; // 0x80 + 0x40
    PIE1 = 0x20; //abilito interrupt ricezione
}

void clearBuffer(char *buf, unsigned char dim, unsigned char *index)
{
    for (unsigned char i = 0; i < dim; i++ )
    {
        buf[i] = 0;
    }
    *index = 0;
}

void decode(void)
{
    if((dataReceived[RXADDRESS] == MYADDRESS) || (dataReceived[RXADDRESS] == BROADADDRESS)) //definisco il numero della mia scheda 
    {
        valore = (dataReceived[RXDECINE] - '0') * 10 + (dataReceived[RXUNITA] - '0');
        
        /*
        for (unsigned char i = 1; i < BUFMAX; i++ )
        {
            switch (i)
            {
                case RXDECINE:
                    
                    break;
                case RXUNITA:
                    break;
            }
        }
        */
    }
}

void stampa (unsigned char val)
{
    char ledDaAccendere = val / 10;
    char cifra = val % 10;
    
    PORTB &= 0xF0;
    PORTB |= (1 << ledDaAccendere);
    
    PORTAbits.RA5 = 0;
    PORTD = disp[cifra];
    PORTAbits.RA5 = 1;
}

void __interrupt() ISR()
{
    if(T0IF)
    {
        TMR0 = 6;
        timeCount++;
        if(timeCount > 30)
        {
            PORTBbits.RB7 = !PORTBbits.RB7; // ^ = XOR oppure PORTBbits.RB7 = !PORTBbits.RB7
            timeCount = 0;
        }
        T0IF = 0;
    }
    
    if (RCIF)
    {
        dataReceived[indexRC++] = RCREG;
        if (indexRC == BUFMAX)
        {
            isReceived = 1;
        }
        RCIF = 0;
    }
}