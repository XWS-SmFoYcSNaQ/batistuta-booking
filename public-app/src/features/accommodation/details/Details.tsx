import FullCalendar from "@fullcalendar/react";
import { Box, Button, Container } from "@mui/material";
import dayGridPlugin from "@fullcalendar/daygrid";
import ArrowBackIcon from "@mui/icons-material/ArrowBack";
import { Link } from "react-router-dom";

const events = [
  { title: "event 1", date: "2023-05-05" },
  { title: "event 2", date: "2023-05-09" },
];

export const Details = () => {
  function renderEventContent(eventInfo: any) {
    return (
      <>
        <b>{eventInfo.timeText}</b>
        <i>{eventInfo.event.title}</i>
      </>
    );
  }

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
        plugins={[dayGridPlugin]}
        initialView="dayGridMonth"
        weekends={false}
        events={events}
        eventContent={renderEventContent}
      />
    </Container>
  );
};
