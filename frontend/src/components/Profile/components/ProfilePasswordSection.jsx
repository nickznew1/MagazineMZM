
import {memo, useState} from "react";
import {useForm} from "react-hook-form";
import "./style.css"
import {config} from "../../../config";


const ProfilePassword = ({profile}) => {
    const { user} = profile;
    const {
        register,
        handleSubmit,
        formState: { errors },
        reset,
        watch,
    } = useForm({
        mode: "onBlur",
        values:{
            old_password:"",
            new_password:"",
            confirm_new_password:"",
        }
    });
    const validatePassword = (value) => {
        if (value.length < 8) {
            return 'Пароль должен состоять из минимум 8 символов';
        } else if (!/[A-Z]/.test(value)) {
            return 'Пароль должен содержать как минимум одну заглавную букву';
        } else if (!/[0-9]/.test(value)) {
            return 'Пароль должен содержать как минимум одну цифру';
        }
        return true;
    };

    const password = watch("new_password")

    const [visiblePasswordChange, setVisiblePasswordChange] = useState(false)


    const togglePasswordChange = () => {
        setVisiblePasswordChange(!visiblePasswordChange);
        reset ({
            new_password:"",
            old_password:"",
            confirm_new_password:"",
        })
    }

        const handleChangePassword = async (data) => {

            const payload = {
                old_password: data.old_password,
                new_password: data.new_password,
                id: user.id,
            }
            try {
                const response = await fetch(`${config.apiUrl}/profile/changep`, {
                    method: 'PUT',
                    headers: {
                        'Content-Type': 'application/json',
                    },
                    body: JSON.stringify(payload),
                })

                setVisiblePasswordChange(false);
            } catch (error) {
                throw new Error(error)
            }
        }
    return (
        <div className = "user-profile-password">
            <div className ="user-profile-password-change__btn-wrapper">
            <div className ="user-profile-password-change__btn">
        <button type ="button" className={visiblePasswordChange ? "active" : ""} onClick = {togglePasswordChange}>Изменить пароль</button>
                </div>
            </div>
        {visiblePasswordChange ?  (
            <form onSubmit={handleSubmit(handleChangePassword)}>
            <div className="user-profile-password__form-wrapper">
                <div className = "user-profile-password__form">
                <input type="password"
                       autoComplete= "current-password"
                       placeholder = "Введите старый пароль"
                       style={{ borderColor: errors.old_password ? 'red' : '#ccc'  }}
                       {...register("old_password", {
                       })
                       }>
                </input>
                {errors.old_password && <span className = "user-profile-errors">{errors.old_password.message}</span> }
            </div>
                <div className = "user-profile-password__form">
                <input type="password"
                       placeholder = "Введите новый пароль"
                       autoComplete= "new-password"
                       style={{ borderColor: errors.new_password ? 'red' : '#ccc'  }}
                       {...register("new_password", {
                        validate: validatePassword,
                       })
                       }>
                </input>
                {errors.new_password && <span className = "user-profile-errors">{errors.new_password.message}</span> }
                </div>
                <div className = "user-profile-password__form">
                <input type="password"
                       autoComplete= "new-password"
                       placeholder = "Введите новый пароль еще раз"
                       style={{ borderColor: errors.confirm_new_password ? 'red' : '#ccc'  }}
                       {...register("confirm_new_password", {
                           validate: (value) => value === password || "Пароли не совпадают",
                       })
                       }>
                </input>
                {errors.confirm_new_password && <span className = "user-profile-errors">{errors.confirm_new_password.message}</span> }
                </div>
            <div className ="user-profile-submit__btn-wrapper">
            <div className = "user-profile-submit__btn">
                    <button type = "submit">Отправить</button>
            </div>
            </div>
            </div>
            </form>
        ) : (
            <div></div>
        )}
            </div>
    )
}

export default memo(ProfilePassword);