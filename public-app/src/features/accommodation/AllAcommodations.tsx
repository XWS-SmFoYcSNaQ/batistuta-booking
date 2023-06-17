import * as React from "react";
import Box from '@mui/material/Box';
import Card from '@mui/material/Card';
import CardActions from '@mui/material/CardActions';
import CardContent from '@mui/material/CardContent';
import Button from '@mui/material/Button';
import Typography from '@mui/material/Typography';
import { Accommodation } from "../../shared/model";
import { AppState, appStore } from "../../core/store";
import { useEffect, useState } from "react";
import { CircularProgress, Paper, Table, TableBody, TableCell, TableContainer, TableHead, TableRow, TextField } from "@mui/material";
import { DateRangePicker } from 'rsuite';
import { DateRange } from "rsuite/esm/DateRangePicker";
import axios from "axios";

export const AllAccommodations = () => {
	const loading = appStore((state: AppState) => state.accommodation.loading);
    const fetchSearchedAccommodations = appStore(
        (state: AppState) => state.accommodation.fetchSearchedAccommodations
      );
    const fetchAccommodations = appStore(
    (state: AppState) => state.accommodation.fetchAccommodations
    );
    const accommodations = appStore(
    (state: AppState) => state.accommodation.data
    );
    
      useEffect(() => {
        fetchAccommodations(null);
      }, [fetchAccommodations]);


    const [location, setLocation] = useState("");
    const [numberOfGuests, setNumberOfGuests] = useState(0);
    const [selectedDates, setSelectedDates] = useState<[Date, Date]>([
        new Date(),
        new Date()
      ]);

    const handleDateRangeChange = (value: DateRange | null) => {
        if (value !== null) {
          setSelectedDates([value[0], value[1]]);
        }
      };

      const handleSearch = () => {
        let requestbody = {
            numberOfGuests: numberOfGuests, start: selectedDates[0].toISOString(), end: selectedDates[1].toISOString(), location: location
        }
        fetchSearchedAccommodations(requestbody)
      };

    return (
        <div>
                <Box sx={{ marginTop: "30px", marginLeft: "30px" }}>
                <TextField
          required
          label="Location"
          value={location}
          onChange={(event) => setLocation(event.target.value)}
          />
        <TextField
          required
          type="number"
          label="Number of guests"
          value={numberOfGuests}
          onChange={(event) => setNumberOfGuests(parseInt(event.target.value))}
        />
              <DateRangePicker
        style={{ width: 300 }}
        placeholder="Select Date Range"
        title="Select Date Range"
        onChange={handleDateRangeChange}
        value={selectedDates}
      />
        <Button size="large" color="primary" type="button" onClick={handleSearch}>
              Search
        </Button>
      </Box>
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
          {!loading && accommodations.length > 0 && (
            <TableContainer component={Paper}>
              <Table sx={{ minWidth: 650 }} aria-label="accommodation table">
                <TableHead>
                  <TableRow>
                    <TableCell align="right">Name</TableCell>
                    <TableCell align="right">Price</TableCell>
                    <TableCell align="right">Benefits</TableCell>
                    <TableCell align="right">Min Guests</TableCell>
                    <TableCell align="right">Max Guests</TableCell>
                    <TableCell align="center">Location</TableCell>
                  </TableRow>
                </TableHead>
                <TableBody>
                  {accommodations.map((d: any, index: number) => (
                    <TableRow
                      key={index}
                      sx={{ "&:last-child td, &:last-child th": { border: 0 } }}
                    >
                      <TableCell component="th" scope="row">
                        {d.name}
                      </TableCell>
                      <TableCell align="right">{d.basePrice}&nbsp;EUR</TableCell>
                      <TableCell align="right">{d.benefits}</TableCell>
                      <TableCell align="right">{d.minGuests}</TableCell>
                      <TableCell align="right">{d.maxGuests}</TableCell>
                      <TableCell align="right">{d.location}</TableCell>
                    </TableRow>
                  ))}
                </TableBody>
              </Table>
            </TableContainer>
          )}
        </div>
      );
                  }