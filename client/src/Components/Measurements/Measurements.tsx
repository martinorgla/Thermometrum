import React from "react";
import {Grid} from "@material-ui/core";
import Temperature from "./Temperature/Temperature";
import Humidity from "./Humidity/Humidity";
import axios from "axios"

class Measurements extends React.Component {
    state = {
        temperature: 0,
        humidity: 0,
    };

    componentDidMount(): void {
        // TODO: Create endpoint for last temperature
        axios.get(`http://185.31.243.56:8001/api/temperature`)
            .then(res => {
                const temperature = res.data[0].temperature;
                const humidity = res.data[0].humidity;

                this.setState({temperature, humidity});
            })
    }

    render() {
        return <Grid container spacing={2}>
            <Grid item xs={6}>
                <Temperature temperature={this.state.temperature}/>
            </Grid>
            <Grid item xs={6}>
                <Humidity humidity={this.state.humidity}></Humidity>
            </Grid>
        </Grid>
    }
}

export default Measurements;