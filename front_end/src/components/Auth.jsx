import './Auth.css'
import { useState } from 'react'

const Auth = () => {
    const [loginState, setLoginState] = useState('')
    const [passwordState, setPasswordState] = useState('')

    const signIn = () => {
        console.log(loginState, passwordState)
        // TODO: сделать запрос к backend
    }

    const seePassword = () => {
        const el = document.getElementById('passwordInput')
        if (el.type === 'password') {
            el.type = 'text'
        } else {
            el.type = 'password'
        }
    }

    return <>
        <div className='container'>
            <img src='auth.png' width='250px' alt='auth.png'/>
            <div className='login_input'>
                <input
                    type='text'
                    placeholder='Логин'
                    onChange={(e) => {
                        setLoginState(e.target.value)
                    }}
                    value={loginState}
                />
                <img src='user.png' className='icon' width='15px' height='15px' alt='user.png' />
            </div>
            <div className='password_input'>
                <input
                    id='passwordInput'
                    type='password'
                    placeholder='Пароль'
                    onChange={(e) => {
                        setPasswordState(e.target.value)
                    }}
                    value={passwordState}
                />
                <img src='hide.png' className='icon hide' onClick={seePassword} width='15px' height='15px' alt='hide.png'/>
            </div>
            <button className='sign_in' onClick={signIn}>Вход</button>
        </div>
    </>
}

export default Auth