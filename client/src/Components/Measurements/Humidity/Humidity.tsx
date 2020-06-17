import React from "react";
import Paper from "@material-ui/core/Paper";
import OpacityIcon from '@material-ui/icons/Opacity';

class Humidity extends React.Component<ownProps> {
    render() {
        return <Paper elevation={3}>
            <h3 style={{whiteSpace: "nowrap"}}>
                <div>
                    <OpacityIcon/> Õhuniiskus<br/>
                    {this.props.humidity}%
                </div>
            </h3>
        </Paper>;
    }
}

export interface ownProps {
    humidity?: number;
}

export default Humidity;