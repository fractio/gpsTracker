import React, { Component } from "react";
import logo from "./logo.svg";
import "./App.css";
import axios from "axios";
import moment from "moment";

import { StaticMap } from "react-map-gl";
import DeckGL, { GeoJsonLayer } from "deck.gl";
import { Location } from "./models";

const MAPBOX_TOKEN =
  "pk.eyJ1IjoiZnJhY3RpbyIsImEiOiJjam53MG16cmcxZXNpM3ZxcXNjbGJlNnA2In0.s0MG7Na_OojYeU8yDzyQEg";

const INITIAL_VIEW_STATE = {
  latitude: 0.42842386927851,
  longitude: -0.10665203462363348,
  zoom: 8,
  minZoom: 2,
  maxZoom: 20
};

type Props = {};

type State = {
  gpsLocations: Array<Location>,
  viewport: any
};
class App extends Component<Props, State> {
  state = {
    gpsLocations: [],
    viewport: {
      width: 1000,
      height: 1000,
      latitude: 51.42842386927851,
      longitude: -0.10665203462363348,
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
    const gpsLocations: Array<Location> = this.state.gpsLocations;

    return (
      <div>
        <div style={{ float: "left" }}>
          {/* <ReactMapGL
            mapboxApiAccessToken="pk.eyJ1IjoiZnJhY3RpbyIsImEiOiJjam53MG16cmcxZXNpM3ZxcXNjbGJlNnA2In0.s0MG7Na_OojYeU8yDzyQEg"
            {...this.state.viewport}
            onViewportChange={viewport => this.setState({ viewport })}
          >
            {gpsLocations.map((location: Location, index) => (
              <Marker
                key={`marker-index`}
                longitude={location.lng}
                latitude={location.lat}
              >
                X
              </Marker>
            ))}
          </ReactMapGL> */}
          <DeckGL
            layers={[
              new GeoJsonLayer({
                id: "geojson",
                data: "/all-geo-json",
                opacity: 1,
                stroked: true,
                filled: false,
                lineWidthMinPixels: 0.5,
                parameters: {
                  depthTest: false
                },

                getLineColor: f => [215, 48, 39],
                getLineWidth: f => 1,

                pickable: false
                // onHover: this._onHover,

                // updateTriggers: {
                //   getLineColor: { year },
                //   getLineWidth: { year }
                // },

                // transitions: {
                //   getLineColor: 1000,
                //   getLineWidth: 1000
                // }
              })
            ]}
            pickingRadius={5}
            initialViewState={INITIAL_VIEW_STATE}
            // viewState={viewState}
            controller={true}
          >
            <StaticMap
              reuseMaps
              mapStyle="mapbox://styles/mapbox/light-v9"
              preventStyleDiffing={true}
              mapboxApiAccessToken={MAPBOX_TOKEN}
            />
          </DeckGL>
        </div>
        <div style={{ float: "left" }}>FIndihs</div>
      </div>
    );
  }
}

export default App;
