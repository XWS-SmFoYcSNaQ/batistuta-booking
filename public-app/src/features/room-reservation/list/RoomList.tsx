import Card from '@mui/material/Card';
import CardActions from '@mui/material/CardActions';
import CardContent from '@mui/material/CardContent';
import CardMedia from '@mui/material/CardMedia';
import Button from '@mui/material/Button';
import Typography from '@mui/material/Typography';
import { AppState, appStore } from '../../../core/store';
import { useEffect } from 'react';
import { Box, CircularProgress } from '@mui/material';
import { Link } from 'react-router-dom';

export const RoomList = () => {
	const loading = appStore((state: AppState) => state.accommodation.loading);
	const fetchAccommodations = appStore(
		(state: AppState) => state.accommodation.fetchAccommodations
	);
	const accommodations = appStore(
		(state: AppState) => state.accommodation.data
	);

	useEffect(() => {
		fetchAccommodations();
	}, [fetchAccommodations]);
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
			{!loading && accommodations.length > 0 && (
				<>
					<h2>Rooms</h2>
					<div style={{ display: 'flex', flexWrap: 'wrap' }}>
						{accommodations.map((room: any, index: number) => (
							<Card sx={{ minWidth: 250, maxWidth: 350, margin: '1rem' }}>
								<CardMedia
									sx={{ height: 140 }}
									image="" />
								<CardContent>
									<Typography gutterBottom variant="h4" component="div">
										{room.name}
									</Typography>
									<Typography variant="body2" color="text.secondary">
										{room.basePrice}€ per night
									</Typography>
								</CardContent>
								<CardActions>
									<Link to={`/rooms/${room.id}`}>
										<Button variant="outlined" size="small">Book</Button>
									</Link>
								</CardActions>
							</Card>
						))}
					</div>
				</>
			)}
		</>
	);
};