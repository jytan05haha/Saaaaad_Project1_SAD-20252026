import Auth from "./pages/auth"
import './App.css'
import { BrowserRouter, Routes, Route } from 'react-router-dom'
import { MantineProvider } from '@mantine/core';
import { useEffect, useState } from "react";
import LoggedInFlow from "./pages/app";
import Home from "./pages/home";
import Orders from "./pages/orders";

function App() {

  const [isLoggedIn, setIsLoggedIn] = useState(false);

  useEffect(() => {
    const token= localStorage.getItem('token');
    if(token){
      setIsLoggedIn(true);
    }
    },[]);
  
  return (
    <MantineProvider>
      <BrowserRouter>
        <Routes>
          <Route path="/" element={isLoggedIn ? <LoggedInFlow/>: <Auth /> } />
          <Route path="/home" element={<Home />} />
          <Route path="/orders" element={isLoggedIn ? <Orders/>: <Auth />} />
        </Routes>
      </BrowserRouter>
    </MantineProvider>
  );
}

export default App
