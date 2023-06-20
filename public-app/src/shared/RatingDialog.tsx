import {
  Dialog,
  DialogTitle,
  DialogContent,
  DialogContentText,
  DialogActions,
  Button,
  FormControl,
  InputLabel,
  MenuItem,
  Select,
  SelectChangeEvent,
} from "@mui/material";
import { ReactElement, useEffect, useState } from "react";

export const RatingDialog = ({
  open,
  setOpen,
  title,
  description,
  onClose,
  onRate,
  initialRating,
}: {
  open: boolean;
  setOpen: (open: boolean) => void;
  title: string;
  onRate: (value: number) => Promise<void>;
  description?: ReactElement;
  onClose?: () => void;
  initialRating?: number;
}) => {
  const [rating, setRating] = useState(initialRating ?? 1);

  const handleChange = (event: SelectChangeEvent) => {
    setRating(parseInt(event.target.value));
  };
  useEffect(() => {
    setRating(initialRating ?? 1);
  }, [initialRating]);

  const handleClose = () => {
    setOpen(false);
    if (onClose) onClose();
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
      <DialogTitle id="rating-dialog-title">
        {title}
      </DialogTitle>
      <DialogContent sx={{ minWidth: "300px" }}>
        <DialogContentText id="rating-dialog-description">
          {description}
        </DialogContentText>
        <FormControl fullWidth sx={{ marginTop: "15px" }}>
          <InputLabel id="demo-simple-select-label">Age</InputLabel>
          <Select
            labelId="demo-simple-select-label"
            id="demo-simple-select"
            value={rating.toString()}
            label="Age"
            onChange={handleChange}
          >
            <MenuItem value={1}>1</MenuItem>
            <MenuItem value={2}>2</MenuItem>
            <MenuItem value={3}>3</MenuItem>
            <MenuItem value={4}>4</MenuItem>
            <MenuItem value={5}>5</MenuItem>
          </Select>
        </FormControl>
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
