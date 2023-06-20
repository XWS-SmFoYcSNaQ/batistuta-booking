import { Box, Container } from "@mui/material";
import { Outlet } from "react-router";

export const AccommodationRoot = () => {
  return (
    <Container>
      <Box sx={{ paddingTop: "10px", paddingBottom: "130px" }}>
        <Outlet />
      </Box>
    </Container>
  );
};
