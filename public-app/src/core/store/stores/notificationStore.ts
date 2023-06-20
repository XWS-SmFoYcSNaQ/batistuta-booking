import { HttpTransportType, HubConnection, HubConnectionBuilder, HubConnectionState, LogLevel } from "@microsoft/signalr";
import { AppState, GetAppState, SetAppState } from "..";
import { Notification } from "../../../shared/model/notifications";
import { toast } from "react-toastify";
import styles from "../../../shared/styles/notificationStyles.module.css";

export interface NotificationStoreType {
  data: Notification[],
  connection?: HubConnection | null,
  connect: () => void
}

export const notificationStore = (
  set: SetAppState,
  get: GetAppState
) : NotificationStoreType => ({
  data: [],
  connect: async () => {
    const jwt = window.localStorage.getItem("jwt");
    const hubEndpoint = `${process.env.REACT_APP_NOTIFICATIONS_API_URL}/hubs/notification`
    const conn = new HubConnectionBuilder()
      .withUrl(hubEndpoint, {
        skipNegotiation: true,
        transport: HttpTransportType.WebSockets,
        accessTokenFactory: () => jwt ? jwt : ''
      })
      .withAutomaticReconnect()
      .configureLogging(LogLevel.Information)
      .build();

    conn.serverTimeoutInMilliseconds = 15 * 60 * 1000;
    conn.keepAliveIntervalInMilliseconds = 2 * 60 * 60 * 1000;

    conn.on("Notification", (notification: Notification) => {
      if (notification.new){
        toast.info(notification.content, { position: "top-right", className: styles.toastMessage});
      }
      set((draft: AppState) => {
        draft.notification.data = [...draft.notification.data, notification];
        return draft;
      });
    })
    
    conn.on("Disconnected", () => {
      const jwt = window.localStorage.getItem('jwt');
      if (jwt) {
        const appState = get();
        setTimeout(() => {
          appState.notification.connect();
        }, 2000);
      }
    })

    conn.onclose(async () => {
      set((draft: AppState) => {
        draft.notification.connection = null;
        draft.notification.data = [];
        return draft;
      })
    })

    try {
      await conn.start();
      console.assert(conn.state === HubConnectionState.Connected);
      console.log("SignalR connected.");
    } catch (err) {
      console.assert(conn.state === HubConnectionState.Disconnected);
      console.log(err);
    }

    set((draft: AppState) => {
      draft.notification.connection = conn;
      return draft;
    })
  
  }
});