export class Temperature {
    temperature: number;
    humidity: number;
    date: string;

    constructor(temperature: number, humidity: number, date: string) {
        this.temperature = temperature;
        this.humidity = humidity;
        this.date = date;
    }
}