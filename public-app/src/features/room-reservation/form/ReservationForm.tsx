import React, { useState } from 'react';
import {
  Box,
  Button,
  FormControl,
  Grid,
  InputLabel,
  MenuItem,
  Select,
  TextField,
} from '@mui/material';
import { Accommodation } from '../../../shared/model';

interface ReservationFormData {
  checkIn: string;
  checkOut: string;
  guests: number;
}

interface ReservationFormProps {
  onSubmit: (data: ReservationFormData) => void;
}

function ReservationForm(room: Accommodation | null) {

  const availableDates = ['2023-05-15', '2023-05-16', '2023-05-17'];

  const isDateDisabled = (date: string): boolean => {
    return !availableDates.includes(date);
  };

  // const { onSubmit } = props;

  const [checkIn, setCheckIn] = useState('');
  const [checkOut, setCheckOut] = useState('');
  const [guests, setGuests] = useState(0);

  const handleSubmit = (event: React.FormEvent<HTMLFormElement>) => {
    event.preventDefault();
    // onSubmit({ checkIn, checkOut, guests });
  };

  return (
    <form onSubmit={handleSubmit}>
      <Grid container spacing={2}>
        <Grid item xs={12} md={4}>
          <TextField
            fullWidth
            required
            id="check-in"
            label="Check In"
            type="date"
            value={checkIn}
            onChange={(event) => setCheckIn(event.target.value)}
            InputLabelProps={{ shrink: true }}
          />
        </Grid>
        <Grid item xs={12} md={4}>
          <TextField
            fullWidth
            required
            id="check-out"
            label="Check Out"
            type="date"
            value={checkOut}
            onChange={(event) => setCheckOut(event.target.value)}
            InputLabelProps={{ shrink: true }}
          />
        </Grid>
        <Grid item xs={12} md={4}>
          <FormControl fullWidth required>
            <InputLabel id="guests-label">Guests</InputLabel>
            <Select
              labelId="guests-label"
              id="guests"
              value={guests}
              onChange={(event) => setGuests(Number(event.target.value))}
            >
              {[1, 2, 3, 4, 5, 6, 7, 8, 9, 10].map((num) => (
                <MenuItem key={num} value={num}>
                  {num}
                </MenuItem>
              ))}
            </Select>
          </FormControl>
        </Grid>
        <Grid item xs={12}>
          <Box textAlign="center">
            <Button type="submit" variant="contained" color="primary">
              Reserve
            </Button>
          </Box>
        </Grid>
      </Grid>
    </form>
  );
}

export default ReservationForm;
