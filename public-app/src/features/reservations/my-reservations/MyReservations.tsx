import * as React from 'react';
import Tabs from '@mui/material/Tabs';
import Tab from '@mui/material/Tab';
import Typography from '@mui/material/Typography';
import Box from '@mui/material/Box';
import { ReactNode, SyntheticEvent, useEffect, useState } from 'react';
import { Link, useLocation } from 'react-router-dom';
import { PendingReservations } from './PendingReservations';
import { AcceptedReservations } from './AcceptedReservations';

interface TabPanelProps {
  children?: ReactNode;
  index: number;
  value: number;
}

function TabPanel(props: TabPanelProps) {
  const { children, value, index, ...other } = props;

  return (
    <div
      role="tabpanel"
      hidden={value !== index}
      id={`simple-tabpanel-${index}`}
      aria-labelledby={`simple-tab-${index}`}
      {...other}
    >
      {value === index && (
        <Box sx={{ p: 3 }}>
          <Typography>{children}</Typography>
        </Box>
      )}
    </div>
  );
}

export const MyReservations = () => {
  const location = useLocation();
  const [value, setValue] = useState(0);

  const handleChange = (event: SyntheticEvent, newValue: number) => {
    setValue(newValue);
  };

  useEffect(() => {
    const tabValue = location.pathname.includes('accepted') ? 1 : 0;
    setValue(tabValue);
  }, [location.pathname]);

  return (
    <>
      <Box sx={{ width: '100%' }}>
        <Box sx={{ borderBottom: 1, borderColor: 'divider' }}>
          <Tabs value={value} onChange={handleChange}>
            <Tab
              label="Pending"
              component={Link}
              to="/reservations/pending"
            />
            <Tab
              label="Accepted"
              component={Link}
              to="/reservations/accepted"
            />
          </Tabs>
        </Box>
        <TabPanel value={value} index={0}>
          <PendingReservations />
        </TabPanel>
        <TabPanel value={value} index={1}>
          <AcceptedReservations />
        </TabPanel>
      </Box>
    </>
  );
};