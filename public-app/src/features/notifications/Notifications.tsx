import { AppState, appStore } from "../../core/store";
import { Container, Grid } from "@mui/material";
import SingleNotification from "./SingleNotification";

const Notifications = () => {
  const notifications = appStore((state: AppState) => state.notification.data);

  return (
    <Container sx={{ py: '18px'}}>
      <h2 style={{ textAlign: 'center' }}>Notifications</h2>
      <Grid sx={{ mt: '16px'}} container justifyContent="left" spacing={6}>
        {notifications.map(x => (
          <Grid item xs={9} md={6}  xl={4} key={x.id}>
            <SingleNotification notification={x}/>
          </Grid>
        ))}
      </Grid>
    </Container>
  )
}

export default Notifications;