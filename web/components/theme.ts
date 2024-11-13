import { createTheme } from '@mui/material/styles';

const theme = createTheme({
  palette: {
    primary: {
      main: '#1976d2',  // Blue
    },
    secondary: {
      main: '#ff4081',  // Pink
    },
    background: {
      default: '#f5f5f5',  // Light Gray
    },
  },
  typography: {
    h1: {
      fontSize: '2.5rem',
      fontWeight: 'bold',
    },
    h2: {
      fontSize: '2rem',
    },
    body1: {
      fontSize: '1rem',
    },
  },
});

export default theme;

