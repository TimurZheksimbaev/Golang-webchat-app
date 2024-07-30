"use client"
import "./globals.css"
import LoginPage from "./login/page";
import AuthContextProvider from "@/modules/auth_provider";
import { Routes, Route } from 'react-router-dom';
import NotFound from "./notfound/NotFound";
import HomePage from "./home/HomePage";
import { BrowserRouter } from 'react-router-dom';
import WebSocketProvider from "@/modules/websocket_provider";

// className="flex flex-col md:flex-row h-full min-h-screen font-sans"
export default function Home() {
  return (
    <>
      <div>
        <AuthContextProvider> 
          <WebSocketProvider>
            <BrowserRouter>
              <Routes>
                <Route path="/" element={<HomePage />} />
                <Route path="/login" element={<LoginPage />} />
                <Route path="*" element={<NotFound />} />
              </Routes>
            </BrowserRouter>
          </WebSocketProvider>
        </AuthContextProvider>
      </div>
    </>
  );
}
