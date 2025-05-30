import { useState } from 'react'
import reactLogo from './assets/react.svg'
import viteLogo from '/vite.svg'
import './App.css'
import CheckinPage from "./components/CheckinPage.js";

function App() {
  const [count, setCount] = useState(0)

  return (
    <>
        <CheckinPage/>
    </>
  )
}

export default App
