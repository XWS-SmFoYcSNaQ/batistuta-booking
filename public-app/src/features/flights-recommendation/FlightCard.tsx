import React, { useState } from 'react';
import { Box, Button, Card, CardContent, Grid, Modal, TextField, Typography } from '@mui/material';
import axios from 'axios';
import { toast } from 'react-toastify';

interface CardProps {
  id: string;
  departureDestination: string;
  arrivalDestination: string;
  departureDate: string;
  arrivalDate: string;
  ticketPrice: number;
  remainingTickets: number;
}

const FlightCard: React.FC<CardProps> = ({
  id,
  departureDestination,
  arrivalDestination,
  departureDate,
  arrivalDate,
  ticketPrice,
  remainingTickets
}) => {
  const [isModalOpen, setModalOpen] = useState(false);
  const [apiKey, setApiKey] = useState('');

  const handleOpenModal = () => {
    setModalOpen(true);
  };

  const handleCloseModal = () => {
    setModalOpen(false);
  };
  const handleBookFlight = async () => {
    try {
      const arrivalResponse = await axios.post(`http://localhost:9000/api/tickets`, {
        flightId: id,
        quantity: 1
      },
      {
        params: {
          api_key: apiKey
        }
      })

      if (arrivalResponse.status === 201) {
        toast.success("Flight booked successfully!")
      }
      handleCloseModal();
    } catch (error: any) {
      if (error.response.request.status === 401) {
        toast.error("Invalid API Key, please try again...")
      } else {
        toast.error("Airline server responded with error: " + error.response.data.error);
        handleCloseModal();
      }
    }
  };

  const departureDateObj = new Date(departureDate);
  const formattedDepartureDate = `${departureDateObj.getFullYear()}-${departureDateObj.getMonth() + 1}-${departureDateObj.getDate()} ${departureDateObj.getHours() - 2}:${departureDateObj.getMinutes()}`;

  const arrivalDateObj = new Date(arrivalDate);
  const formattedArrivalDate = `${arrivalDateObj.getFullYear()}-${arrivalDateObj.getMonth() + 1}-${arrivalDateObj.getDate()} ${arrivalDateObj.getHours() - 2}:${arrivalDateObj.getMinutes()}`;

  return (
    <>
      <Card>
        <CardContent>
          <Typography variant="h6">{departureDestination} &#8594; {arrivalDestination}</Typography>
          <Grid container spacing={1} sx={{ marginTop: '0.2rem' }}>
            <Grid item xs={5}>
              <Typography>Departure date:</Typography>
            </Grid>
            <Grid item xs={5}>
              <Typography>{formattedDepartureDate}</Typography>
            </Grid>
          </Grid>
          <Grid container spacing={1} sx={{ marginTop: '0.2rem' }}>
            <Grid item xs={5}>
              <Typography>Arrival date:</Typography>
            </Grid>
            <Grid item xs={5}>
              <Typography>{formattedArrivalDate}</Typography>
            </Grid>
          </Grid>
          <Grid container spacing={1} sx={{ marginTop: '0.2rem' }}>
            <Grid item xs={5}>
              <Typography>Ticket price:</Typography>
            </Grid>
            <Grid item xs={5}>
              <Typography>{ticketPrice}&#8364;</Typography>
            </Grid>
          </Grid>
          <Grid container spacing={1} sx={{ marginTop: '0.2rem' }}>
            <Grid item xs={5}>
              <Typography>Tickets left:</Typography>
            </Grid>
            <Grid item xs={5}>
              <Typography>{remainingTickets}</Typography>
            </Grid>
          </Grid>
        </CardContent>
        <Box
          sx={{
            display: 'flex',
            justifyContent: 'flex-end',
            paddingBottom: '1rem',
            paddingRight: '1rem'
          }}
        >
          <Button variant="contained" color="primary" onClick={(handleOpenModal)}>
            Book Flight
          </Button>
        </Box>
      </Card>
      <Modal open={isModalOpen} onClose={handleCloseModal}>
        <Box
          sx={{
            position: 'absolute',
            top: '50%',
            left: '50%',
            transform: 'translate(-50%, -50%)',
            backgroundColor: 'white',
            boxShadow: 24,
            padding: '2rem',
            width: '600px',
            borderRadius: '8px'
          }}
        >
          <Typography variant="h6">Enter API Key from your Airline account:</Typography>
          <TextField
            label="API Key"
            variant="outlined"
            fullWidth
            value={apiKey}
            onChange={(event) => setApiKey(event.target.value)}
            sx={{ marginTop: '1rem' }}
          />
          <Box sx={{ display: 'flex', justifyContent: 'flex-end', marginTop: '1rem' }}>
            <Button variant="contained" color="primary" onClick={handleBookFlight}>
              Confirm Booking
            </Button>
          </Box>
        </Box>
      </Modal>
    </>
  );
};

export default FlightCard;
