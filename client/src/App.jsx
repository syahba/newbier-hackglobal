import { useState } from 'react'
import { BrowserRouter as Router, Route, Routes } from 'react-router-dom';

import NotFound from './pages/NotFound';
import Index from './pages/Index';
import Home from './pages/Home';
import Main from './layouts/Main';

function App() {

  return (
    <Router>
      <Routes>
        <Route path="/" element={<Index />} />
        <Route path="/home" element={<Main><Home /></Main>} />
        {/* Route 404 untuk halaman yang tidak ditemukan */}
        <Route path="*" element={<NotFound />} />
      </Routes>
    </Router>
  )
}

export default App
