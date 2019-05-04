#include <RCSwitch.h>

RCSwitch mySwitch = RCSwitch();

void setup() 
{
  Serial.begin(9600);
  mySwitch.enableReceive(0);
}

void loop() {  
  if (mySwitch.available())
  {
    int value = mySwitch.getReceivedValue();
    Serial.println(mySwitch.getReceivedValue());       
    mySwitch.resetAvailable();
  }
  delay(1); 
}
