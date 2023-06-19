import { useEffect, useState } from 'react';
import { Box, CircularProgress } from '@mui/material';
import { Reservation } from '../../../shared/model/reservation';
import { VerifyResponse } from '../../../shared/model/authentication';
import { apiUrl, appStore, AppState } from '../../../core/store';
import axios, { AxiosRequestConfig } from 'axios';
import ReservationCard from './ReservationCard';

export const PendingReservations = () => {
    const [loading, setLoading] = useState<number>(1)
    const [pendingReservations, setPendingReservations] = useState<Reservation[]>([]);
    const verify = appStore((state: AppState) => state.auth.verify)
    const userId = appStore(
        (state: AppState) => state.auth.userId
    );

    const fetchReservations = async () => {
        const verifyConfig: AxiosRequestConfig = {
            headers: {
                'Content-Type': 'application/json',
                'Authorization': 'Bearer ' + window.localStorage.getItem("jwt")
            }
        }
        const resp = await axios.post<VerifyResponse>(`${apiUrl}/api/auth/verify`, {}, verifyConfig);
        let userID = resp.data.UserId
        const requestOptions = {
            method: 'GET'
        };
        const response = await fetch(`${apiUrl}` + '/booking/user/' + userID, requestOptions);
        const data = await response.json();
        setPendingReservations(data.data);
        setLoading(0)
    }

    const handleCancel = async (reservationId: string) => {
        try {
            await axios.delete(`${apiUrl}/booking/request/${reservationId}`);
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
                        {pendingReservations.map((reservation: any) => (
                            <ReservationCard reservation={reservation} onCancel={handleCancel} key={reservation.id} />
                        ))}
                    </div>
                </>
            )}
        </>
    );
};