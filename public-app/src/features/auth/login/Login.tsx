import { AccountCircle } from "@mui/icons-material";
import KeyIcon from '@mui/icons-material/Key';
import { Button, Card, CardActions, CardContent, CircularProgress, Container, Grid, InputAdornment, TextField } from "@mui/material";
import { useState } from "react";
import { useNavigate } from "react-router";
import { AppState, appStore } from "../../../core/store";

const Login = () => {
  const [username, setUsername] = useState('');
  const [password, setPassword] = useState('');
  const loginFunc = appStore((state: AppState) => state.auth.login);
  const navigate = useNavigate();

  async function login() {
    if (!username || !password) return;
    const success = await loginFunc(username, password);
    if (success) {
      navigate("/");
    }
  }

  return (
    <Container sx={{ marginTop: 8}}>
      <Card sx={{ minWidth: 275, maxWidth: 500, mx: "auto", px: '12px', py: '16px' }}>
        <CardContent>
          <Grid container spacing={2} alignItems="center">
            <Grid item xs={8}>
              <TextField 
                id="username" 
                required 
                label="Username" 
                variant="standard"
                value={username}
                onChange={e => { setUsername(e.target.value); }}
                onKeyDown={e => { if(e.key === "Enter") login(); }}
                InputProps={{
                  startAdornment: (
                    <InputAdornment position="start">
                      <AccountCircle />
                    </InputAdornment>
                  ),
                }}/>
            </Grid>
            <Grid item xs={8}>
              <TextField 
                id="password" 
                required 
                type="password"
                label="Password" 
                variant="standard" 
                value={password}
                onChange={e => { setPassword(e.target.value); }}
                onKeyDown={e => { if(e.key === "Enter") login(); }}
                InputProps={{
                  startAdornment: (
                    <InputAdornment position="start">
                      <KeyIcon />
                    </InputAdornment>
                  ),
                }}/>
            </Grid>
          </Grid>
        </CardContent>
        <CardActions>
          <Button onClick={login} size="medium" variant="outlined" color="primary">Login</Button>
        </CardActions>
      </Card>
    </Container>
  )
}

export default Login;