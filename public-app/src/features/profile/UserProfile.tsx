import { useState } from "react";
import { AppState, appStore } from "../../core/store";
import { Button, Card, CardActions, CardContent, CircularProgress, Container, Grid, TextField, Typography } from "@mui/material";
import { toast } from "react-toastify";

const UserProfile = () => {
  const user = appStore((state: AppState) => state.auth.user);
  const updateUserInfo = appStore((state: AppState) => state.auth.updateUserInfo);
  const loading = appStore((state: AppState) => state.auth.loading);
  const [userDetail, setUserDetail] = useState({ firstName: user?.FirstName, lastName: user?.LastName, livingPlace: user?.LivingPlace});

  const onClick = async () => {
    if (user?.FirstName === userDetail.firstName && user?.LastName === userDetail.lastName && user?.LivingPlace === userDetail.livingPlace) {
      return;
    }
    if (!userDetail.firstName || !userDetail.lastName || !userDetail.livingPlace) {
      toast.error("All fields are mandatory");
      setUserDetail({ firstName: user?.FirstName, lastName: user?.LastName, livingPlace: user?.LivingPlace});
      return;
    }
    const success = await updateUserInfo({ FirstName: userDetail.firstName, LastName: userDetail.lastName, LivingPlace: userDetail.livingPlace })
    if (!success) {
      setUserDetail({ firstName: user?.FirstName, lastName: user?.LastName, livingPlace: user?.LivingPlace});
    }
  };

  if (loading) return (
    <Container sx={{ height: '100vh', display: 'flex'}}>
      <CircularProgress color="primary" sx={{ mx: 'auto', mt: '50px' }}/>
    </Container>
  )

  return (
    <Container sx={{ mt: 8 }}>
      <Card sx={{ maxWidth: '600px', mx: 'auto', py: '18px', px: '10px'}}>
        <CardContent>
          <Grid container spacing={4} alignItems="center" justifyContent="center">
            <Grid item xs={12}>
              <Typography variant="h6" textAlign="center">User Information</Typography>
            </Grid>
            <Grid item xs={10} md={8} lg={7}>
              <TextField 
                sx={{ width: '100%' }}
                id="firstName"
                label="First Name"
                variant="outlined"
                value={userDetail.firstName}
                onChange={(e) => {setUserDetail({...userDetail, firstName: e.target.value})}}
                onKeyDown={(e) => { if(e.key === 'Enter' ) onClick(); }} />
            </Grid>
            <Grid item xs={10} md={8} lg={7}>
              <TextField 
                  sx={{ width: '100%' }}
                  id="lastName"
                  label="Last Name"
                  variant="outlined"
                  value={userDetail.lastName}
                  onChange={(e) => {setUserDetail({...userDetail, lastName: e.target.value})}}
                  onKeyDown={(e) => { if(e.key === 'Enter' ) onClick(); }} />
            </Grid>
            <Grid item xs={10} md={8} lg={7}>
              <TextField 
                  sx={{ width: '100%' }}
                  id="livingPlace"
                  label="Living Place"
                  variant="outlined"
                  value={userDetail.livingPlace}
                  onChange={(e) => {setUserDetail({...userDetail, livingPlace: e.target.value})}}
                  onKeyDown={(e) => { if(e.key === 'Enter' ) onClick(); }} />
            </Grid>
          </Grid>
        </CardContent>
        <CardActions sx={{ justifyContent: 'center' }}>
          <Button onClick={onClick} sx={{ ml: '8px'}} variant="contained" size="medium" color="primary">Save</Button>
        </CardActions>
      </Card>
    </Container>
  );
}

export default UserProfile;