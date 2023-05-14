import { useEffect } from "react"
import { AppState, appStore } from "../../../core/store"
import ReservationForm from "../form/ReservationForm"
import { useParams } from "react-router-dom"
import { Box, CircularProgress } from "@mui/material"

export const SingleRoom = () => {
  const loading = appStore((state: AppState) => state.accommodation.loading);
  const fetchAccommodationDetails = appStore((state: AppState) => state.accommodation.fetchDetails)
  const room = appStore(
    (state: AppState) => state.accommodation.accommodation
  )
  const params = useParams();

  function handleReservationSubmit() {
  }

  useEffect(() => {
    fetchAccommodationDetails(params.id ?? "", 'periods');
  }, [fetchAccommodationDetails]);

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
          <ReservationForm />
        </>
      )}
    </>
  )
}