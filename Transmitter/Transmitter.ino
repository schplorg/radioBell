#include <RCSwitch.h>
#include <avr/sleep.h>

RCSwitch mySwitch = RCSwitch();
void setup() {
  Serial.begin(9600);
  Serial.println("setup");
  pinMode(LED_BUILTIN,OUTPUT);
  pinMode(2,INPUT_PULLUP);
  mySwitch.enableTransmit(10);
}
void loop() {
  Serial.println("loop");
  for(int i = 0; i < 10; i++){
    mySwitch.send(1234, 24);
    Serial.println("send");
    delay(10);
  }
  Sleep();
}  
void Sleep(){
    sleep_enable();
    attachInterrupt(0, Wake, LOW);
    Serial.println("Sleep");
    set_sleep_mode(SLEEP_MODE_PWR_DOWN);
    delay(3000);
    sleep_cpu();
  }
void Wake(){
  Serial.println("Wake");
  sleep_disable();
  detachInterrupt(0);
}
