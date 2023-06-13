import Card from '@mui/material/Card';
import CardActions from '@mui/material/CardActions';
import CardContent from '@mui/material/CardContent';
import CardMedia from '@mui/material/CardMedia';
import Button from '@mui/material/Button';
import Typography from '@mui/material/Typography';
import { useEffect, useState } from 'react';
import { Box, CircularProgress } from '@mui/material';
import { Link } from 'react-router-dom';
import { Reservation } from '../../shared/model/reservation';
import { VerifyResponse } from '../../shared/model/authentication';
import React from 'react';
import { apiUrl, appStore, AppState } from '../../core/store';
import axios, { AxiosRequestConfig } from 'axios';

export const PendingReservations = () => {
    const [loading, setLoading] = useState<number>(1) // to get
    const [pendingReservations, setPendingReservations] = useState<Reservation[]>([]);
    const verify = appStore((state: AppState) => state.auth.verify)
    const userId = appStore(
		(state: AppState) => state.auth.userId
	);

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
        const response = await fetch(`${apiUrl}` + '/booking/user/' + userID  , requestOptions);
        const data = await response.json();
        setPendingReservations(data.data);
        setLoading(0)
        console.log(data)
    }

    const handleCancel = async (reservationId : string) => {
        try {
          const response = await axios.delete(`${apiUrl}/booking/request/${reservationId}`);
          console.log('Reservation canceled successfully');
          fetchReservations();
        } catch (error) {
          console.error('Error canceling reservation:', error);
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
                    <h2>Pending reservations</h2>
                    <div style={{ display: 'flex', flexWrap: 'wrap' }}>
                        {pendingReservations.map((reservation: any, index: number) => (
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
};