/*
   Created by △ BegoonLab
   
   @author Alexander Begoon

   Firmware for Railway Module version 1.0

   2023
*/

#include "NmraDcc.h"
#include "PCA9685.h"
#include "ezLED.h"

//////// CONFIGURATION ///////////

//#define Two_PWM_Controllers
#define This_Decoder_Address 10
#define PWM_Frequency 50
// Define the Arduino input Pin number for the DCC Signal 
#define DCC_PIN     2

#define NOTIFY_DCC_FUNC

// Uncomment to print all DCC F functions
//#define DEBUG_DCC_FUNC
// Uncomment to print all DCC Packets
//#define NOTIFY_DCC_MSG

//////// END OF CONFIGURATION ///////////

PCA9685 pwmController1(B000000);        // Library using B000000 (A5-A0) i2c address, and default Wire @400kHz
#ifdef Two_PWM_Controllers
PCA9685 pwmController2(B000001);        // Library using B000001 (A5-A0) i2c address, and default Wire @400kHz
#endif
// Not a real device, will act as a proxy to pwmController1 and pwmController2, using all-call i2c address 0xE0, and default Wire @400kHz
PCA9685 pwmControllerAll(PCA9685_I2C_DEF_ALLCALL_PROXYADR);

NmraDcc Dcc;
DCC_MSG Packet;
ezLED led(13);  // Status LED
 
struct CVPair
{
  uint16_t  CV;
  uint8_t   Value;
};

CVPair FactoryDefaultCVs [] =
{
  // The CV Below defines the Short DCC Address
  {CV_MULTIFUNCTION_PRIMARY_ADDRESS, This_Decoder_Address&0x7F},

  // These two CVs define the Long DCC Address
  {CV_MULTIFUNCTION_EXTENDED_ADDRESS_MSB, ((This_Decoder_Address>>8)&0x7F)+192},
  {CV_MULTIFUNCTION_EXTENDED_ADDRESS_LSB, This_Decoder_Address&0xFF},

// ONLY uncomment 1 CV_29_CONFIG line below as appropriate
//  {CV_29_CONFIG,                                      0}, // Short Address 14 Speed Steps
  {CV_29_CONFIG,                       CV29_F0_LOCATION}, // Short Address 28/128 Speed Steps
//  {CV_29_CONFIG, CV29_EXT_ADDRESSING | CV29_F0_LOCATION}, // Long  Address 28/128 Speed Steps  
};
uint8_t FactoryDefaultCVIndex = sizeof(FactoryDefaultCVs)/sizeof(CVPair);
void notifyCVResetFactoryDefault()
{
  // Make FactoryDefaultCVIndex non-zero and equal to num CV's to be reset 
  // to flag to the loop() function that a reset to Factory Defaults needs to be done
  FactoryDefaultCVIndex = sizeof(FactoryDefaultCVs)/sizeof(CVPair);
};

////////////////// SETUP ////////////////////
void setup() {
  pinMode(3, OUTPUT);
  digitalWrite(3, HIGH); // Disable PWM Drivers

  Serial.begin(115200); 
  Wire.begin();

  Serial.println("Setup...");

  digitalWrite(3, LOW); // Enable PWM Drivers

  pwmControllerAll.resetDevices();    // Resets all PCA9685 devices on i2c line
  
  pwmController1.init();              // Initializes first module using default totem-pole driver mode, and default disabled phase balancer
  pwmController1.setPWMFrequency(PWM_Frequency);
  #ifdef Two_PWM_Controllers
  pwmController2.init();              // Initializes second module using default totem-pole driver mode, and default disabled phase balancer
  pwmController2.setPWMFrequency(PWM_Frequency);
  #endif
  
  pwmControllerAll.initAsProxyAddresser(); // Initializes 'fake' module as all-call proxy addresser
  
  // Enables all-call support to module from 'fake' all-call proxy addresser
  pwmController1.enableAllCallAddress(pwmControllerAll.getI2CAddress());
  
  #ifdef Two_PWM_Controllers
  pwmController2.enableAllCallAddress(pwmControllerAll.getI2CAddress()); // On both
  #endif
  
  pwmController1.setChannelOff(0);    // Turn channel 0 off
  
  #ifdef Two_PWM_Controllers
  pwmController2.setChannelOff(0);    // On both
  #endif
  
  // pwmControllerAll.setChannelPWM(0, 4096); // Enables full on on both pwmController1 and pwmController2
  
  delay(500);
  Serial.println("End setup.");
  // Setup which External Interrupt, the Pin it's associated with that we're using and enable the Pull-Up
  // Many Arduino Cores now support the digitalPinToInterrupt() function that makes it easier to figure out the
  // Interrupt Number for the Arduino Pin number, which reduces confusion. 
  #ifdef digitalPinToInterrupt
    Dcc.pin(DCC_PIN, 0);
  #else
    Dcc.pin(0, DCC_PIN, 1);
  #endif
  Dcc.init( MAN_ID_DIY, 100, FLAGS_MY_ADDRESS_ONLY, 0 ); // Enable DCC Receiver
  Serial.println("DCC initiated");
  led.blinkNumberOfTimes(250, 750, 10);
}


