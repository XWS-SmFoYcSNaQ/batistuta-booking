import { useEffect, useState } from "react";
import { useParams } from "react-router";
import { appStore, AppState } from "../../../core/store";
import { Discount } from "../../../shared/model";
import { toast } from "react-toastify";
import { Calendar, CalendarEvent } from "../../../shared";
import { Container, Box, Button, TextField } from "@mui/material";

const getInitialData = (): Discount => {
  return {
    start: undefined,
    end: undefined,
    discount: 0,
  };
};

export const Discounts = () => {
  const params = useParams();
  const [data, setData] = useState<Discount>(getInitialData());
  const createDiscount = appStore(
    (state: AppState) => state.discount.createDiscount
  );
  const fetchAccommodationDetails = appStore(
    (state: AppState) => state.accommodation.fetchDetails
  );
  const accommodation = appStore(
    (state: AppState) => state.accommodation.accommodation
  );
  const currentUser = appStore((state: AppState) => state.auth.user);

  const [events, setEvents] = useState<CalendarEvent[]>([]);

  const handleSubmit = async (e: any) => {
    e.preventDefault();
    try {
      if (!data.start || !data.end){
        toast.warn("Please select starting and ending dates")
        return
      }
      await createDiscount({ ...data, accommodationId: params.id, userId: "" });
      toast.success("Discount created successfully");
      setData(getInitialData());
      fetchAccommodationDetails(params.id ?? "");
    } catch (e: any) {
      toast.error(e.message);
    }
  };

  const renderEventContent = (eventInfo: any) => {
    return <Box sx={{ padding: "0px 5px" }}>{eventInfo.event.title}</Box>;
  };

  const selectHandler = (e: any) => {
    const start: Date = e.start;
    const end: Date = new Date(e.end.getTime() - 1);
    setData({
      ...data,
      start: start.toISOString(),
      end: end.toISOString(),
    });
  };

  useEffect(() => {
    fetchAccommodationDetails(params.id ?? "");
  }, [fetchAccommodationDetails, params.id]);

  useEffect(() => {
    if (accommodation?.discounts) {
      setEvents(
        accommodation.discounts.map((d: Discount) => ({
          title: !d.userId || d.userId === "" ? `discount(${d.discount}%)` : "",
          start: new Date(d.start ?? "") ?? undefined,
          end: new Date(d.end ?? "") ?? undefined,
          color: "green",
        }))
      );
    }
  }, [accommodation]);

  return (
    <Container>
      <h1>Discounts</h1>
      <Calendar
        events={events}
        selectHandler={selectHandler}
        renderEventContent={renderEventContent}
      />
      <Box sx={{ marginTop: "30px" }}>
        <Box sx={{ marginBottom: "25px" }}>
          Base Price: {accommodation?.basePrice}&nbsp;EUR
        </Box>
        {currentUser?.Role === 1 && (
          <form onSubmit={handleSubmit}>
            <div>
              <TextField
                required
                label="Discount"
                type="number"
                value={data.discount ?? ""}
                onChange={(e) =>
                  setData({
                    ...data,
                    discount:
                      e.target.value === ""
                        ? undefined
                        : parseFloat(e.target.value),
                  })
                }
              />
            </div>
            <Button type="submit" variant="outlined" sx={{ marginTop: "15px" }}>
              Create Discount
            </Button>
          </form>
        )}
      </Box>
    </Container>
  );
};
