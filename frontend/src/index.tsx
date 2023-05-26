import ReactDOM from 'react-dom/client';
import './index.css';
import { BrowserRouter } from 'react-router-dom';
import { AuthProvider } from './features/auth/context/auth';
import App from './components/app/app';
import { Provider } from 'react-redux'
import { store } from './api/store/store';

const root = ReactDOM.createRoot(
  document.getElementById('root') as HTMLElement
);
root.render(
  <Provider store={store}>
     <BrowserRouter>
     <AuthProvider>
      <App />
     </AuthProvider>
    </BrowserRouter>
    </Provider>
);
