import {
  Box,
  CircularProgress,
  TableContainer,
  Paper,
  TableHead,
  TableRow,
  TableCell,
  TableBody,
  Table,
  Container,
} from "@mui/material";
import { useParams } from "react-router-dom";
import { AppState, appStore } from "../../core/store";
import { Rating } from "../../shared/model";
import { useEffect } from "react";

export const RatingList = () => {
  const params = useParams();
  const loading = appStore((state: AppState) => state.rating.loading);
  const fetchRatings = appStore(
    (state: AppState) => state.rating.fetchRatingsByTargetId
  );
  const ratings = appStore((state: AppState) => state.rating.data);

  useEffect(() => {
    fetchRatings(params.id ?? "");
  }, [fetchRatings, params.id]);
  return (
    <Container>
      <Box sx={{ paddingTop: "10px", paddingBottom: "130px" }}>
        <h2>Ratings</h2>
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
        {!loading && ratings.length > 0 && (
          <TableContainer component={Paper}>
            <Table sx={{ minWidth: 650 }} aria-label="rating table">
              <TableHead>
                <TableRow>
                  <TableCell align="right">Rating</TableCell>
                  <TableCell align="right">Firstname</TableCell>
                  <TableCell align="right">LastName</TableCell>
                  <TableCell align="right">Email</TableCell>
                  <TableCell align="right">Date</TableCell>
                </TableRow>
              </TableHead>
              <TableBody>
                {ratings.map((d: Rating, index: number) => (
                  <TableRow
                    key={index}
                    sx={{ "&:last-child td, &:last-child th": { border: 0 } }}
                  >
                    <TableCell component="th" scope="row">
                      {d.Value}
                    </TableCell>
                    <TableCell align="right">{d.UserFirstName}</TableCell>
                    <TableCell align="right">{d.UserLastName}</TableCell>
                    <TableCell align="right">{d.UserEmail}</TableCell>
                    <TableCell align="right">{new Date(d.LastModified).toLocaleString()}</TableCell>
                  </TableRow>
                ))}
              </TableBody>
            </Table>
          </TableContainer>
        )}
        {!loading && ratings.length === 0 && <div>No ratings to display</div>}
      </Box>
    </Container>
  );
};
