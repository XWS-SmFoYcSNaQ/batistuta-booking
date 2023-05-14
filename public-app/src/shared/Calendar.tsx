import FullCalendar from "@fullcalendar/react";
import dayGridPlugin from "@fullcalendar/daygrid";
import interactionPlugin from "@fullcalendar/interaction";

export interface CalendarEvent {
  title: string;
  start?: Date;
  end?: Date;
  color?: string;
}

export const Calendar = ({
  events,
  selectHandler,
  renderEventContent
}: {
  events: CalendarEvent[];
  selectHandler: (e: any) => void;
  renderEventContent: (e: any) => React.ReactNode;
}) => {
  

  return (
    <FullCalendar
      plugins={[dayGridPlugin, interactionPlugin]}
      initialView="dayGridMonth"
      weekends={true}
      events={events}
      eventContent={renderEventContent}
      selectable={true}
      select={selectHandler}
    />
  );
};
