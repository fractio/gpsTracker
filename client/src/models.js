// @flow
// Automatically generated by typewriter. Do not edit.
// http://www.github.com/natdm/typewriter

// Location stores the GPS coords
export type Location = {
  startTime: Time,
  clientTimeStamp: Time,
  serverTimeStamp: Time,
  accuracy: number,
  lat: number,
  lng: number,
  altitude: number,
  speed: number,
  serial: string,
  numberOfSatellites: number,
  direction: number,
  provider: string
};