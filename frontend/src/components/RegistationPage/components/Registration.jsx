import {useNavigate} from "react-router";
import {useAuth} from "../../../context/AuthContext";
import {memo, useState} from "react";
import {useForm} from "react-hook-form";
import "./style.css"
import {config} from "../../../config";



const Registration = () => {
    const {
        register,
        handleSubmit,
        formState: { errors }
    } = useForm({
        mode: "onBlur"
    });
    const navigate = useNavigate()
    const {setUser, setToken} = useAuth()


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

    const validateEmail = (value) => {
        if (!/^[A-Z0-9._%+-]+@[A-Z0-9.-]+\.[A-Z]{2,4}$/i.test(value)){
            return "Некорректный email"
        } else if (value === "") {
            return "Поле обязательно"
        }
    }

    const [error, setError] = useState('')




    const onSubmit = async (data) => {
        const payload = {
            login: data.login,
            email: data.email,
            password: data.password,
        }
        try {
            const response = await fetch(`${config.apiUrl}/auth/registry`, {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json',
                },
                body: JSON.stringify(payload),
            })
            if (!response.ok) {
                setError("Пользователь с такими данными уже существует!")
            } else {
                const result = await response.json();
                setToken(result.access_token)
                setUser(result)
                navigate('/');
            }
        } catch (error){
            throw new Error ("Ошибка при регистрации")
        }
    }

    return (
        <>
            <div className = "user-cab">
                <div className = "user-cab__content">
        <div className ="user-cab__info">
            <h2>Пожалуйста, зарегистрируйтесь</h2>
        </div>
    <form onSubmit={handleSubmit(onSubmit)}>
        <div className = "user-cab__fields">
            <label className="text" form="input">Логин:</label>
            <input
                type="text"
                placeholder="Введите логин"
                style={{ borderColor: errors.login ? 'red' : '#ccc'  }}
                {...register("login", {
                    required: 'Логин обязателен',
                    validate: validateLogin,
                })
                }>
            </input>
            {errors.login && <span className="user-cab__errors">>{errors.login.message}</span>}
        </div>
        <div className = "user-cab__fields">
            <label className="text" form="input">Пароль:</label>
            <input
                type="password"
                placeholder="Введите пароль"
                style={{ borderColor: errors.password ? 'red' : '#ccc'  }}
                {...register("password", {
                    required: 'Пароль обязателен',
                    validate: validatePassword,
                })
                }>
            </input>
        </div>
        <div className = "user-cab__fields">
            <label className="text" form="input">Повторите пароль:</label>
            <input type= "password"
                   name ="confirm_password"
                   placeholder="повторите пароль"
                   style={{ borderColor: errors.password ? 'red' : '#ccc'  }}
                {...register("confirm_password", {
                    required: 'Password is required',
                    validate: validatePassword,
                })
                   }></input>
        {errors.password && <span className="user-cab__errors">>{errors.password.message}</span>}
        </div>
        <div className = "user-cab__fields">
            <label className="text" form="input">Email:</label>
            <input
                type="email"
                style={{ borderColor: errors.email ? 'red' : '#ccc'  }}
                {...register("email", {
                    required: 'Email обязателен',
                    validate: validateEmail,
                })
                }></input>
            {errors.email && <span className="user-cab__errors">{errors.email.message}</span>}
            {error && <span className="user-cab__errors">{error}</span> }
        </div>
        <div className="user-cab__btns">
            <button type = "submit">Зарегистрироваться</button>
        </div>
    </form>
                </div>
            </div>
        </>
    )
}

export default memo(Registration);