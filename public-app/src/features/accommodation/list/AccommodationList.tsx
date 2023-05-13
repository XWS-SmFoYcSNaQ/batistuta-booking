import { Box, CircularProgress, TableContainer, Paper, Table, TableHead, TableRow, TableCell, TableBody, Stack, Button } from "@mui/material";
import { useEffect } from "react";
import { useLocation, Outlet } from "react-router";
import { Link } from "react-router-dom";
import { appStore, AppState } from "../../../core/store";

export const AccommodationList = () => {
  const location = useLocation();
  const loading = appStore((state: AppState) => state.accommodation.loading);
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
        {!loading && (
          <TableContainer component={Paper}>
            <Table sx={{ minWidth: 650 }} aria-label="accommodation table">
              <TableHead>
                <TableRow>
                  <TableCell align="right">Name</TableCell>
                  <TableCell align="right">Price</TableCell>
                  <TableCell align="right">Benefits</TableCell>
                  <TableCell align="right">Max Guests</TableCell>
                  <TableCell align="right">Min Guests</TableCell>
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
                      <Stack direction="row">
                        <Button
                          variant="outlined"
                          color="error"
                          sx={{ mr: 3 }}
                          type="button"
                        >
                          Delete
                        </Button>
                        <Link to={`/accommodation/${d.id}`}>
                          <Button
                            variant="outlined"
                            color="primary"
                            sx={{ whiteSpace: "nowrap" }}
                            type="button"
                          >
                            Details
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
  )
}