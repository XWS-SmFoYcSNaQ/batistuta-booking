import { Button, Card, CardActions, CardContent, Container, FormControl, Grid, InputLabel, MenuItem, Select, SelectChangeEvent, TextField } from "@mui/material";
import { useState } from "react";
import { RegisterRequest } from "../../../shared/model/authentication";
import { AppState, appStore } from "../../../core/store";
import { useNavigate } from "react-router";
import { toast } from "react-toastify";

interface Error {
  Field?: string;
  Error?: string;
}

const Register = () => {
  const [username, setUsername] = useState('');
  const [password, setPassword] = useState('');
  const [firstName, setFirstName] = useState('');
  const [lastName, setLastName] = useState('');
  const [email, setEmail] = useState('');
  const [livingPlace, setLivingPlace] = useState('');
  const [role, setRole] = useState('');
  const registerFunc = appStore((state: AppState) => state.auth.register)
  const navigate = useNavigate();
  
  async function register() {
    const errors = validateFields();
    if (errors.length !== 0) {
      console.log(errors);
      return;
    }
    const registerRequest : RegisterRequest = {
      Username: username,
      Password: password,
      FirstName: firstName,
      LastName: lastName,
      Email: email,
      LivingPlace: livingPlace,
      Role: role
    } 
    const success = await registerFunc(registerRequest);
    if (success) {
      navigate("/");
      toast.success("Registration successfull.");
    } 
  }

  function validateFields(): Error[] {
    const errors : Error[] = [];
    if (!username) { errors.push({ Field: 'username', Error: 'Username is mandatory' }); }
    if (!password || password.length < 8) errors.push({ Field: 'password', Error: 'Password is mandatory and must be at least 8 characters long.'});
    if (!firstName || firstName.charAt(0).match(new RegExp("[a-z]{1}"))) errors.push({ Field: 'firstName', Error: 'First name is mandatory and must begin with uppercase'});
    if (!lastName || lastName.charAt(0).match(new RegExp("[a-z]{1}"))) errors.push({ Field: 'lastName', Error: 'Last name is mandatory and must begin with uppercase'});
    if (!email || !email.includes("@")) errors.push({ Field: 'email', Error: 'Email is mandatory and must contain @'});
    if (!livingPlace) errors.push({ Field: 'livingPlace', Error: 'Living place is mandatory' });
    if (!role) errors.push({ Field: 'role', Error: 'Role is mandatory'});

    return errors;
  }

  const handleChange = (event: SelectChangeEvent) => {
    setRole(event.target.value as string);
  };

  return (
    <Container sx={{ marginTop: 8}}>
    <Card sx={{ minWidth: 275, maxWidth: 500, mx: "auto", px: '12px', py: '16px' }}>
      <CardContent>
        <Grid container spacing={2} alignItems="center">
          <Grid item xs={6}>
            <TextField 
              id="username" 
              required 
              label="Username" 
              variant="standard"
              value={username}
              onChange={(e) => { setUsername(e.target.value)}}/>
          </Grid>
          <Grid item xs={6}>
            <TextField 
              id="password" 
              required 
              label="Password"
              type="password" 
              variant="standard" 
              value={password}
              onChange={(e) => { setPassword(e.target.value)}}/>
          </Grid>
          <Grid item xs={6}>
            <TextField 
              id="firstName" 
              required 
              label="First Name" 
              variant="standard" 
              value={firstName}
              onChange={(e) => { setFirstName(e.target.value)}}/>
          </Grid>
          <Grid item xs={6}>
            <TextField 
              id="lastName" 
              required 
              label="Last Name" 
              variant="standard"
              value={lastName}
              onChange={(e) => { setLastName(e.target.value)}}/>
          </Grid>
          <Grid item xs={6}>
            <TextField 
              id="email" 
              required 
              label="Email" 
              variant="standard" 
              value={email}
              onChange={(e) => { setEmail(e.target.value)}}/>
          </Grid>
          <Grid item xs={6}>
            <TextField 
              id="living place" 
              required 
              label="Living Place" 
              variant="standard" 
              value={livingPlace}
              onChange={(e) => { setLivingPlace(e.target.value)}}/>
          </Grid>
          <Grid item xs={6}>
            <FormControl fullWidth>
              <InputLabel id="role-label">Role</InputLabel>
              <Select
                labelId="role-label"
                id="role"
                variant="standard"
                value={role}
                label="Role"
                onChange={handleChange}
              >
                <MenuItem value="Host">Host</MenuItem>
                <MenuItem value="Guest">Guest</MenuItem>
              </Select>
            </FormControl>
          </Grid>
        </Grid>
      </CardContent>
      <CardActions>
        <Button onClick={register} size="medium" variant="outlined" color="primary">Register</Button>
      </CardActions>
    </Card>
  </Container>
  )
}

export default Register;