import { useEffect, useState } from "react"
import { AppState, apiUrl, appStore } from "../../../core/store"
import { ReservationForm } from "../form/ReservationForm"
import { useParams } from "react-router-dom"
import { Box, Button, CircularProgress, FormControl, Grid, InputLabel, MenuItem, Select } from "@mui/material"
import DateRangePicker, { DateRange } from "rsuite/esm/DateRangePicker"
import axios from "axios"
import { toast } from "react-toastify"
import { format } from 'date-fns';

export const SingleRoom = () => {
  const loading = appStore((state: AppState) => state.accommodation.loading);
  const fetchAccommodationDetails = appStore((state: AppState) => state.accommodation.fetchDetails)
  const room = appStore(
    (state: AppState) => state.accommodation.accommodation
  )
  const periods = appStore(
    (state: AppState) => state.period
  )
  const params = useParams();

	const fetchUser = appStore(
		(state: AppState) => state.auth.verify
	);
	const userId = appStore(
		(state: AppState) => state.auth.userId
	);

  useEffect(() => {
    fetchUser();
  }, [fetchUser]);


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

  const isDisabled = (day : Date) : boolean  => {
      if(room?.periods !== undefined && room?.periods?.length! !== 0)
      {
        for (let i = 0; i < room?.periods?.length!; i++) {
            const startDate = new Date(room?.periods?.at(i)?.start!)
            const endDate = new Date(room?.periods?.at(i)?.end!)
            if(day >= startDate && day <= endDate){
              console.log('darko')
              return true
            }

        }
      }
      return false;
  }

  const createNumberArray = () : number[] => {
    const result = [];
    for (let i = Math.min(room?.minGuests!, room?.maxGuests!); i <= Math.max(room?.minGuests!, room?.maxGuests!); i++) {
      result.push(i);
    }
    return result;
  }

  useEffect(() => {
    fetchAccommodationDetails(params.id ?? "");
  }, [fetchAccommodationDetails]);

  const handleSubmit = async (event: React.FormEvent<HTMLFormElement>) => {
    event.preventDefault();
    const startDate = format(selectedDates[0],'yyyy-MM-dd HH:mm:ssxxx');

    const endDate = format(selectedDates[1],'yyyy-MM-dd HH:mm:ssxxx');
    let requestbody = {
      accommodationId: room?.id, startDate: startDate, endDate: endDate, numberOfGuests: guests, userId: userId
  }
    try {
      await axios.post(`${apiUrl}/booking/request`, requestbody)
      toast.success("Booking request successfully sent!")
    } catch (e: any) {
      if(e.message){
        throw new Error(e.message)
      }
      throw new Error("Error while creating accommodation.")
    }
  }

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
          <h1> { room?.name } </h1>
          <p>{ room?.basePrice }€ per night </p>
          <p>Total: </p>
          {/* {isAuthenticated() ? <Link to="/books">Books</Link> : null} */}
      {<form onSubmit={handleSubmit}>
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
              required
              labelId="guests-label"
              id="guests"
              value={guests}
              onChange={(event) => setGuests(Number(event.target.value))}
            >
              {createNumberArray().map((num : number) => (
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
    </form>}
        </>
      )}
    </>
  )
}