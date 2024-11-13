'use client';

import { useState } from 'react';
import { Container, TextField, Button, Typography, Box } from '@mui/material';

export default function JWTPage() {
  const [input, setInput] = useState('');
  const [output, setOutput] = useState('');

  const base64UrlEncode = (str: string): string => {
    return btoa(str).replace(/\+/g, '-').replace(/\//g, '_').replace(/=+$/, '');
  };

  const base64UrlDecode = (str: string): string => {
    str = str.replace(/-/g, '+').replace(/_/g, '/');
    try {
      return atob(str);
    } catch (e) {
      return 'Invalid JWT';
    }
  };

  return (
    <Container maxWidth="sm" style={{ marginTop: '2rem' }}>
      <Typography variant="h1" gutterBottom>
        JWT Encoder/Decoder
      </Typography>

      <TextField
        label="Enter your JWT"
        variant="outlined"
        fullWidth
        multiline
        rows={4}
        value={input}
        onChange={(e) => setInput(e.target.value)}
        margin="normal"
      />

      <Box display="flex" justifyContent="space-between" my={2}>
        <Button
          variant="contained"
          color="primary"
          onClick={() => setOutput(base64UrlEncode(input))}
        >
          Encode
        </Button>
        <Button
          variant="contained"
          color="secondary"
          onClick={() => setOutput(base64UrlDecode(input))}
        >
          Decode
        </Button>
      </Box>

      <TextField
        label="Output"
        variant="outlined"
        fullWidth
        multiline
        rows={4}
        value={output}
        InputProps={{
          readOnly: true,
        }}
        margin="normal"
      />
    </Container>
  );
}

