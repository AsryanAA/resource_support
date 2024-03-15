import React from 'react'
import './App.css'
import { Route, Routes } from 'react-router-dom'
import Layout from './components/Layout/Layout'
import Accessors from './components/Dictionaries/Accessors/Accessors'
import Auth from './components/Auth'
import Munitions from './components/Dictionaries/Munition/Munitions'
import PageNotFound from './components/PageNotFound/PageNotFound'

const App = () => {
    return <>
        <Routes>
            <Route index element={ <Auth />} />
            <Route path='/' element={ <Layout /> }>
                <Route path='file' element={ <h1>File</h1> } />
                <Route path='documents' element={ <h1>Documents</h1> } />
                <Route path='accounting' element={ <h1>Accounting</h1> } />
                <Route path='functions' element={ <h1>Functions</h1> } />
                <Route path='reports' element={ <h1>Reports</h1> } />
                <Route path='dictionaries' element={ <h1>Выберите словарь</h1> } />
                <Route path='dictionaries/accessors' element={ <Accessors /> } />
                <Route path='dictionaries/munitions' element={ <Munitions /> } />
                <Route path='window' element={ <h1>Window</h1> } />
                <Route path='help' element={ <h1>Help</h1> } />
                <Route path='*' element={ <PageNotFound /> } />
            </Route>
        </Routes>
    </>
}

export default App