import React from 'react';
import './App.css';
import SignInForm from './features/auth/components/signin/sign-in-form';
import SignUpForm from './features/auth/components/signup/sign-up-form';

function App() {
  return (
    <div className="App">
     <SignInForm/>
    </div>
  );
}

export default App;
