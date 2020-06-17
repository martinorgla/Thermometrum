import React from "react";
import Paper from "@material-ui/core/Paper";
import {AcUnit} from "@material-ui/icons";

class Temperature extends React.Component<ownProps> {
    render() {
        return <Paper elevation={3}>
            <h3 style={{whiteSpace: "nowrap"}}>
                <AcUnit/> Temperatuur<br/>
                {this.props.temperature}°C
            </h3>
        </Paper>;
    }
}

export interface ownProps {
    temperature?: number;
}

export default Temperature;