import {
  Dialog,
  DialogTitle,
  DialogContent,
  DialogContentText,
  DialogActions,
  Button,
} from "@mui/material";
import { ReactElement, useState } from "react";

export const RatingDialog = ({
  open,
  setOpen,
  title,
  description,
  onClose,
  onRate,
}: {
  open: boolean;
  setOpen: (open: boolean) => void;
  title: string;
  onRate: (value: number) => Promise<void>;
  description?: ReactElement;
  onClose?: () => void;
}) => {
  const [rating, setRating] = useState(0)  

  const handleClose = () => {
    setOpen(false);
    if(onClose)
      onClose()
  };
  const handleRate = async () => {
    try {
      await onRate(rating);
      handleClose();
    } catch (err) {
      console.log(err);
    }
  };
  return (
    <Dialog
      open={open}
      onClose={handleClose}
      aria-labelledby="rating-dialog-title"
      aria-describedby="rating-dialog-description"
    >
      <DialogTitle id="rating-dialog-title">{title}</DialogTitle>
      <DialogContent>
        <DialogContentText id="rating-dialog-description">
          {description}
        </DialogContentText>
      </DialogContent>
      <DialogActions>
        <Button onClick={handleClose}>Cancel</Button>
        <Button onClick={handleRate} autoFocus>
          Rate
        </Button>
      </DialogActions>
    </Dialog>
  );
};