////////////////// LOOP ////////////////////
void loop() {
  // You MUST call the NmraDcc.process() method frequently from the Arduino loop() function for correct library operation
  Dcc.process();
  
  if( FactoryDefaultCVIndex && Dcc.isSetCVReady())
  {
    FactoryDefaultCVIndex--; // Decrement first as initially it is the size of the array 
    Dcc.setCV( FactoryDefaultCVs[FactoryDefaultCVIndex].CV, FactoryDefaultCVs[FactoryDefaultCVIndex].Value);
  }

  // MUST call the led.loop() function in loop()
  led.loop();
}

#ifdef  NOTIFY_DCC_FUNC
void notifyDccFunc(uint16_t Addr, DCC_ADDR_TYPE AddrType, FN_GROUP FuncGrp, uint8_t FuncState)
{
  led.blinkNumberOfTimes(15, 15, 1);
#ifdef DEBUG_DCC_FUNC
  Serial.print("notifyDccFunc: Addr: ");
  Serial.print(Addr,DEC);
  Serial.print( (AddrType == DCC_ADDR_SHORT) ? 'S' : 'L' );
  Serial.print("  Function Group: ");
  Serial.print(FuncGrp,DEC);
#endif
  switch( FuncGrp )
   {
#ifdef NMRA_DCC_ENABLE_14_SPEED_STEP_MODE    
     case FN_0:
#ifdef DEBUG_DCC_FUNC
       Serial.print(" FN0: ");
       Serial.println((FuncState & FN_BIT_00) ? "1  " : "0  "); 
#endif
       setPWMController1(0 , FuncState & FN_BIT_00)
       break;
#endif
     case FN_0_4:
       if(Dcc.getCV(CV_29_CONFIG) & CV29_F0_LOCATION) // Only process Function 0 in this packet if we're not in Speed Step 14 Mode
       {
#ifdef DEBUG_DCC_FUNC
         Serial.print(" FN 0: ");
         Serial.print((FuncState & FN_BIT_00) ? "1  ": "0  ");
#endif 
         setPWMController1(0 , FuncState & FN_BIT_00);
       }
#ifdef DEBUG_DCC_FUNC       
       Serial.print(" FN 1-4: ");
       Serial.print((FuncState & FN_BIT_01) ? "1  ": "0  ");
       Serial.print((FuncState & FN_BIT_02) ? "1  ": "0  ");
       Serial.print((FuncState & FN_BIT_03) ? "1  ": "0  ");
       Serial.println((FuncState & FN_BIT_04) ? "1  ": "0  ");
#endif
       setPWMController1(1 , FuncState & FN_BIT_01);
       setPWMController1(2 , FuncState & FN_BIT_02);
       setPWMController1(3 , FuncState & FN_BIT_03);
       setPWMController1(4 , FuncState & FN_BIT_04);
       break;
    
     case FN_5_8:
#ifdef DEBUG_DCC_FUNC 
       Serial.print(" FN 5-8: ");
       Serial.print((FuncState & FN_BIT_05) ? "1  ": "0  ");
       Serial.print((FuncState & FN_BIT_06) ? "1  ": "0  ");
       Serial.print((FuncState & FN_BIT_07) ? "1  ": "0  ");
       Serial.println((FuncState & FN_BIT_08) ? "1  ": "0  ");
#endif
       setPWMController1(5 , FuncState & FN_BIT_05);
       setPWMController1(6 , FuncState & FN_BIT_06);
       setPWMController1(7 , FuncState & FN_BIT_07);
       setPWMController1(8 , FuncState & FN_BIT_08);
       break;
    
     case FN_9_12:
#ifdef DEBUG_DCC_FUNC
       Serial.print(" FN 9-12: ");
       Serial.print((FuncState & FN_BIT_09) ? "1  ": "0  ");
       Serial.print((FuncState & FN_BIT_10) ? "1  ": "0  ");
       Serial.print((FuncState & FN_BIT_11) ? "1  ": "0  ");
       Serial.println((FuncState & FN_BIT_12) ? "1  ": "0  ");
#endif
       setPWMController1(9 , FuncState & FN_BIT_09);
       setPWMController1(10 , FuncState & FN_BIT_10);
       setPWMController1(11 , FuncState & FN_BIT_11);
       setPWMController1(12 , FuncState & FN_BIT_12);
       break;

     case FN_13_20:
#ifdef DEBUG_DCC_FUNC
       Serial.print(" FN 13-20: ");
       Serial.print((FuncState & FN_BIT_13) ? "1  ": "0  ");
       Serial.print((FuncState & FN_BIT_14) ? "1  ": "0  ");
       Serial.print((FuncState & FN_BIT_15) ? "1  ": "0  ");
       Serial.print((FuncState & FN_BIT_16) ? "1  ": "0  ");
       Serial.print((FuncState & FN_BIT_17) ? "1  ": "0  ");
       Serial.print((FuncState & FN_BIT_18) ? "1  ": "0  ");
       Serial.print((FuncState & FN_BIT_19) ? "1  ": "0  ");
       Serial.println((FuncState & FN_BIT_20) ? "1  ": "0  ");
#endif
       setPWMController1(13 , FuncState & FN_BIT_13);
       setPWMController1(14 , FuncState & FN_BIT_14);
       setPWMController1(15 , FuncState & FN_BIT_15);
#ifdef Two_PWM_Controllers
       setPWMController2( 0 , FuncState & FN_BIT_16);
       setPWMController2( 1 , FuncState & FN_BIT_17);
       setPWMController2( 2 , FuncState & FN_BIT_18);
       setPWMController2( 3 , FuncState & FN_BIT_19);
       setPWMController2( 4 , FuncState & FN_BIT_20);
#endif
       break;
  
     case FN_21_28:
#ifdef DEBUG_DCC_FUNC
       Serial.print(" FN 21-28: ");
       Serial.print((FuncState & FN_BIT_21) ? "1  ": "0  ");
       Serial.print((FuncState & FN_BIT_22) ? "1  ": "0  ");
       Serial.print((FuncState & FN_BIT_23) ? "1  ": "0  ");
       Serial.print((FuncState & FN_BIT_24) ? "1  ": "0  ");
       Serial.print((FuncState & FN_BIT_25) ? "1  ": "0  ");
       Serial.print((FuncState & FN_BIT_26) ? "1  ": "0  ");
       Serial.print((FuncState & FN_BIT_27) ? "1  ": "0  ");
       Serial.println((FuncState & FN_BIT_28) ? "1  ": "0  ");
#endif
#ifdef Two_PWM_Controllers
       setPWMController2( 5 , FuncState & FN_BIT_21);
       setPWMController2( 6 , FuncState & FN_BIT_22);
       setPWMController2( 7 , FuncState & FN_BIT_23);
       setPWMController2( 8 , FuncState & FN_BIT_24);
       setPWMController2( 9 , FuncState & FN_BIT_25);
       setPWMController2( 10 , FuncState & FN_BIT_26);
       setPWMController2( 11 , FuncState & FN_BIT_27);
       setPWMController2( 12 , FuncState & FN_BIT_28);
#endif
       break;  
   }
}
#endif

void setPWMController1 (int channel, bool value) {
  if (value == true) {
    pwmController1.setChannelPWM(channel, 200);
  } else {
    pwmController1.setChannelPWM(channel, 350);
  }
}

#ifdef Two_PWM_Controllers
void setPWMController2 (int channel, bool value) {
  if (value == true) {
    pwmController2.setChannelPWM(channel, 200);
  } else {
    pwmController2.setChannelPWM(channel, 350);
  }
}
#endif

#ifdef  NOTIFY_DCC_MSG
void notifyDccMsg( DCC_MSG * Msg)
{
  Serial.print("notifyDccMsg: ") ;
  for(uint8_t i = 0; i < Msg->Size; i++)
  {
    Serial.print(Msg->Data[i], HEX);
    Serial.write(' ');
  }
  Serial.println();
}
#endif
