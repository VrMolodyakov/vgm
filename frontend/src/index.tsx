import ReactDOM from 'react-dom/client';
import './index.css';
import { BrowserRouter } from 'react-router-dom';
import { AuthProvider } from './features/auth/context/auth';
import App from './app/app';
import { QueryClient, QueryClientProvider } from '@tanstack/react-query';
import { ClientProvider } from './features/music/client-provider/provider';

const queryClient = new QueryClient();

const root = ReactDOM.createRoot(
  document.getElementById('root') as HTMLElement
);
root.render(
  <BrowserRouter>
    <QueryClientProvider client={queryClient}>
      <AuthProvider>
        <ClientProvider>
          <App />
        </ClientProvider>
      </AuthProvider>
    </QueryClientProvider>
  </BrowserRouter>
);
