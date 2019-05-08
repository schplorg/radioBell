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
void loop() {
  Serial.println("loop");
  mySwitch.send(counter, 24);
  counter++;
  /*
  for(int i = 0; i < 3; i++){
    mySwitch.send(counter, 24);
    counter++;
    Serial.println("send");
    delay(10);
  }
  */
  Sleep();
}  
void Sleep(){
    sleep_enable();
    attachInterrupt(0, Wake, LOW);
    Serial.println("Sleep");
    set_sleep_mode(SLEEP_MODE_PWR_DOWN);
    delay(100);
    sleep_cpu();
  }
void Wake(){
  Serial.println("Wake");
  sleep_disable();
  detachInterrupt(0);
}
