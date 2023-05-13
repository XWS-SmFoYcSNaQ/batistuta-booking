import { Outlet } from "react-router";
import { AppState, appStore } from "../../core/store";
import { useEffect } from "react";
import {
  Box,
  Paper,
  Table,
  TableBody,
  TableCell,
  TableContainer,
  TableHead,
  TableRow,
} from "@mui/material";

export const AccommodationRoot = () => {
  const fetchAccommodations = appStore(
    (state: AppState) => state.accommodation.fetchAccommodations
  );
  const accommodations = appStore(
    (state: AppState) => state.accommodation.data
  );

  useEffect(() => {
    fetchAccommodations();
  }, [fetchAccommodations]);

  return (
    <Box sx={{ padding: "30px" }}>
      <div>
        <h3>Accommodations</h3>
        <TableContainer component={Paper}>
          <Table sx={{ minWidth: 650 }} aria-label="simple table">
            <TableHead>
              <TableRow>
                <TableCell align="right">Name</TableCell>
                <TableCell align="right">Price</TableCell>
                <TableCell align="right">Benefits</TableCell>
                <TableCell align="right">Max Guests</TableCell>
                <TableCell align="right">Min Guests</TableCell>
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
                </TableRow>
              ))}
            </TableBody>
          </Table>
        </TableContainer>
      </div>
      <Outlet />
    </Box>
  );
};
