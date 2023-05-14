import ReservationForm from "../form/ReservationForm"

export const SingleRoom = () => {

  function handleReservationSubmit() {
  }

  return <div>
    <ReservationForm onSubmit={handleReservationSubmit} />
  </div>
}