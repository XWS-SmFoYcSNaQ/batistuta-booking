import { useState } from "react";
import { Accommodation } from "../../../shared/model";
import { Box, Button, TextField } from "@mui/material";
import { AppState, appStore } from "../../../core/store";
import { useNavigate } from "react-router";
import { Link } from "react-router-dom";
import { toast } from "react-toastify";

const getInitialData = (): Accommodation => {
  return {
    name: "",
    benefits: "",
    minGuests: 0,
    maxGuests: 0,
    basePrice: 0,
  };
};

export const Create = () => {
  const navigate = useNavigate();
  const [data, setData] = useState<Accommodation>(getInitialData());
  const createAccommodation = appStore(
    (state: AppState) => state.accommodation.createAccommodation
  );
  const fetchMyAccommodations = appStore(
    (state: AppState) => state.accommodation.fetchMyAccommodations
  )

  const handleSubmit = async (e: any) => {
    e.preventDefault();
    try {
      await createAccommodation(data);
      navigate("/accommodation/my");
      toast.success("Accommodation created successfully")
      setData(getInitialData());
      fetchMyAccommodations()
    } catch (e: any) {
      toast.error(e.message)
    }
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
            onChange={(e) => setData({ ...data, name: e.target.value })}
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
                  e.target.value === "" ? undefined : parseFloat(e.target.value),
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
                  e.target.value === "" ? undefined : parseFloat(e.target.value),
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
                  e.target.value === "" ? undefined : parseFloat(e.target.value),
              })
            }
          />
        </div>
        <Box marginTop="20px" sx={{ display: "flex", justifyContent: "right", gap: "15px" }}>
          <Link to="/accommodation/my">
            <Button size="large" color="error" type="button">
              Cancel
            </Button>
          </Link>
          <Button size="large" variant="outlined" type="submit">
            Create Accommodation
          </Button>
        </Box>
      </form>
    </div>
  );
};
