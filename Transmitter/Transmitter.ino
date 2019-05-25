#include <RCSwitch.h>
#include <avr/sleep.h>

RCSwitch mySwitch = RCSwitch();
int counter = 1;
void setup() {
  Serial.begin(9600);
  Serial.println("setup");
  pinMode(LED_BUILTIN,OUTPUT);
  pinMode(2,INPUT_PULLUP);
  mySwitch.enableTransmit(10);
}
bool wake = false;
void loop() {
  Serial.println("loop");
  wake = true;
  attachInterrupt(0, Press, LOW);
  while(wake){
    delay(5000);
    wake = false;
  }
  detachInterrupt(0);
  Sleep();
}  
bool pressed = false;
void Press(){  
  wake = true;
  if(pressed){
    counter++;
    return;
  }else{
    pressed = true;
  }
  Serial.println("Press");  
  mySwitch.send(counter, 24);
  /*
  for(int i = 0; i < 8; i++){
    mySwitch.send(counter, 24);
    delay(5);
  }
  */
  counter++;  
  pressed = false;
}
void Sleep(){
    Serial.println("Sleep");
    sleep_enable();
    attachInterrupt(0, Wake, LOW);
    set_sleep_mode(SLEEP_MODE_PWR_DOWN);
    delay(100);
    sleep_cpu();
  }
void Wake(){
  Serial.println("Wake");
  sleep_disable();
  detachInterrupt(0);
}
