import Link from 'next/link';
import { Container, Typography, Box, Button } from '@mui/material';

export default function HomePage() {
  return (
    <Container maxWidth="sm" style={{ marginTop: '2rem' }}>
      <Typography variant="h1" align="center" gutterBottom>
        vtools.dev
      </Typography>
      <Typography variant="h2" align="center" gutterBottom>
        Developer Tools for Encoding/Decoding
      </Typography>

      <Box display="flex" justifyContent="space-around" mt={4}>
        <Link href="/base64" passHref>
          <Button variant="contained" color="primary">
            Base64 Encoder/Decoder
          </Button>
        </Link>
        <Link href="/jwt" passHref>
          <Button variant="contained" color="secondary">
            JWT Encoder/Decoder
          </Button>
        </Link>
      </Box>
    </Container>
  );
}

