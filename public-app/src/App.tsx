import AppBar from "@mui/material/AppBar";
import Box from "@mui/material/Box";
import CssBaseline from "@mui/material/CssBaseline";
import Divider from "@mui/material/Divider";
import Drawer from "@mui/material/Drawer";
import List from "@mui/material/List";
import ListItem from "@mui/material/ListItem";
import ListItemButton from "@mui/material/ListItemButton";
import Toolbar from "@mui/material/Toolbar";
import Typography from "@mui/material/Typography";
import { Outlet, useNavigate } from "react-router";
import { ListItemIcon, Button } from "@mui/material";
import { NavLink } from "react-router-dom";
import HomeIcon from "@mui/icons-material/Home";
import HouseIcon from "@mui/icons-material/House";
import BedIcon from "@mui/icons-material/Bed";
import LoginIcon from "@mui/icons-material/Login";
import PersonIcon from "@mui/icons-material/Person";
import AccountBoxIcon from "@mui/icons-material/AccountBox";
import BookOnlineIcon from "@mui/icons-material/BookOnline";
import CheckIcon from "@mui/icons-material/Check";
import { ToastContainer } from "react-toastify";
import { AppState, appStore } from "./core/store";
import "rsuite/dist/rsuite.min.css";
import { useEffect, useState } from "react";
import AssignmentIndIcon from "@mui/icons-material/AssignmentInd";

const drawerWidth = 300;

interface NavItem {
  route: string;
  text: string;
  icon: JSX.Element;
  roles?: number[];
}

const upperNavItems: NavItem[] = [
  {
    route: "/",
    text: "Home",
    icon: <HomeIcon />,
  },
  {
    route: "/accommodation/my",
    text: "My accommodations",
    icon: <HouseIcon />,
    roles: [1],
  },
  {
    route: "/accommodation/all",
    text: "Accommodations",
    icon: <HouseIcon />,
  },
  {
    route: "/user/hosts",
    text: "Hosts",
    icon: <AssignmentIndIcon />,
  },
  {
    route: "/rooms",
    text: "Rooms",
    icon: <BedIcon />,
  },
  {
    route: "/reservations",
    text: "My reservations",
    icon: <BookOnlineIcon />,
  },
  {
    route: "reservations-to-confirm",
    text: "Reservations confirmation",
    icon: <CheckIcon />,
  },
  {
    route: "/profile",
    text: "Profile",
    icon: <AccountBoxIcon />,
  },
];

const lowerNavItems: NavItem[] = [
  {
    route: "/login",
    text: "Login",
    icon: <LoginIcon />,
  },
  {
    route: "/register",
    text: "Register",
    icon: <PersonIcon />,
  },
];

export default function App() {
  const isAuthenticated = appStore(
    (state: AppState) => state.auth.user != null
  );
  const logoutUser = appStore((state: AppState) => state.auth.logout);
  const currentUser = appStore((state: AppState) => state.auth.user);
  const navigate = useNavigate();

  function logout() {
    window.localStorage.removeItem("jwt");
    logoutUser();
    navigate("/login");
  }

  const filteredLowerNavItems = isAuthenticated
    ? lowerNavItems.filter(
        (item) => item.route !== "/login" && item.route !== "/register"
      )
    : lowerNavItems;

  const [upperNav, setUpperNav] = useState<NavItem[]>([] as NavItem[]);

  useEffect(() => {
    let filteredUpperNavItems = !isAuthenticated
      ? upperNavItems.filter((item) => item.route !== "/profile")
      : upperNavItems;
    filteredUpperNavItems = filteredUpperNavItems.filter(
      (item) => !item.roles || item.roles?.includes(currentUser?.Role ?? -1)
    );
    setUpperNav(filteredUpperNavItems);
  }, [currentUser, isAuthenticated]);

  return (
    <Box sx={{ display: "flex" }}>
      <CssBaseline />
      <AppBar
        position="fixed"
        sx={{ width: `calc(100% - ${drawerWidth}px)`, ml: `${drawerWidth}px` }}
      >
        <Toolbar>
          <Typography variant="h5" noWrap component="div">
            Welcome {currentUser && currentUser?.FirstName} to Batistuta Booking
          </Typography>
        </Toolbar>
      </AppBar>
      <Drawer
        sx={{
          width: drawerWidth,
          flexShrink: 0,
          "& .MuiDrawer-paper": {
            width: drawerWidth,
            boxSizing: "border-box",
          },
        }}
        variant="permanent"
        anchor="left"
      >
        <Toolbar />
        <Divider />
        {/* Upper nav items */}
        <List>
          {upperNav.map((navItem, index) => (
            <NavLink to={navItem.route} key={navItem.route}>
              <ListItem disablePadding>
                <ListItemButton>
                  <ListItemIcon>{navItem.icon}</ListItemIcon>
                  {navItem.text}
                </ListItemButton>
              </ListItem>
            </NavLink>
          ))}
        </List>
        <Divider />
        {/* Lower nav items */}
        <List>
          {filteredLowerNavItems.map((navItem, index) => (
            <NavLink to={navItem.route} key={navItem.route}>
              <ListItem disablePadding>
                <ListItemButton>
                  <ListItemIcon>{navItem.icon}</ListItemIcon>
                  {navItem.text}
                </ListItemButton>
              </ListItem>
            </NavLink>
          ))}
        </List>
        {isAuthenticated && <Button onClick={logout}>Logout</Button>}
      </Drawer>
      <Box
        component="main"
        sx={{
          flexGrow: 1,
          bgcolor: "background.default",
          py: 8,
          overflowX: "visible",
        }}
      >
        <Outlet />
      </Box>
      <ToastContainer />
    </Box>
  );
}
