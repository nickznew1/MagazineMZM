import {useNavigate} from "react-router";
import {useAuth} from "../../../context/AuthContext";
import {memo, useState} from "react";
import {useForm} from "react-hook-form"
import "./style.css"
import {config} from "../../../config";


const Auth = () => {
    const {
        register,
        handleSubmit,
        formState: { errors }
    } = useForm({
        mode: "onBlur"
    });
    const navigate = useNavigate()
    const {setUser, setToken} = useAuth();
    const [error, setError] = useState('')

    const validatePassword = (value) => {
        if (value.length < 8) {
            return 'Пароль должен состоять из минимум 8 символов';
        } else if (!/[A-Z]/.test(value)) {
            return 'Пароль должен содержать как минимум одну заглавную букву';
        } else if (!/[0-9]/.test(value)) {
            return 'Пароль должен содержать как минимум одну цифру';
        } else if (value === ""){
            return "Поле обязательно"
        }
        return true;
    };

    const validateLogin = (value) => {
        if (value.length < 6) {
            return 'Логин должен состоять из минимум 6 символов';
        } else if (!/^[a-zA-Z0-9]+$/.test(value)) {
            return "Логин должен состоять только из латинских букв и цифр"
        } else if (value === "") {
            return "Поле обязательно"
        }
        return true
    }


    const onSubmit = async (data) => {
        const payload = {
            login:data.login,
            password:data.password,
        }
        try {
            const response = await fetch(`${config.apiUrl}/auth`, {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json',
                },
                body: JSON.stringify(payload),
            })
            if (!response.ok) {
                setError("Такого пользователя не существует!")
                return
            }
            const result = await response.json();
            setToken(result.access_token)

            setUser(payload)
            navigate('/')
        } catch (error){
            throw new Error(error)
        }
    }

    return (
        <div className="user-cab">
            <div className="user-cab__content">
    <div className="user-cab__info">
       <h1>Войдите в профиль</h1>
    <form id ="submit" onSubmit={handleSubmit(onSubmit)}>
            <div className = "user-cab__fields">
            <label className="text" form="input">Логин:</label>
            <input type="text"
                   placeholder ="Введите логин"
                   style={{ borderColor: errors.login ? 'red' : '#ccc' }}
                   {...register("login", {
                       required: 'Логин обязателен',
                       validate: validateLogin,
                   })
                   }>
            </input>
            {errors.login && <span className="user-cab__errors">{errors.login.message}</span> }
        </div>
            <div className = "user-cab__fields">
            <label className="text" form="input">Пароль:</label>
            <input type="password"
                   placeholder = "Введите пароль"
                   style={{ borderColor: errors.password ? 'red' : '#ccc' }}
                   {...register("password", {
                       required: 'Пароль обязателен',
                       validate: validatePassword,
                   })
                   }>
            </input>
            {errors.password && <span className="user-cab__errors">{errors.password.message}</span> }
                {error && <span className="user-cab__errors">{error}</span> }

        </div>
        <div className = "user-cab__fields">
        </div>
    </form>
    <div className="user-cab__btns">
        <button type="submit" form = "submit">Войти в профиль</button>
        <button onClick={() => navigate('/auth/registry')}>Еще нет профиля? Зарегистрируйтесь!</button>
    </div>
    </div>
            </div>
        </div>
)
    }

export default memo(Auth);
