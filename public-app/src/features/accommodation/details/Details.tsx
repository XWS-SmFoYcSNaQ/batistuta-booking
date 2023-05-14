import FullCalendar from "@fullcalendar/react";
import { Box, Button, Container } from "@mui/material";
import dayGridPlugin from "@fullcalendar/daygrid";
import interactionPlugin from "@fullcalendar/interaction";
import ArrowBackIcon from "@mui/icons-material/ArrowBack";
import { Link, useParams } from "react-router-dom";
import { Period } from "../../../shared/model";
import { useEffect, useState } from "react";
import { AppState, appStore } from "../../../core/store";
import { toast } from "react-toastify";

const getInitialData = (): Period => {
  return {
    start: undefined,
    end: undefined,
  };
};

interface Event {
  title: string;
  start?: Date;
  end?: Date;
  color?: string;
}

export const Details = () => {
  const params = useParams();
  const [data, setData] = useState<Period>(getInitialData());
  const createPeriod = appStore((state: AppState) => state.period.createPeriod);
  const fetchPeriods = appStore((state: AppState) => state.period.fetchPeriods);
  const periods = appStore((state: AppState) => state.period.data);

  const [events, setEvents] = useState<Event[]>([]);

  const renderEventContent = (eventInfo: any) => {
    return (
      <>
        <i>{eventInfo.event.title}</i>
      </>
    );
  };

  const handleSubmit = async (e: any) => {
    e.preventDefault();
    try {
      if (!data.start || !data.end) return;
      await createPeriod({ ...data, accommodationId: params.id, userId: "" });
      toast.success("Period created successfully");
      setData(getInitialData());
      fetchPeriods(params.id ?? "");
    } catch (e: any) {
      toast.error(e.message);
    }
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
    fetchPeriods(params.id ?? "");
  }, [fetchPeriods, params.id]);

  useEffect(() => {
    setEvents(
      periods.map((p: Period) => ({
        title: !p.userId || p.userId === "" ? "Unavailable" : "",
        start: new Date(p.start ?? "") ?? undefined,
        end: new Date(p.end ?? "") ?? undefined,
        color: !p.userId || p.userId === "" ? "#800" : "green",
      }))
    );
  }, [periods]);
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
      <FullCalendar
        plugins={[dayGridPlugin, interactionPlugin]}
        initialView="dayGridMonth"
        weekends={true}
        events={events}
        eventContent={renderEventContent}
        selectable={true}
        select={selectHandler}
      />
      <Box sx={{marginTop: "30px"}}>
        <form onSubmit={handleSubmit}>
          <Button type="submit" variant="outlined">
            Set Unavailable
          </Button>
        </form>
      </Box>
    </Container>
  );
};
