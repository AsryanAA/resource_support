import { Outlet } from 'react-router-dom'
import CustomLink from './CustomLink/CustomLink'
import './Layout.css'

const Layout = () => {
    return <>
        <header>
            <nav className='nav'>
                <CustomLink to='/file'>Файл</CustomLink>
                <CustomLink to='/documents'>Документы</CustomLink>
                <CustomLink to='/accounting'>Учет</CustomLink>
                <CustomLink to='/functions'>Функции</CustomLink>
                <CustomLink to='/reports'>Отчеты</CustomLink>
                <CustomLink to='/dictionaries'>Словари</CustomLink>
                <CustomLink to='/window'>Окно</CustomLink>
                <CustomLink to='/help'>Справка</CustomLink>
            </nav>
        </header>
        <Outlet />
        <footer>

        </footer>
    </>
}

export default Layout