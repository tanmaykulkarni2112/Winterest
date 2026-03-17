import { BrowserRouter, Route, Routes } from 'react-router-dom'
import './App.css'
import React, { Suspense } from 'react'
const Landing = React.lazy(()=>import("./pages/Landing"))

function App() {

  return(
    <>
      <div>
          <BrowserRouter>
            <Routes>
              <Route path='/' element={<Suspense fallback="loading the page,please wait...."><Landing/></Suspense>}/>
            </Routes>
          </BrowserRouter>
      </div>
    </>
  )
}

export default App
