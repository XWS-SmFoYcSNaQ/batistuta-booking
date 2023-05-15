import { Box, Button, Container } from "@mui/material";
import ArrowBackIcon from "@mui/icons-material/ArrowBack";
import { Link, useParams } from "react-router-dom";
import { Period } from "../../../shared/model";
import { useEffect, useState } from "react";
import { AppState, appStore } from "../../../core/store";
import { toast } from "react-toastify";
import { Calendar, CalendarEvent } from "../../../shared";

const getInitialData = (): Period => {
  return {
    start: undefined,
    end: undefined,
  };
};

export const Availability = () => {
  const params = useParams();
  const [data, setData] = useState<Period>(getInitialData());
  const createPeriod = appStore((state: AppState) => state.period.createPeriod);
  const fetchAccommodationDetails = appStore((state: AppState) => state.accommodation.fetchDetails)
  const accommodation = appStore((state: AppState) => state.accommodation.accommodation)

  const [events, setEvents] = useState<CalendarEvent[]>([]);

  const handleSubmit = async (e: any) => {
    e.preventDefault();
    try {
      if (!data.start || !data.end) return;
      await createPeriod({ ...data, accommodationId: params.id, userId: "" });
      toast.success("Period created successfully");
      setData(getInitialData());
      fetchAccommodationDetails(params.id ?? "");
    } catch (e: any) {
      toast.error(e.message);
    }
  };

  const renderEventContent = (eventInfo: any) => {
    return (
      <>
        <i>{eventInfo.event.title}</i>
      </>
    );
  };

  const selectHandler = (e: any) => {
    const start: Date = e.start;
    const end: Date = new Date(e.end.getTime() - 1);
    setData({
      start: start.toISOString(),
      end: end.toISOString(),
    });
  };

  useEffect(() => {
    fetchAccommodationDetails(params.id ?? "")
  }, [fetchAccommodationDetails, params.id]);

  useEffect(() => {
    if(accommodation?.periods){
      setEvents(
        accommodation.periods.map((p: Period) => ({
          title: !p.userId || p.userId === "" ? "Unavailable" : "",
          start: new Date(p.start ?? "") ?? undefined,
          end: new Date(p.end ?? "") ?? undefined,
          color: !p.userId || p.userId === "" ? "#800" : "green",
        }))
      );
    }
  }, [accommodation]);
  return (
    <Container>
      <Box sx={{ margin: "10px 0px" }}>
        <Link to="/accommodation">
          <Button type="submit">
            <ArrowBackIcon sx={{ marginRight: "10px" }} />
            <span>Go back</span>
          </Button>
        </Link>
      </Box>
      <h1>Availability</h1>
      <Calendar events={events} selectHandler={selectHandler} renderEventContent={renderEventContent}/>
      <Box sx={{marginTop: "30px"}}>
        <Box sx={{ marginBottom: "25px" }}>
          Maxiumum Guests: {accommodation?.maxGuests}
        </Box>
        <form onSubmit={handleSubmit}>
          <Button type="submit" variant="outlined">
            Set Unavailable
          </Button>
        </form>
      </Box>
    </Container>
  );
};
