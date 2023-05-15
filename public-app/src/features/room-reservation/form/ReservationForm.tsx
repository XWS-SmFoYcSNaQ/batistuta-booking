import React, { useEffect, useState } from 'react';
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
import { Accommodation, Period } from '../../../shared/model';
import DateRangePicker, { DateRange } from 'rsuite/esm/DateRangePicker';
import { AppState, appStore } from '../../../core/store';

interface ReservationFormData {
  checkIn: string;
  checkOut: string;
  guests: number;
}

interface ReservationFormProps {
  onSubmit: (data: ReservationFormData) => void;
}

export const ReservationForm = (room: any, periods: Period[]) => {

  const availableDates = ['2023-05-15', '2023-05-16', '2023-05-17'];

  const isDateDisabled = (date: string) => {
  };

  const [guests, setGuests] = useState(0);
  const [selectedDates, setSelectedDates] = useState<[Date, Date]>([
    new Date(),
    new Date()
  ]);

const handleDateRangeChange = (value: DateRange | null) => {
    if (value !== null) {
      setSelectedDates([value[0], value[1]]);
    }
  };

  const handleSubmit = (event: React.FormEvent<HTMLFormElement>) => {
    event.preventDefault();
    // onSubmit({ checkIn, checkOut, guests });
  };

  const isDisabled = (day : Date) : boolean  => {
    console.log(periods.length)
      if(periods !== undefined && periods.length !== 0)
      {
        for (let i = 0; i < periods.length; i++) {
            const startDate = new Date(periods[i].start!)
            const endDate = new Date(periods[i].end!)
            if(day.getDate > startDate.getDate && day.getDate < endDate.getDate)
              return true

        }
      }
      return false;
  }

  return (
    <form onSubmit={handleSubmit}>
      <Grid container spacing={2}>
        <Grid item xs={12} md={4}>
        <DateRangePicker
        style={{ width: 300 }}
        placeholder="Select Date Range"
        title="Select Date Range"
        onChange={handleDateRangeChange}
        value={selectedDates}
        shouldDisableDate={(day : Date) => {
          var now = new Date();
          if(day < now)
            return true;
          if(isDisabled(day))
            return true
          return false;
        }}
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
