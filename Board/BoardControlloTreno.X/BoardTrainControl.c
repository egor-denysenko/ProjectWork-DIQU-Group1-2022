// CONFIG1
#pragma config FEXTOSC = HS     // External Oscillator mode selection bits (HS (crystal oscillator) above 4MHz; PFM set to high power)
#pragma config RSTOSC = EXT1X   // Power-up default value for COSC bits (EXTOSC operating per FEXTOSC bits)
#pragma config CLKOUTEN = OFF   // Clock Out Enable bit (CLKOUT function is disabled; i/o or oscillator function on OSC2)
#pragma config CSWEN = ON       // Clock Switch Enable bit (Writing to NOSC and NDIV is allowed)
#pragma config FCMEN = ON       // Fail-Safe Clock Monitor Enable bit (FSCM timer enabled)

// CONFIG2
#pragma config MCLRE = ON       // Master Clear Enable bit (MCLR pin function is port defined function)
#pragma config PWRTE = OFF      // Power-up Timer Enable bit (PWRT disabled)
#pragma config LPBOREN = OFF    // Low-Power BOR enable bit (ULPBOR disabled)
#pragma config BOREN = ON       // Brown-out reset enable bits (Brown-out Reset Enabled, SBOREN bit is ignored)
#pragma config BORV = LO        // Brown-out Reset Voltage Selection (Brown-out Reset Voltage (VBOR) set to 1.9V on LF, and 2.45V on F Devices)
#pragma config ZCD = OFF        // Zero-cross detect disable (Zero-cross detect circuit is disabled at POR.)
#pragma config PPS1WAY = ON     // Peripheral Pin Select one-way control (The PPSLOCK bit can be cleared and set only once in software)
#pragma config STVREN = ON      // Stack Overflow/Underflow Reset Enable bit (Stack Overflow or Underflow will cause a reset)

// CONFIG3
#pragma config WDTCPS = WDTCPS_31// WDT Period Select bits (Divider ratio 1:65536; software control of WDTPS)
#pragma config WDTE = OFF        // WDT operating mode (WDT enabled regardless of sleep; SWDTEN ignored)
#pragma config WDTCWS = WDTCWS_7 // WDT Window Select bits (window always open (100%); software control; keyed access not required)
#pragma config WDTCCS = SC       // WDT input clock selector (Software Control)

// CONFIG4
#pragma config WRT = OFF        // UserNVM self-write protection bits (Write protection off)
#pragma config SCANE = available// Scanner Enable bit (Scanner module is available for use)
#pragma config LVP = ON         // Low Voltage Programming Enable bit (Low Voltage programming enabled. MCLR/Vpp pin function is MCLR.)

// CONFIG5
#pragma config CP = OFF         // UserNVM Program memory code protection bit (Program Memory code protection disabled)
#pragma config CPD = OFF        // DataNVM code protection bit (Data EEPROM code protection disabled)

#define _XTAL_FREQ 20000000
#include <xc.h>
#include <stdlib.h>

void SYSTEM_Initialize(void);
void OSCILLATOR_Initialize(void);
void UART_Init(unsigned long int baudrate);
void UART_TxChar(char);
void UART_TxString(const char*);

char dato[50];
char received;
int i;

void main(void)
{
    SYSTEM_Initialize();
    PORTEbits.RE0 = 1;
    int i = 0;
    char c = 0x44;
    char a = 0x00;
    char b = 0xfe;
    while (1)
    {
        UART_TxChar(a);
        __delay_ms(1);
        UART_TxChar(b);
        __delay_ms(1);
        UART_TxChar(c);
        __delay_ms(1);
        UART_TxChar(a);
        __delay_ms(1);
        UART_TxChar(b);
        __delay_ms(1);
        UART_TxChar(c);
        PORTD = 0x01 << i;
        __delay_ms(1000);
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

void UART_Init(unsigned long int baudrate) 
{    
    // Set the EUSART module to the options selected in the user interface.
    TRISC &= ~0x40; // RESET RC6
    TRISC |= 0x080; // SET RC7 
    // ABDOVF no_overflow; SCKP Non-Inverted; BRG16 16bit_generator; WUE disabled; ABDEN disabled; 
    BAUD1CON = 0x08;

    // SPEN enabled; RX9 8-bit; CREN enabled; ADDEN disabled; SREN disabled; 
    RC1STA = 0x90;

    // TX9 8-bit; TX9D 0; SENDB sync_break_complete; TXEN enabled; SYNC asynchronous; BRGH hi_speed; CSRC slave; 
    TX1STA = 0x24;

    // SP1BRGL 25; 
    SP1BRGL = 31;

    // SP1BRGH 0; 
    SP1BRGH = 2;
    
    RXPPS = 0x17;   //RC7->EUSART:RX;    
    RC6PPS = 0x10;   //RC6->EUSART:TX;
}

void SYSTEM_Initialize(void)
{
    //PMD_Initialize();
    //PIN_MANAGER_Initialize();
    TRISD = 0x00;
    TRISE = 0x00;
    UART_Init(9600);
}

void UART_TxChar(char ch)
{    
    while(0 == PIR3bits.TXIF);

    TX1REG = ch;    // Write the data byte to the USART.
}

void UART_TxString(const char* str)
{
    unsigned char i = 0;

    while (str[i] != 0)
    {
        UART_TxChar(str[i]);
        i++;
    }
}