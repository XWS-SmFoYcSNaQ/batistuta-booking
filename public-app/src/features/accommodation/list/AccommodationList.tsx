import {
  Box,
  CircularProgress,
  TableContainer,
  Paper,
  Table,
  TableHead,
  TableRow,
  TableCell,
  TableBody,
  Stack,
  Button,
} from "@mui/material";
import { useEffect } from "react";
import { useLocation, Outlet } from "react-router";
import { Link } from "react-router-dom";
import { appStore, AppState } from "../../../core/store";

export const AccommodationList = () => {
  const location = useLocation();
  const loading = appStore((state: AppState) => state.accommodation.loading);
  const fetchMyAccommodations = appStore(
    (state: AppState) => state.accommodation.fetchMyAccommodations
  );
  const accommodations = appStore(
    (state: AppState) => state.accommodation.data
  );

  useEffect(() => {
    fetchMyAccommodations();
  }, [fetchMyAccommodations]);
  return (
    <div>
      <h2>Accommodations</h2>
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
                <TableCell align="center">Action</TableCell>
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
                  <TableCell align="right">
                    <Stack direction="row" justifyContent="right" gap={1}>
                      <Link to={`/accommodation/discounts/${d.id}`}>
                        <Button
                          variant="outlined"
                          color="primary"
                          sx={{ whiteSpace: "nowrap" }}
                          type="button"
                        >
                          Discounts
                        </Button>
                      </Link>
                      <Link to={`/accommodation/availability/${d.id}`}>
                        <Button
                          variant="outlined"
                          color="primary"
                          sx={{ whiteSpace: "nowrap" }}
                          type="button"
                        >
                          Availability
                        </Button>
                      </Link>
                    </Stack>
                  </TableCell>
                </TableRow>
              ))}
            </TableBody>
          </Table>
        </TableContainer>
      )}
      {!loading && accommodations.length === 0 && <div>No accommodations to display</div>}
      <Box sx={{ marginTop: "30px" }}>
        {location.pathname === "/accommodation" && (
          <Link to="/accommodation/create">
            <Button type="button" size="large" variant="outlined">
              Create New Accommodation
            </Button>
          </Link>
        )}
        <Outlet />
      </Box>
    </div>
  );
};
