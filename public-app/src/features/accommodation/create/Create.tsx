import { useState } from "react";
import { Accommodation } from "../../../shared/model";
import { Box, Button, TextField } from "@mui/material";

const getInitialData = (): Accommodation => {
  return {
    name: "",
    benefits: "",
    minGuests: 0,
    maxGuests: 0,
    basePrice: 0
  };
};

export const Create = () => {
  const [data, setData] = useState<Accommodation>(getInitialData());

  const handleSubmit = (e: any) => {
    e.preventDefault();
    console.log(data)
    setData(getInitialData());
  };

  return (
    <div>
      <form onSubmit={handleSubmit}>
      <h3>Add new accommodation</h3>
      <div className="grid">
        <TextField
          required
          label="Name"
          value={data.name ?? ""}
          onChange={(e) =>
            setData({ ...data, name: e.target.value })
          }
        />
        <TextField
          required
          label="Benefits"
          value={data.benefits ?? ""}
          onChange={(e) => setData({ ...data, benefits: e.target.value })}
        />
        <TextField
          required
          label="Price"
          type="number"
          value={data.basePrice ?? ""}
          onChange={(e) =>
            setData({
              ...data,
              basePrice:
                e.target.value === "" ? 0 : parseFloat(e.target.value),
            })
          }
        />
        <TextField
          required
          label="Min Guests"
          type="number"
          value={data.minGuests ?? ""}
          onChange={(e) =>
            setData({
              ...data,
              minGuests:
                e.target.value === "" ? 0 : parseFloat(e.target.value),
            })
          }
        />
        <TextField
          required
          label="Max Guests"
          type="number"
          value={data.maxGuests ?? ""}
          onChange={(e) =>
            setData({
              ...data,
              maxGuests:
                e.target.value === "" ? 0 : parseFloat(e.target.value),
            })
          }
        />
      </div>
      <Box marginTop="20px" sx={{ display: "flex", justifyContent: "right" }}>
        <Button size="large" variant="outlined" type="submit">
          Create Accommodation
        </Button>
      </Box>
    </form>
    </div>
  )
}