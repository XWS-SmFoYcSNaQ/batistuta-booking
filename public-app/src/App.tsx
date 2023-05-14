import AppBar from '@mui/material/AppBar';
import Box from '@mui/material/Box';
import CssBaseline from '@mui/material/CssBaseline';
import Divider from '@mui/material/Divider';
import Drawer from '@mui/material/Drawer';
import List from '@mui/material/List';
import ListItem from '@mui/material/ListItem';
import ListItemButton from '@mui/material/ListItemButton';
import Toolbar from '@mui/material/Toolbar';
import Typography from '@mui/material/Typography';
import { useState } from "react";
import "./App.css";
import { Outlet } from 'react-router';
import { ListItemIcon, Button } from '@mui/material';
import { NavLink } from 'react-router-dom';
import HomeIcon from '@mui/icons-material/Home';
import HouseIcon from '@mui/icons-material/House';
import BedIcon from '@mui/icons-material/Bed';
import { ToastContainer } from 'react-toastify';

const drawerWidth = 240;

interface NavItem {
  route: string;
  text: string;
  icon: JSX.Element;
}

export default function App() {
  const [isAuthenticated, setIsAuthenticated] = useState(true);

  const upperNavItems: NavItem[] = [
    {
      route: '/',
      text: 'Home',
      icon: <HomeIcon/>
    },
    {
      route: '/accommodation',
      text: 'Accommodations',
      icon: <HouseIcon/>
    },
    {
      route: '/rooms',
      text: 'Rooms',
      icon: <BedIcon/>
    }
  ];

  const lowerNavItems: NavItem[] = [];

  const filteredLowerNavItems = isAuthenticated
    ? lowerNavItems.filter(item => item.route !== '/login' && item.route !== '/register')
    : lowerNavItems;

  return (
    <Box sx={{ display: 'flex' }}>
      <CssBaseline />
      <AppBar
        position="fixed"
        sx={{ width: `calc(100% - ${drawerWidth}px)`, ml: `${drawerWidth}px` }}
      >
        <Toolbar>
          <Typography variant="h5" noWrap component="div">
            Welcome to Batistuta Booking
          </Typography>
        </Toolbar>
      </AppBar>
      <Drawer
        sx={{
          width: drawerWidth,
          flexShrink: 0,
          '& .MuiDrawer-paper': {
            width: drawerWidth,
            boxSizing: 'border-box',
          },
        }}
        variant="permanent"
        anchor="left"
      >
        <Toolbar />
        <Divider />
        {/* Upper nav items */}
        <List>
          {upperNavItems.map((navItem, index) => (
            <NavLink to={navItem.route} key={navItem.route}>
              <ListItem disablePadding>
                <ListItemButton>
                  <ListItemIcon>
                    {navItem.icon}
                  </ListItemIcon>
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
            <NavLink to={navItem.route} key={navItem.route} >
              <ListItem disablePadding>
                <ListItemButton>
                  <ListItemIcon>
                    {navItem.icon}
                  </ListItemIcon>
                    {navItem.text}
                </ListItemButton>
              </ListItem>
            </NavLink>
          ))}
        </List>
        {isAuthenticated && <Button>Logout</Button>}
      </Drawer>
      <Box
        component="main"
        sx={{ flexGrow: 1, bgcolor: 'background.default', py: 8, overflowX: 'visible' }}
      >
        <Outlet />
      </Box>
      <ToastContainer />
    </Box>
  );
}