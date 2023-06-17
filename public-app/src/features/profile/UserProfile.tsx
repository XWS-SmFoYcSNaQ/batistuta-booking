import { useState } from "react";
import { AppState, apiUrl, appStore } from "../../core/store";
import { Box, Button, Card, CardActions, CardContent, CircularProgress, Container, Grid, Stack, TextField, Typography } from "@mui/material";
import { toast } from "react-toastify";
import { UserRole } from "../../shared/model/user";
import useFetch from "../../shared/hooks/useFetch";
import CheckBoxIcon from '@mui/icons-material/CheckBox';

const UserProfile = () => {
  const user = appStore((state: AppState) => state.auth.user);
  const updateUserInfo = appStore((state: AppState) => state.auth.updateUserInfo);
  const loading = appStore((state: AppState) => state.auth.loading);
  const [userDetail, setUserDetail] = useState({ firstName: user?.FirstName, lastName: user?.LastName, livingPlace: user?.LivingPlace});
  const [updatePasswordInfo, setUpdatePasswordInfo] = useState({ currentPassword: '', newPassword: ''});
  const changePassword = appStore((state: AppState) => state.auth.changePassword);
  const { loading: loadingIsFeatured, error: featuredError, data: featuredData } = useFetch(`${apiUrl}/api/featured`);

  const updateUser = async () => {
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

  const changeUserPassword = async () => {
    if (!updatePasswordInfo.currentPassword || !updatePasswordInfo.newPassword) {
      toast.error("All fields are mandatory");
      return;
    }
    
    await changePassword({ CurrentPassword: updatePasswordInfo.currentPassword, NewPassword: updatePasswordInfo.newPassword });
  }

  if (loading || loadingIsFeatured) return (
    <Container sx={{ height: '100vh', display: 'flex'}}>
      <CircularProgress color="primary" sx={{ mx: 'auto', mt: '50px' }}/>
    </Container>
  )

  if (user?.Role === UserRole.Host && featuredError) return (
    <Container sx={{ height: '100vh', display: 'flex'}}>
      { featuredError }
    </Container>
  )

  return (
    <Container sx={{ mt: 10 }}>
      { user?.Role === UserRole.Host && featuredData && featuredData.Featured && (
          <Box sx={{ display: 'flex', alignItems: 'center', justifyContent: 'center', mb: '20px', fontSize: '1.5rem' }}>
                {featuredData.Message} <CheckBoxIcon sx={{ ml: '5px'}}/> 
          </Box>
      )}
      <Stack direction="row" spacing={8} >
        <Box sx={{ minHeight: '450px' }} >
          <Card sx={{ maxWidth: '600px', height: '100%', mx: 'auto', py: '18px', px: '10px'}}>
            <CardContent>
              <Grid container spacing={4} alignItems="center" justifyContent="center">
                <Grid item xs={12}>
                  <Typography variant="h6" textAlign="center">Your Information</Typography>
                </Grid>
                <Grid item xs={10} md={8} lg={7}>
                  <TextField 
                    sx={{ width: '100%' }}
                    id="firstName"
                    label="First Name"
                    variant="outlined"
                    value={userDetail.firstName}
                    onChange={(e) => {setUserDetail({...userDetail, firstName: e.target.value})}}
                    onKeyDown={(e) => { if(e.key === 'Enter' ) updateUser(); }} />
                </Grid>
                <Grid item xs={10} md={8} lg={7}>
                  <TextField 
                      sx={{ width: '100%' }}
                      id="lastName"
                      label="Last Name"
                      variant="outlined"
                      value={userDetail.lastName}
                      onChange={(e) => {setUserDetail({...userDetail, lastName: e.target.value})}}
                      onKeyDown={(e) => { if(e.key === 'Enter' ) updateUser(); }} />
                </Grid>
                <Grid item xs={10} md={8} lg={7}>
                  <TextField 
                      sx={{ width: '100%' }}
                      id="livingPlace"
                      label="Living Place"
                      variant="outlined"
                      value={userDetail.livingPlace}
                      onChange={(e) => {setUserDetail({...userDetail, livingPlace: e.target.value})}}
                      onKeyDown={(e) => { if(e.key === 'Enter' ) updateUser(); }} />
                </Grid>
                <Grid item xs={10} md={8} lg={7}>
                  <TextField 
                      sx={{ width: '100%' }}
                      id="Role"
                      label="Role"
                      variant="outlined"
                      value={UserRole[user?.Role ? user.Role : 0]}
                      disabled/>
                </Grid>
              </Grid>
            </CardContent>
            <CardActions sx={{ justifyContent: 'center' }}>
              <Button onClick={updateUser} sx={{ ml: '8px'}} variant="contained" size="medium" color="primary">Save</Button>
            </CardActions>
          </Card>
        </Box >

        <Box sx={{ minHeight: '450px' }}>
          <Card sx={{ maxWidth: '600px', height: '100%', mx: 'auto', py: '18px', px: '10px'}}>
            <CardContent>
              <Grid container spacing={4} alignItems="center" justifyContent="center">
                <Grid item xs={12}>
                  <Typography variant="h6" textAlign="center">Change Password</Typography>
                </Grid>
                <Grid item xs={10} md={8} lg={7}>
                  <TextField 
                    sx={{ width: '100%' }}
                    id="currentPassword"
                    label="Current Password"
                    variant="outlined"
                    type="password"
                    value={updatePasswordInfo.currentPassword}
                    onChange={(e) => {setUpdatePasswordInfo({...updatePasswordInfo, currentPassword: e.target.value})}}
                    onKeyDown={(e) => { if(e.key === 'Enter' ) changeUserPassword(); }} />
                </Grid>
                <Grid item xs={10} md={8} lg={7}>
                  <TextField 
                      sx={{ width: '100%' }}
                      id="newPassword"
                      label="New Password"
                      variant="outlined"
                      type="password"
                      value={updatePasswordInfo.newPassword}
                      onChange={(e) => {setUpdatePasswordInfo({...updatePasswordInfo, newPassword: e.target.value})}}
                      onKeyDown={(e) => { if(e.key === 'Enter' ) changeUserPassword(); }} />
                </Grid>
              </Grid>
            </CardContent>
            <CardActions sx={{ justifyContent: 'center' }}>
              <Button onClick={changeUserPassword} sx={{ ml: '8px'}} variant="contained" size="medium" color="primary">Save</Button>
            </CardActions>
          </Card>
        </Box> 
      </Stack>
    </Container>
  );
}

export default UserProfile;