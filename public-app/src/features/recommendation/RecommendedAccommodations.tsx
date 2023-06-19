import { useState, useEffect } from "react";
import { Accommodation } from "../../shared/model";
import axios, { AxiosRequestConfig } from "axios";
import { apiUrl } from "../../core/store";
import { VerifyResponse } from "../../shared/model/authentication";
import { Box, CircularProgress } from '@mui/material';
import Card from '@mui/material/Card';
import CardActions from '@mui/material/CardActions';
import CardContent from '@mui/material/CardContent';
import CardMedia from '@mui/material/CardMedia';
import Typography from '@mui/material/Typography';

export const RecommendedAccommodations = () => {
    const [loading, setLoading] = useState<number>(1) // to get
    const [accommodations, setAccommodations] = useState<Accommodation[]>([]);

    const fetchAccommodations = async () => {
        const verifyConfig : AxiosRequestConfig = {
            headers: {
              'Content-Type': 'application/json',
              'Authorization': 'Bearer '+window.localStorage.getItem("jwt")
            }
          }
        const resp = await axios.post<VerifyResponse>(`${apiUrl}/api/auth/verify`,{},  verifyConfig);
        let userID = resp.data.UserId
        const requestOptions = {
          method: 'GET',
          headers: {
            'Content-Type': 'application/json',
            'Authorization': 'Bearer '+window.localStorage.getItem("jwt")
          }
        };
        const response = await fetch(`${apiUrl}` + '/accommodation/recommendation/' + userID  , requestOptions);
        const data = await response.json();
        setAccommodations(data.data);
        setLoading(0)
    }
    
    useEffect(() => {
        fetchAccommodations();
    }, []);
    
    return (
        <>
            {loading && (
                <Box
                    sx={{
                        display: "flex",
                        justifyContent: "center",
                        padding: "100px 0px",
                    }}
                >
                    <CircularProgress />
                </Box>
            )}
            {!loading && (
                <>
                    <h2>Recommended accommodations</h2>
                    <div style={{ display: 'flex', flexWrap: 'wrap' }}>
                        {accommodations.map((accommodation: any, index: number) => (
                            <Card sx={{ minWidth: 250, maxWidth: 350, margin: '1rem' }}>
                                <CardMedia
                                    sx={{ height: 140 }}
                                    image="" />
                                <CardContent>
                                    <Typography gutterBottom variant="h4" component="div">
                                        Name : {accommodation.name}
                                    </Typography>
                                    <Typography variant="body2" color="text.secondary">
                                        Benefits : {accommodation.benefits}
                                    </Typography>
                                    <Typography variant="body2" color="text.secondary">
                                        Minimum guests : {accommodation.minGuests}
                                    </Typography>
                                    <Typography variant="body2" color="text.secondary">
                                        Maximum guests : {accommodation.maxGuests}
                                    </Typography>
                                    <Typography variant="body2" color="text.secondary">
                                        Base price : {accommodation.basePrice}
                                    </Typography>
                                </CardContent>
                            </Card>
                        ))}
                    </div>
                </>
            )}
        </>
    );
}