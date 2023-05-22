import React from 'react';
import ReactDOM from 'react-dom/client';
import './index.css';
import { BrowserRouter } from 'react-router-dom';
import { AuthProvider } from './features/auth/context/auth';
import App from './components/app/app';


const root = ReactDOM.createRoot(
  document.getElementById('root') as HTMLElement
);
root.render(
     <BrowserRouter>
     <AuthProvider>
      <App />
     </AuthProvider>
    </BrowserRouter>
);
