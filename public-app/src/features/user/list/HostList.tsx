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
import { User } from "../../../shared/model";
import { RatingDialog } from "../../../shared";

export const HostList = () => {
  const loading = appStore((state: AppState) => state.user.loading);
  const fetchUsers = appStore((state: AppState) => state.user.fetchUsers);
  const currentUser = appStore((state: AppState) => state.auth.user);
  const hosts = appStore((state: AppState) => state.user.data).filter(
    (user) => user.Role === 1 || (user.Role as any) === "Host"
  );
  const [isDialogOpen, setDialogOpen] = useState(false);
  const [selectedHost, setSelectedHost] = useState<User | null>(null);
  const openRatingDialog = (host: User) => {
    setSelectedHost(host);
    setDialogOpen(true);
  };

  useEffect(() => {
    fetchUsers();
  }, [fetchUsers]);

  const handleRating = async (value: number) => {
    //fetching latest data after rating action needs to be replaced by notifications since rating is asynchronous
    console.log(value, selectedHost?.Id);
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
                  <TableCell align="right">
                    <Stack direction="row" justifyContent="right" gap={1}>
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
      {!loading && hosts.length === 0 && <div>No hosts to display</div>}
      <RatingDialog
        open={isDialogOpen}
        setOpen={setDialogOpen}
        onClose={() => setSelectedHost(null)}
        // initialRating={getCurrentUserAccommodationRating(
        //   selectedAccommodation,
        //   currentUser
        // )}
        title="Rate host"
        onRate={(value: number) => handleRating(value)}
      />
    </div>
  );
};
