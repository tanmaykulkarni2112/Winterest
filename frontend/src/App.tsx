import { BrowserRouter, Route, Routes } from 'react-router-dom'
import './App.css'
import React, { Suspense } from 'react'
const Landing = React.lazy(()=>import("./pages/Landing"))
const Ideas = React.lazy(():any=>import("./pages/Ideas"))

function App() {

  return(
    <>
      <div>
          <BrowserRouter>
            <Routes>
              <Route path='/' element={<Suspense fallback="loading the page,please wait...."><Landing/></Suspense>}/>
              <Route path='/ideas' element={<Suspense fallback="loading the page,please wait...."><Ideas/></Suspense>}/>
            </Routes>
          </BrowserRouter>
      </div>
    </>
  )
}

export default App
