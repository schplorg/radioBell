#include <RCSwitch.h>

RCSwitch mySwitch = RCSwitch();
void setup() {
  mySwitch.enableTransmit(10);
}
void loop() {
  mySwitch.send(1234, 24); // Der 433mhz Sender versendet die Dezimalzahl „1234“
  delay(1000);
}  
