import { Card, CardContent, Typography } from "@mui/material";
import { Notification, NotificationType  } from "../../shared/model/notifications";

interface Props {
  notification: Notification;
}

const SingleNotification = ({ notification } : Props) => {
  return (
    <Card sx={{ minWidth: 275, backgroundColor: 'var(--primary-light)', color: 'var(--primary-text)', px: '10px', py: '6px' }}>
      <CardContent>
        <Typography sx={{ fontSize: 14 }} color="text.secondary" gutterBottom>
          { notification.createdAt && notification.createdAt.toString() }
        </Typography>
        <Typography variant="h5" component="div">
          {notification.title }
        </Typography>
        <Typography sx={{ mb: 1.5 }} fontSize="0.85rem">
          {  notification.type }
        </Typography>
        <Typography variant="body2">
          { notification.content }
        </Typography>
      </CardContent>
    </Card>
  )
};

export default SingleNotification;