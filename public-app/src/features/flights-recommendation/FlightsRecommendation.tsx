import React, { useState } from 'react';
import { TextField, Grid, Typography, Button, Box } from '@mui/material';
import { DatePicker, LocalizationProvider } from '@mui/x-date-pickers';
import { AdapterDayjs } from '@mui/x-date-pickers/AdapterDayjs';
import dayjs from 'dayjs';
import axios from 'axios';
import { Flight } from '../../shared/model/flight';
import FlightCard from './FlightCard';

interface FlightsRecommendationProps {
  startDate: string;
  endDate: string;
  location: string;
}

interface FlightData {
  date: string
  originLocation: string;
  destinationLocation: string;
}

export const FlightsRecommendation: React.FC<FlightsRecommendationProps> = ({ startDate, endDate, location }) => {
  const [departure, setDeparture] = useState<FlightData>({
    date: dayjs(startDate).format('YYYY-MM-DD'),
    originLocation: '',
    destinationLocation: location
  });
  const [arrival, setArrival] = useState<FlightData>({
    date: dayjs(endDate).format('YYYY-MM-DD'),
    originLocation: location,
    destinationLocation: ''
  });

  const [departureFlights, setDepartureFlights] = useState<Flight[] | undefined>(undefined);
  const [arrivalFlights, setArrivalFlights] = useState<Flight[] | undefined>(undefined);
  const [flightsFlag, setFlightsFlag] = useState(true);

  const handleGetFlights = async () => {
    console.log("departure: ", departure)
    console.log("arrival: ", arrival)
    try {
      const departureResponse = await axios.get(`http://localhost:9000/api/flight/search`, {
        params: {
          date: departure.date,
          destination_location: departure.destinationLocation,
          origin_location: departure.originLocation
        },
      });
      setDepartureFlights(departureResponse.data);
      setFlightsFlag(false)
    } catch (error: any) {
      console.log(error.message);
      return
    }
    try {
      const arrivalResponse = await axios.get(`http://localhost:9000/api/flight/search`, {
        params: {
          date: arrival.date,
          destination_location: arrival.destinationLocation,
          origin_location: arrival.originLocation
        },
      });
      setArrivalFlights(arrivalResponse.data)
      setFlightsFlag(false)
    } catch (error: any) {
      console.log(error.message);
    }
  };

  return (
    <LocalizationProvider dateAdapter={AdapterDayjs}>
      <Grid container spacing={2}>
        <Grid item xs={6}>
          <Typography variant='h5'>
            Departure
          </Typography>
          {flightsFlag && (
            <>
              <Box sx={{ mt: 2 }}>
                <TextField
                  label="Origin"
                  variant="outlined"
                  fullWidth
                  onChange={(event) => setDeparture({ ...departure, originLocation: event.target.value })}
                />
              </Box>
              <Box sx={{ mt: 2 }}>
                <TextField
                  label="Destination"
                  variant="outlined"
                  fullWidth
                  defaultValue={location}
                  onChange={(event) => setDeparture({ ...departure, destinationLocation: event.target.value })}
                />
              </Box>
              <Box sx={{ mt: 2 }}>
                <DatePicker
                  label="Date"
                  defaultValue={dayjs(startDate)}
                  onChange={(date) => setDeparture({ ...departure, date: date ? date.format('YYYY-MM-DD') : '' })}
                />
              </Box>
            </>
          )}
          {!flightsFlag && departureFlights && (
            <>
              {departureFlights.map((flight) => (
                <Box sx={{ mt: 2 }}>
                  <FlightCard
                    id={flight.id}
                    departureDestination={flight.placeSource}
                    arrivalDestination={flight.placeDestination}
                    departureDate={flight.dateSource}
                    arrivalDate={flight.dateDestination}
                    ticketPrice={flight.ticketPrice}
                    remainingTickets={flight.totalTickets - flight.boughtTickets}
                  />
                </Box>
              ))}
            </>
          )}
          {!flightsFlag && !departureFlights && (
            <Box sx={{ mt: 2 }}>
              <Typography>There are no flights available for the specified date and place of departure and arrival</Typography>
            </Box>
          )}
        </Grid>
        <Grid item xs={6}>
          <Typography variant='h5'>
            Arrival
          </Typography>
          {flightsFlag && (
            <>
              <Box sx={{ mt: 2 }}>
                <TextField
                  label="Origin"
                  variant="outlined"
                  fullWidth
                  defaultValue={location}
                  onChange={(event) => setArrival({ ...arrival, originLocation: event.target.value })}
                />
              </Box>
              <Box sx={{ mt: 2 }}>
                <TextField
                  label="Destination"
                  variant="outlined"
                  fullWidth
                  onChange={(event) => setArrival({ ...arrival, destinationLocation: event.target.value })}
                />
              </Box>
              <Box sx={{ mt: 2 }}>
                <DatePicker
                  label="Date"
                  defaultValue={dayjs(endDate)}
                  onChange={(date) => setArrival({ ...arrival, date: date ? date.format('YYYY-MM-DD') : '' })}
                />
              </Box>
            </>
          )}
          {!flightsFlag && arrivalFlights && (
            <>
              {arrivalFlights.map((flight) => (
                <Box sx={{ mt: 2 }}>
                  <FlightCard
                    key={flight.id}
                    id={flight.id}
                    departureDestination={flight.placeSource}
                    arrivalDestination={flight.placeDestination}
                    departureDate={flight.dateSource}
                    arrivalDate={flight.dateDestination}
                    ticketPrice={flight.ticketPrice}
                    remainingTickets={flight.totalTickets - flight.boughtTickets}
                  />
                </Box>
              ))}
            </>
          )}
          {!flightsFlag && !arrivalFlights && (
            <Box sx={{ mt: 2 }}>
              <Typography>There are no flights available for the specified date and place of departure and arrival</Typography>
            </Box>
          )}
        </Grid>
        {flightsFlag && (
          <Grid item xs={12}>
            <Box sx={{ display: 'flex', justifyContent: 'flex-end', marginTop: '1rem' }}>
              <Button variant="contained" color="primary" onClick={handleGetFlights}>
                Get flights
              </Button>
            </Box>
          </Grid>
        )}
      </Grid>
    </LocalizationProvider >
  );
};
