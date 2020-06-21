import React from "react";
import {Line} from 'react-chartjs-2';
import axios from "axios";
import {Temperature} from "../../Models/Temperature";

class Chart extends React.Component {
    state = {
        labels: [],
        temperatures: [],
        humidity: [],
    };

    loadTemperatures() {
        // TODO: Move to global config?
        let url = `http://185.31.243.56:8001/api/temperatures`;

        // TODO: Global?
        if (window.location.hostname === 'localhost') {
            url = `http://localhost:8001/api/temperatures`;
        }

        axios.get(url)
            .then(res => {
                if (res.data.data.length) {
                    const data = res.data.data;
                    const labels = this.getLabels(data);
                    const temperatures = this.getTemperatures(data);
                    const humidity = this.getHumidity(data);

                    this.setState({labels, temperatures, humidity});
                }
            })
    }

    getLabels(temperatures: Temperature[]) {
        return temperatures.map((temp: Temperature) => {
            return temp.date
        });
    }

    getTemperatures(temperatures: Temperature[]) {
        return temperatures.map((temperature: Temperature) => {
            return temperature.temperature
        });
    }

    getHumidity(temperatures: Temperature[]) {
        return temperatures.map((temp: Temperature) => {
            return temp.humidity
        });
    }

    renderLineChart(data: object) {
        return <Line data={data}/>
    }

    buildLineChart() {
        const data = {
            labels: this.state.labels,
            datasets: [
                {
                    label: 'Temperatuur',
                    fill: false,
                    lineTension: 0.1,
                    backgroundColor: 'rgba(25,55,55,0.4)',
                    borderColor: 'rgb(25,55,55)',
                    borderCapStyle: 'butt',
                    borderDash: [],
                    borderDashOffset: 0.0,
                    borderJoinStyle: 'miter',
                    pointBorderColor: 'rgb(20,55,55)',
                    pointBackgroundColor: '#373737',
                    pointBorderWidth: 1,
                    pointHoverRadius: 5,
                    pointHoverBackgroundColor: 'rgb(17,55,55)',
                    pointHoverBorderColor: 'rgb(55,55,55)',
                    pointHoverBorderWidth: 2,
                    pointRadius: 1,
                    pointHitRadius: 10,
                    redraw: true,
                    data: this.state.temperatures
                },
                {
                    label: 'Õhuniiskus',
                    fill: false,
                    lineTension: 0.1,
                    backgroundColor: 'rgba(75,192,192,0.4)',
                    borderColor: 'rgba(75,192,192,1)',
                    borderCapStyle: 'butt',
                    borderDash: [],
                    borderDashOffset: 0.0,
                    borderJoinStyle: 'miter',
                    pointBorderColor: 'rgba(75,192,192,1)',
                    pointBackgroundColor: '#fff',
                    pointBorderWidth: 1,
                    pointHoverRadius: 5,
                    pointHoverBackgroundColor: 'rgba(75,192,192,1)',
                    pointHoverBorderColor: 'rgba(220,220,220,1)',
                    pointHoverBorderWidth: 2,
                    pointRadius: 1,
                    pointHitRadius: 10,
                    data: this.state.humidity
                }
            ]
        };

        return data;
    }

    componentDidMount(): void {
        this.loadTemperatures();
        console.log(this.state);
    }

    componentDidUpdate(prevProps: Readonly<{}>, prevState: Readonly<{}>, snapshot?: any): void {
        console.log(this.state);
    }

    render() {
        return (
            <div>
                <h2>Viimased 24 tundi</h2>
                {this.renderLineChart(this.buildLineChart())}
            </div>
        );
    }
}

export default Chart;