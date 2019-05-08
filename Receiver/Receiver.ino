#include <RCSwitch.h>

RCSwitch mySwitch = RCSwitch();
int counter = -1;
void setup() 
{
  Serial.begin(9600);
  mySwitch.enableReceive(0);
}

void loop() {  
  if (mySwitch.available())
  {
    int value = mySwitch.getReceivedValue();
    if(counter != value){
      counter = value;
    Serial.println(value);  
    }     
    mySwitch.resetAvailable();
  }
  delay(10); 
}
