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

#define cessless 
//#define cessyess 
//#define chadVagon 


void main(void) {
    
    #ifdef cessless
    PORTA = 0x01;
    #endif

    #ifdef cessyess
    PORTB = 0x01;
    #endif
    
    #ifdef chadVagon
    PORTC = 0x01;
    #endif
    return;
}


void init(void)
{
    TRISA = 0x00;
    TRISB = 0x00;
    TRISC = 0x00;
    TRISD = 0x0F;
    TRISE = 0x00;
    PORTA = 0x20;
}