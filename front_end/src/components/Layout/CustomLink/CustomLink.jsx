import { Link, useMatch } from 'react-router-dom'
import './CustomLink.css'

const CustomLink = ({ children, to, ...props }) => {
    const match = useMatch(to)
    // console.log(match)

    return <>
        <Link
            to={to}
            className={match ? 'active_link' : 'default_link'}
            {...props}
        >
            { children }
        </Link>
    </>
}

export default CustomLink