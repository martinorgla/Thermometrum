import React from "react";
import {Grid} from "@material-ui/core";
import Temperature from "./Temperature/Temperature";
import Humidity from "./Humidity/Humidity";
import axios from "axios"
import Chart from "../Chart/Chart";

class Measurements extends React.Component {
    state = {
        temperature: 0,
        humidity: 0,
        lastReading: '',
    };

    componentDidMount(): void {
        // TODO: Move to global config?
        let url = `http://185.31.243.56:8001/api/temperature`;

        // TODO: Global?
        if (window.location.hostname === 'localhost') {
            url = `http://localhost:8001/api/temperature`;
        }

        axios.get(url)
            .then(res => {
                const temperature = res.data.data.temperature;
                const humidity = res.data.data.humidity;
                const lastReading = res.data.data.date;
                this.setState({temperature, humidity, lastReading});
            })
    }

    render() {
        return <Grid container spacing={2}>
            <Grid item xs={12}>
                <h5>{this.state.lastReading}</h5>
            </Grid>
            <Grid item xs={6}>
                <Temperature temperature={this.state.temperature}/>
            </Grid>
            <Grid item xs={6}>
                <Humidity humidity={this.state.humidity}></Humidity>
            </Grid>
            <Grid item xs={12}>
                <Chart></Chart>
            </Grid>
        </Grid>
    }
}

export default Measurements;