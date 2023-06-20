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
import { AppState, appStore } from "../../../core/store";
import { useEffect, useState } from "react";
import { Rating, User } from "../../../shared/model";
import { RatingDialog } from "../../../shared";
import { Link } from "react-router-dom";
import ClearIcon from '@mui/icons-material/Clear';
import CheckIcon from '@mui/icons-material/Check';

const getAverageRating = (a: User) => {
  const sum = a.Ratings?.map((r) => r.Value).reduce(
    (sum, value) => (sum += value),
    0
  );
  return (
    (sum ?? 0) / (a.Ratings && a.Ratings.length > 0 ? a.Ratings.length : 1)
  );
};

const getCurrentUserHostRating = (
  host: User | null,
  currentUser?: User
): Rating | undefined => {
  return host?.Ratings?.find((r) => r.UserId === currentUser?.Id);
};

export const HostList = () => {
  const loading = appStore((state: AppState) => state.user.loading);
  const fetchHosts = appStore((state: AppState) => state.user.fetchHosts);
  const currentUser = appStore((state: AppState) => state.auth.user);
  const hosts = appStore((state: AppState) => state.user.data);
  const [isDialogOpen, setDialogOpen] = useState(false);
  const rateHost = appStore((state: AppState) => state.user.rateHost);
  const removeRating = appStore((state: AppState) => state.rating.removeRating);
  const [selectedHost, setSelectedHost] = useState<User | null>(null);
  const openRatingDialog = (host: User) => {
    setSelectedHost(host);
    setDialogOpen(true);
  };

  useEffect(() => {
    fetchHosts();
  }, [fetchHosts]);

  const handleRating = async (value: number) => {
    //fetching latest data after rating action needs to be replaced by notifications since rating is asynchronous
    try {
      await rateHost({ id: selectedHost?.Id!, value });
      fetchHosts();
    } catch (e) {
      throw e;
    }
  };

  const handleRatingRemoval = async (host: User) => {
    try {
      await removeRating(getCurrentUserHostRating(host, currentUser)?.Id ?? "");
      fetchHosts();
    } catch (e) {
      console.log(e);
    }
  };

  return (
    <div>
      <h2>Hosts</h2>
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
      {!loading && hosts.length > 0 && (
        <TableContainer component={Paper}>
          <Table sx={{ minWidth: 650 }} aria-label="host table">
            <TableHead>
              <TableRow>
                <TableCell align="right">Username</TableCell>
                <TableCell align="right">First Name</TableCell>
                <TableCell align="right">Last Name</TableCell>
                <TableCell align="right">Email</TableCell>
                <TableCell align="right">Place of Living</TableCell>
                <TableCell align="right">Average Rating</TableCell>
                <TableCell align="right">My Rating</TableCell>
                <TableCell align="center">Featured</TableCell>
                <TableCell align="center">Action</TableCell>
              </TableRow>
            </TableHead>
            <TableBody>
              {hosts.map((d: User, index: number) => (
                <TableRow
                  key={index}
                  sx={{ "&:last-child td, &:last-child th": { border: 0 } }}
                >
                  <TableCell component="th" scope="row">
                    {d.Username}
                  </TableCell>
                  <TableCell align="right">{d.FirstName}</TableCell>
                  <TableCell align="right">{d.LastName}</TableCell>
                  <TableCell align="right">{d.Email}</TableCell>
                  <TableCell align="right">{d.LivingPlace}</TableCell>
                  <TableCell align="right">{getAverageRating(d)}</TableCell>
                  <TableCell align="right">
                    {getCurrentUserHostRating(d, currentUser)?.Value}
                  </TableCell>
                  <TableCell align="center">
                    { d.Featured ? <CheckIcon/> : <ClearIcon/> }
                  </TableCell>
                  <TableCell align="right">
                    <Stack direction="row" justifyContent="right" gap={1}>
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
                      <Link to={`/ratings/${d.Id}`}>
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
                    </Stack>
                  </TableCell>
                </TableRow>
              ))}
            </TableBody>
          </Table>
        </TableContainer>
      )}
      {!loading && hosts.length === 0 && <div>No hosts to display</div>}
      <RatingDialog
        open={isDialogOpen}
        setOpen={setDialogOpen}
        onClose={() => setSelectedHost(null)}
        initialRating={
          getCurrentUserHostRating(selectedHost, currentUser)?.Value
        }
        title="Rate host"
        onRate={(value: number) => handleRating(value)}
      />
    </div>
  );
};
