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
import { useEffect, useState } from "react";
import { useLocation, Outlet } from "react-router";
import { Link } from "react-router-dom";
import { appStore, AppState } from "../../../core/store";
import { RatingDialog } from "../../../shared";
import { Accommodation, User } from "../../../shared/model";

const getAverageRating = (a: Accommodation) => {
  const sum = a.ratings
    ?.map((r) => r.value)
    .reduce((sum, value) => (sum += value), 0);
  return (
    (sum ?? 0) / (a.ratings && a.ratings.length > 0 ? a.ratings.length : 1)
  );
};

const getCurrentUserAccommodationRating = (
  a: Accommodation | null,
  currentUser?: User
) => {
  return a?.ratings?.find((r) => r.userId === currentUser?.Id)?.value;
};

export const AccommodationList = ({ host = true }: { host?: boolean }) => {
  const location = useLocation();
  const loading = appStore((state: AppState) => state.accommodation.loading);
  const fetchMyAccommodations = appStore(
    (state: AppState) => state.accommodation.fetchMyAccommodations
  );
  const fetchAccommodations = appStore(
    (state: AppState) => state.accommodation.fetchAccommodations
  );
  const accommodations = appStore(
    (state: AppState) => state.accommodation.data
  );
  const currentUser = appStore((state: AppState) => state.auth.user);
  const [isDialogOpen, setDialogOpen] = useState(false);
  const rateAccommodation = appStore(
    (state: AppState) => state.accommodation.rateAccommodation
  );
  const [selectedAccommodation, setSelectedAccommodation] =
    useState<Accommodation | null>(null);

  useEffect(() => {
    host ? fetchMyAccommodations() : fetchAccommodations();
  }, [fetchAccommodations, fetchMyAccommodations, host]);

  const openRatingDialog = (accommodation: Accommodation) => {
    setSelectedAccommodation(accommodation);
    setDialogOpen(true);
  };
  const handleRating = async (value: number) => {
    //fetching latest data after rating action needs to be replaced by notifications since rating is asynchronous
    try {
      await rateAccommodation({ id: selectedAccommodation?.id!, value });
      host ? fetchMyAccommodations() : fetchAccommodations();
    } catch (e) {
      throw e;
    }
  };

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
                <TableCell align="right">Average Rating</TableCell>
                <TableCell align="right">My Rating</TableCell>
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
                  <TableCell align="right">{getAverageRating(d)}</TableCell>
                  <TableCell align="right">
                    {getCurrentUserAccommodationRating(d, currentUser)}
                  </TableCell>
                  <TableCell align="right">
                    <Stack direction="row" justifyContent="right" gap={1}>
                      <Link to={`/accommodation/discounts/${d.id}`}>
                        <Button
                          variant="outlined"
                          color="primary"
                          sx={{ whiteSpace: "nowrap" }}
                          type="button"
                          size="small"
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
                          size="small"
                        >
                          Availability
                        </Button>
                      </Link>
                      <Button
                        variant="outlined"
                        color="primary"
                        sx={{ whiteSpace: "nowrap" }}
                        type="button"
                        onClick={() => openRatingDialog(d)}
                        size="small"
                      >
                        Rate
                      </Button>
                    </Stack>
                  </TableCell>
                </TableRow>
              ))}
            </TableBody>
          </Table>
        </TableContainer>
      )}
      {!loading && accommodations.length === 0 && (
        <div>No accommodations to display</div>
      )}
      <Box sx={{ marginTop: "30px" }}>
        {location.pathname === "/accommodation/my" &&
          currentUser?.Role === 1 && (
            <Link to="/accommodation/my/create">
              <Button type="button" size="large" variant="outlined">
                Create New Accommodation
              </Button>
            </Link>
          )}
        <Outlet />
      </Box>
      <RatingDialog
        open={isDialogOpen}
        setOpen={setDialogOpen}
        onClose={() => setSelectedAccommodation(null)}
        initialRating={getCurrentUserAccommodationRating(
          selectedAccommodation,
          currentUser
        )}
        title="Rate accommodation"
        onRate={(value: number) => handleRating(value)}
      />
    </div>
  );
};
