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
  TextField,
  Grid,
} from "@mui/material";
import { useCallback, useEffect, useState } from "react";
import { useLocation, Outlet } from "react-router";
import { Link } from "react-router-dom";
import { appStore, AppState } from "../../../core/store";
import { RatingDialog } from "../../../shared";
import { Accommodation, User } from "../../../shared/model";
import { DateRangePicker } from "rsuite";
import { DateRange } from "rsuite/esm/DateRangePicker";
import { AccommodationFilter, Filters } from "../filters";

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
  return a?.ratings?.find((r) => r.userId === currentUser?.Id);
};

export const AccommodationList = ({ host = true }: { host?: boolean }) => {
  const location = useLocation();
  const loading = appStore((state: AppState) => state.accommodation.loading);
  const fetchSearchedAccommodations = appStore(
    (state: AppState) => state.accommodation.fetchSearchedAccommodations
  );
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
  const removeRating = appStore((state: AppState) => state.rating.removeRating);
  const [selectedAccommodation, setSelectedAccommodation] =
    useState<Accommodation | null>(null);
  const [filtersEnabled, setFiltersEnabled] = useState(false);
  const [filters, setFilters] = useState<AccommodationFilter | null>(null);

  const fetchData = useCallback(() => {
    host
      ? fetchMyAccommodations()
      : fetchAccommodations(filtersEnabled ? filters : null);
  }, [
    fetchAccommodations,
    fetchMyAccommodations,
    filters,
    filtersEnabled,
    host,
  ]);

  const [locationSearch, setLocationSearch] = useState("");
  const [numberOfGuests, setNumberOfGuests] = useState(0);
  const [selectedDates, setSelectedDates] = useState<[Date, Date]>([
    new Date(),
    new Date()
  ]);

  useEffect(() => {
    fetchData();
  }, [fetchData]);

  useEffect(() => {
    if (!filtersEnabled) setFilters(null);
  }, [filtersEnabled]);

  const openRatingDialog = (accommodation: Accommodation) => {
    setSelectedAccommodation(accommodation);
    setDialogOpen(true);
  };
  const handleRating = async (value: number) => {
    //fetching latest data after rating action needs to be replaced by notifications since rating is asynchronous
    try {
      await rateAccommodation({ id: selectedAccommodation?.id!, value });
      fetchData();
    } catch (e) {
      throw e;
    }
  };
  const handleRatingRemoval = async (accommodation: Accommodation) => {
    try {
      await removeRating(
        getCurrentUserAccommodationRating(accommodation, currentUser)?.id ?? ""
      );
      fetchData();
    } catch (e) {
      console.log(e);
    }
  };

  const handleDateRangeChange = (value: DateRange | null) => {
    if (value !== null) {
      setSelectedDates([value[0], value[1]]);
    }
  };

  const handleSearch = () => {
    let requestbody = {
      numberOfGuests: numberOfGuests, start: selectedDates[0].toISOString(), end: selectedDates[1].toISOString(), location: locationSearch
    }
    fetchSearchedAccommodations(requestbody)
  };

  return (
    <div>
      <h2>Accommodations</h2>
      <div>
        <Grid container spacing={2} alignItems="center" sx={{ marginTop: "30px", marginBottom: "30px" }}>
          <Grid item>
            <TextField
              required
              label="Location"
              value={locationSearch}
              onChange={(event) => setLocationSearch(event.target.value)}
            />
          </Grid>
          <Grid item>
            <TextField
              required
              type="number"
              label="Number of guests"
              value={numberOfGuests}
              onChange={(event) => setNumberOfGuests(parseInt(event.target.value))}
            />
          </Grid>
          <Grid item>
            <DateRangePicker
              placeholder="Select Date Range"
              title="Select Date Range"
              onChange={handleDateRangeChange}
              value={selectedDates}
            />
          </Grid>
          <Grid item>
            <Button
              size="large"
              color="primary"
              type="button"
              onClick={handleSearch}
            >
              Search
            </Button>
          </Grid>
          <Grid item>
            <Button
              size="large"
              color="error"
              type="button"
              onClick={fetchData}
            >
              Reset
            </Button>
          </Grid>
        </Grid>
      </div>
      {!host && !filtersEnabled && (
        <Button type="button" onClick={() => setFiltersEnabled(true)}>
          Enable filters
        </Button>
      )}
      {!host && filtersEnabled && (
        <Filters
          setFiltersEnabled={(value: boolean) => setFiltersEnabled(value)}
          setFilters={(filters: AccommodationFilter | null) =>
            setFilters(filters)
          }
        />
      )}
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
                <TableCell align="right">Location</TableCell>
                <TableCell align="right">Automatic Acceptance Of Reservations</TableCell>
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
                  <TableCell align="right">{d.basePrice}&nbsp;&#8364;</TableCell>
                  <TableCell align="right">{d.benefits}</TableCell>
                  <TableCell align="right">{d.minGuests}</TableCell>
                  <TableCell align="right">{d.maxGuests}</TableCell>
                  <TableCell align="right">{getAverageRating(d)}</TableCell>
                  <TableCell align="right">
                    {getCurrentUserAccommodationRating(d, currentUser)?.value}
                  </TableCell>
                  <TableCell align="right">{d.location}</TableCell>
                  <TableCell align="right">
                    <span style={{ color: d.automaticReservation === 1 ? "green" : "red" }}>{d.automaticReservation === 1 ? "enabled" : "disabled"}</span>
                  </TableCell>
                  <TableCell align="right">
                    <Stack
                      flexWrap="wrap"
                      direction="row"
                      justifyContent="right"
                      gap={1}
                    >
                      <Button
                        variant="outlined"
                        color="error"
                        sx={{ whiteSpace: "nowrap" }}
                        type="button"
                        onClick={() => handleRatingRemoval(d)}
                        size="small"
                      >
                        Remove Rating
                      </Button>
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
                      <Link to={`/ratings/${d.id}`}>
                        <Button
                          variant="outlined"
                          color="primary"
                          sx={{ whiteSpace: "nowrap" }}
                          type="button"
                          size="small"
                        >
                          Ratings
                        </Button>
                      </Link>
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
        initialRating={
          getCurrentUserAccommodationRating(selectedAccommodation, currentUser)
            ?.value
        }
        title="Rate accommodation"
        onRate={(value: number) => handleRating(value)}
      />
    </div>
  );
};
