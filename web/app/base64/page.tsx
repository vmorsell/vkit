'use client';

import { useState } from 'react';
import { Container, TextField, Button, Typography, Box } from '@mui/material';

export default function Base64Page() {
  const [input, setInput] = useState('');
  const [output, setOutput] = useState('');

  const encodeBase64 = (str: string): string => {
    return btoa(str);
  };

  const decodeBase64 = (str: string): string => {
    try {
      return atob(str);
    } catch (e) {
      return 'Invalid Base64 string!';
    }
  };

  return (
    <Container maxWidth="sm" style={{ marginTop: '2rem' }}>
      <Typography variant="h1" gutterBottom>
        Base64 Encoder/Decoder
      </Typography>

      <TextField
        label="Enter your text"
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
          onClick={() => setOutput(encodeBase64(input))}
        >
          Encode
        </Button>
        <Button
          variant="contained"
          color="secondary"
          onClick={() => setOutput(decodeBase64(input))}
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

