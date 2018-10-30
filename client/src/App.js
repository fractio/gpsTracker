import React, { Component } from "react";
import logo from "./logo.svg";
import "./App.css";
import axios from "axios";
import moment from "moment";
import ReactMapGL, { Marker, Popup, NavigationControl } from "react-map-gl";

class App extends Component {
  state = {
    gpsLocations: [],
    viewport: {
      width: 1000,
      height: 1000,
      latitude: 37.7577,
      longitude: -122.4376,
      zoom: 17
    }
  };
  componentWillMount() {
    axios.get("/all").then(response => {
      this.setState({
        gpsLocations: response.data,
        viewport: {
          ...this.state.viewport,
          latitude: response.data[response.data.length - 1].lat,
          longitude: response.data[response.data.length - 1].lng
        }
      });
    });
  }
  render() {
    const { gpsLocations } = this.state;

    return (
      <div>
        <div style={{ float: "left" }}>
          <ReactMapGL
            mapboxApiAccessToken="pk.eyJ1IjoiZnJhY3RpbyIsImEiOiJjam53MG16cmcxZXNpM3ZxcXNjbGJlNnA2In0.s0MG7Na_OojYeU8yDzyQEg"
            {...this.state.viewport}
            onViewportChange={viewport => this.setState({ viewport })}
          >
            {gpsLocations.map((location, index) => (
              <Marker
                key={`marker-index`}
                longitude={location.lng}
                latitude={location.lat}
              >
                X
              </Marker>
            ))}
          </ReactMapGL>
        </div>
        <div style={{ float: "left" }}>FIndihs</div>
      </div>
    );
  }
}

export default App;
