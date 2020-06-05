// Import required libraries
#include "WiFi.h"
#include <Adafruit_Sensor.h>
#include <DHT.h>
#include <HTTPClient.h>

// Replace with your network credentials
const char* ssid = "Not Found 2.4GHz";
const char* password = "wifiparool";

//Your Domain name with URL path or IP address with path
const char* serverName = "http://uvn-243-56.tll07.zonevs.eu:8001/api/temperature";

#define DHTPIN 2

//#define DHTTYPE    DHT11     // DHT 11
#define DHTTYPE    DHT22     // DHT 22 (AM2302)
//#define DHTTYPE    DHT21     // DHT 21 (AM2301)

DHT dht(DHTPIN, DHTTYPE);

// the following variables are unsigned longs because the time, measured in
// milliseconds, will quickly become a bigger number than can be stored in an int.
unsigned long lastTime = 0;
// Set timerDelay to 10 minutes. 1000 millis * 60 seconds * 10 minutes
unsigned long timer = 10;
unsigned long timerDelay = 1000*60*timer;

String readDHTTemperature() {
  // Sensor readings may also be up to 2 seconds 'old' (its a very slow sensor)
  // Read temperature as Celsius (the default)
  float t = dht.readTemperature();

  // Check if any reads failed and exit early (to try again).
  if (isnan(t)) {
    Serial.println("Failed to read from DHT sensor!");
    return String("false");
  }
  else {
    Serial.println(t);
    return String(t);
  }
}

String readDHTHumidity() {
  // Sensor readings may also be up to 2 seconds 'old' (its a very slow sensor)
  float h = dht.readHumidity();

  if (isnan(h)) {
    Serial.println("Failed to read from DHT sensor!");
    return String("false");
  }
  else {
    Serial.println(h);
    return String(h);
  }
}

void setup(){
  // Serial port for debugging purposes
  Serial.begin(115200);
  dht.begin();

  // Connect to Wi-Fi
  WiFi.begin(ssid, password);
  while (WiFi.status() != WL_CONNECTED) {
    delay(1000);
    Serial.println("Connecting to WiFi..");
  }

  // Print ESP32 Local IP Address
  Serial.print("Connected to WiFi network with IP Address: ");
  Serial.println(WiFi.localIP());

  readAndSend();

  Serial.println("Start loop:");
}

void loop() {
    if ((millis() - lastTime) > timerDelay){
        readAndSend();

        lastTime = millis();
    }
}

void readAndSend() {
    //Check WiFi connection status
    if(WiFi.status() == WL_CONNECTED){
        HTTPClient http;

        String temperature = readDHTTemperature();
        String humidity = readDHTHumidity();

        if (temperature == "false" || humidity == "false") {
            Serial.println("Temperature or humidity missing");
        }
        else {
            // Your Domain name with URL path or IP address with path
            http.begin(serverName);

            String request = "{\"room\":\"elutuba\",\"temperature\":"+temperature+",\"humidity\":"+humidity+"}";

            // If you need an HTTP request with a content type: application/json, use the following:
            http.addHeader("Content-Type", "application/json");
            int httpResponseCode = http.POST(request);

            Serial.print("HTTP POST JSON:");
            Serial.println(request);
            Serial.print("HTTP Response code: ");
            Serial.println(httpResponseCode);

            // Free resources
            http.end();
        }
    }
    else {
        Serial.println("WiFi Disconnected");
    }
}
