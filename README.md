# Thermometrum
Thermometrum API running on port 8001

## Get temperature
GET localhost:8001/api/temperature/{roomName}
```bash
{
 	"room": "elutuba",
 	"temperature": 24.9,
 	"humidity": 32
}
```

## Save temperature
POST localhost:8001/api/temperature
```bash
{
 	"room": "elutuba",
 	"temperature": 24.9,
 	"humidity": 32
}
```