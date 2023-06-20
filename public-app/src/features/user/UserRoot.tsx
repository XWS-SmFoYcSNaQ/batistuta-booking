import { Container, Box } from "@mui/material";
import { Outlet } from "react-router";

export const UserRoot = () => {
  return (
    <Container>
      <Box sx={{ paddingTop: "10px", paddingBottom: "130px" }}>
        <Outlet/>
      </Box>
    </Container>
  );
};
