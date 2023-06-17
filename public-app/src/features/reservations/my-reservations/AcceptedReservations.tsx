import Card from '@mui/material/Card';
import CardActions from '@mui/material/CardActions';
import CardContent from '@mui/material/CardContent';
import CardMedia from '@mui/material/CardMedia';
import Button from '@mui/material/Button';
import Typography from '@mui/material/Typography';
import { useEffect, useState } from 'react';
import { Box, CircularProgress } from '@mui/material';
import { Link } from 'react-router-dom';
import { Reservation } from '../../../shared/model/reservation';
import { VerifyResponse } from '../../../shared/model/authentication';
import React from 'react';
import { apiUrl, appStore, AppState } from '../../../core/store';
import axios, { AxiosRequestConfig } from 'axios';
import { toast } from 'react-toastify';

export const AcceptedReservations = () => {
    const [loading, setLoading] = useState<number>(1) // to get
    const [acceptedReservations, setAcceptedReservations] = useState<Reservation[]>([]);

    const fetchReservations = async () => {
        const verifyConfig : AxiosRequestConfig = {
            headers: {
              'Content-Type': 'application/json',
              'Authorization': 'Bearer '+window.localStorage.getItem("jwt")
            }
          }
        const resp = await axios.post<VerifyResponse>(`${apiUrl}/api/auth/verify`,{},  verifyConfig);
        let userID = resp.data.UserId
        const requestOptions = {
          method: 'GET'
        };
        const response = await fetch(`${apiUrl}` + '/booking/reservation/user/' + userID  , requestOptions);
        const data = await response.json();
        setAcceptedReservations(data.data);
        setLoading(0)
        console.log(data)
    }

    const handleCancel = async (reservationId: string) => {
        try {
          const response = await axios.delete(`${apiUrl}/booking/reservation/${reservationId}`);
          toast.success('Reservation canceled successfully');
          fetchReservations();
        } catch (error: any) {
          if (error.response && error.response.data && error.response.data.code === 3) {
            toast.error(error.response.data.message);
          } else {
            toast.error('Error canceling reservation');
          }
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
                    <h2>Accepted reservations</h2>
                    <div style={{ display: 'flex', flexWrap: 'wrap' }}>
                        {acceptedReservations.map((reservation: any, index: number) => (
                            <Card sx={{ minWidth: 250, maxWidth: 350, margin: '1rem' }}>
                                <CardMedia
                                    sx={{ height: 140 }}
                                    image="" />
                                <CardContent>
                                    <Typography gutterBottom variant="h4" component="div">
                                        Start at : {reservation.startDate}
                                        End at : {reservation.endDate}
                                    </Typography>
                                    <Typography variant="body2" color="text.secondary">
                                        {reservation.numberOfGuests} guests
                                    </Typography>
                                </CardContent>
                                <CardActions>
                                    <Button variant="outlined" size="small" onClick={() => handleCancel(reservation.id)}>Cancel</Button>
                                </CardActions>
                            </Card>
                        ))}
                    </div>
                </>
            )}
        </>
    );
}