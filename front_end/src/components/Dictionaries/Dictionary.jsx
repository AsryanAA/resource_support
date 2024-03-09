import { useAccessors } from '../../store'
import { Route, Routes } from 'react-router-dom'
import { useEffect } from 'react'

const Dictionary = () => {
    const data = useAccessors(state => state.readAccessors)
    useEffect(() => {
        data()
    }, [])

    const accessors = useAccessors(state => state.accessors)
    // console.log(munitions)

    return <>
        <Routes>
            <Route path='accessors' element={ <h1>Accessors</h1> }/>
        </Routes>
    </>
}

export default Dictionary