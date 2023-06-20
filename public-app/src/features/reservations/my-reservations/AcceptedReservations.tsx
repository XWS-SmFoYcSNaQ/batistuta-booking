import { useEffect, useState } from 'react';
import { Box, CircularProgress } from '@mui/material';
import { Reservation } from '../../../shared/model/reservation';
import { VerifyResponse } from '../../../shared/model/authentication';
import { apiUrl } from '../../../core/store';
import axios, { AxiosRequestConfig } from 'axios';
import { toast } from 'react-toastify';
import ReservationCard from './ReservationCard';

export const AcceptedReservations = () => {
  const [loading, setLoading] = useState<number>(1)
  const [acceptedReservations, setAcceptedReservations] = useState<Reservation[]>([]);

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
    const response = await fetch(`${apiUrl}` + '/booking/reservation/user/' + userID, requestOptions);
    const data = await response.json();
    setAcceptedReservations(data.data);
    setLoading(0)
  }

  const handleCancel = async (reservationId: string) => {
    try {
      await axios.delete(`${apiUrl}/booking/reservation/${reservationId}`);
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
            {acceptedReservations.map((reservation: any) => (
              <ReservationCard reservation={reservation} onCancel={handleCancel} key={reservation.id} />
            ))}
          </div>
        </>
      )}
    </>
  );
}