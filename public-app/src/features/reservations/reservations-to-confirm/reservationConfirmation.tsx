import { useEffect, useState } from "react";
import { Reservation } from "../../../shared/model/reservation";
import axios, { AxiosRequestConfig } from "axios";
import { VerifyResponse } from "../../../shared/model/authentication";
import { apiUrl } from "../../../core/store";
import { Box, Button, Card, CardActions, CardContent, CardMedia, CircularProgress, Typography } from "@mui/material";

export const ReservationConfirmation = () => {
  const [loading, setLoading] = useState<number>(1) // to get
  const [reservations, setReservations] = useState<Reservation[]>([]);

  const fetchReservations = async () => {
      const requestOptions = {
        method: 'GET',
        headers: {
          'Content-Type': 'application/json',
          'Authorization': 'Bearer '+window.localStorage.getItem("jwt")
        }
      };
      const response = await fetch(`${apiUrl}` + '/booking/reservation/requests/host', requestOptions);
      const data = await response.json();
      setReservations(data.data);
      setLoading(0)
      console.log(data)
  }

  const handleAccept = async (reservationId : string) => {
      try {
        const requestOptions = {
          method: 'GET',
          headers: {
            'Content-Type': 'application/json',
            'Authorization': 'Bearer '+window.localStorage.getItem("jwt")
          }
        };
        const response = await axios.get(`${apiUrl}/booking/reservation/confirm/${reservationId}`, requestOptions);
        console.log('Reservation request accepted successfully');
        fetchReservations();
      } catch (error) {
        console.error('Error accepting reservation:', error);
      }
    };
  
  useEffect(() => {
      fetchReservations();
  }, []);
  
  return (
      <>
          {loading && (
              <Box
                  sx={{
                      display: "flex",
                      justifyContent: "center",
                      padding: "100px 0px",
                  }}
              >
                  <CircularProgress />
              </Box>
          )}
          {!loading && (
              <>
                  <h2>Reservations confirmation</h2>
                  <div style={{ display: 'flex', flexWrap: 'wrap' }}>
                      {reservations.map((reservation: any, index: number) => (
                          <Card sx={{ minWidth: 250, maxWidth: 350, margin: '1rem' }}>
                              <CardMedia
                                  sx={{ height: 140 }}
                                  image="" />
                              <CardContent>
                              <Typography gutterBottom variant="h4" component="div">
                                        {reservation.accommodationName}
                                    </Typography>
                                    <Typography variant="body2" color="text.secondary">
                                        {reservation.numberOfGuests} guests
                                    </Typography>
                                    <Typography variant="body2" color="text.secondary">
                                        Start at : {reservation.startDate}
                                        End at : {reservation.endDate}
                                    </Typography>
                                    <Typography variant="body2" color="text.secondary">
                                        Location : {reservation.location}
                                    </Typography>
                                  <Typography variant="body2" color="text.secondary">
                                      User canceled reservations {reservation.numberOfCanceledReservations} times in past!
                                  </Typography>
                              </CardContent>
                              <CardActions>
                                  <Button variant="outlined" size="small" onClick={() => handleAccept(reservation.id)}>Accept</Button>
                              </CardActions>
                          </Card>
                      ))}
                  </div>
              </>
          )}
      </>
  );
};