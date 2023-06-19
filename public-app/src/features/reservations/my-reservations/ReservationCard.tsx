import { Box, Button, Card, CardActions, CardContent, CardMedia, Modal, Typography } from "@mui/material";
import { useState } from "react";
import { FlightsRecommendation } from "../../flights-recommendation";

interface ReservationCardProps {
  reservation: any;
  onCancel: (id: string) => void;
}

const ReservationCard: React.FC<ReservationCardProps> = ({ reservation, onCancel }) => {
  const [openModal, setOpenModal] = useState(false);

  const handleOpenModal = () => {
    setOpenModal(true);
  };

  const handleCloseModal = () => {
    setOpenModal(false);
  };

  return (
    <Card sx={{ minWidth: 250, maxWidth: 350, margin: '1rem' }}>
      <CardMedia sx={{ height: 140 }} image="" />
      <CardContent>
      <Typography gutterBottom variant="h5" component="div">
          &nbsp;{reservation.accommodationName}
        </Typography>
        <Typography gutterBottom variant="h6" component="div">
          Starts at: &nbsp;{reservation.startDate}
        </Typography>
        <Typography variant="h6" component="div">
          Ends at: &nbsp;&nbsp;&nbsp;{reservation.endDate}
        </Typography>
        <Typography variant="body2" color="text.secondary">
          {reservation.numberOfGuests} guests
        </Typography>
      </CardContent>
      <CardActions sx={{ justifyContent: 'space-between' }}>
        <Button variant="outlined" color="secondary" size="small" onClick={() => onCancel(reservation.id)}>
          Cancel
        </Button>
        <Button variant="outlined" size="small" onClick={handleOpenModal}>
          Flights recommendation
        </Button>
      </CardActions>
      <Modal
        open={openModal}
        onClose={handleCloseModal}
      >
        <Box
          sx={{
            position: 'absolute',
            top: '50%',
            left: '50%',
            transform: 'translate(-50%, -50%)',
            backgroundColor: 'white',
            boxShadow: 24,
            padding: '2rem',
            width: '800px',
            borderRadius: '8px'
          }}
        >
          <FlightsRecommendation startDate={reservation.startDate} endDate={reservation.endDate} location={"Novi Sad"} />
        </Box>
      </Modal>
    </Card>
  );
};

export default ReservationCard;
